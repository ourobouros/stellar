package proxy

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/containerd/containerd"
	"github.com/ehazlett/element"
	"github.com/ehazlett/stellar"
	api "github.com/ehazlett/stellar/api/services/proxy/v1"
	"github.com/ehazlett/stellar/client"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/route"
	"google.golang.org/grpc"
)

const (
	serviceID = "stellar.services.proxy.v1"
)

var (
	empty = &ptypes.Empty{}
)

type service struct {
	containerdAddr string
	namespace      string
	agent          *element.Agent
	proxyHTTPPort  int
	errCh          chan error
	updateCh       chan *proxyUpdate
	currentServers map[string]*backend
	mux            *route.Mux
}

type backend struct {
	host    string
	lb      *roundrobin.RoundRobin
	servers []*url.URL
}

type updateAction string

const (
	updateActionAdd    updateAction = "add"
	updateActionUpdate updateAction = "update"
	updateActionRemove updateAction = "remove"
)

type proxyUpdate struct {
	action  updateAction
	backend *backend
}

func New(cfg *stellar.Config, agent *element.Agent) (*service, error) {
	errCh := make(chan error)
	go func() {
		for {
			err := <-errCh
			logrus.Errorf("proxy: %s", err)
		}
	}()

	return &service{
		containerdAddr: cfg.ContainerdAddr,
		namespace:      cfg.Namespace,
		proxyHTTPPort:  cfg.ProxyHTTPPort,
		errCh:          errCh,
		updateCh:       make(chan *proxyUpdate),
		currentServers: make(map[string]*backend),
		mux:            route.NewMux(),
		agent:          agent,
	}, nil
}

func (s *service) Register(server *grpc.Server) error {
	api.RegisterProxyServer(server, s)
	return nil
}

func (s *service) ID() string {
	return serviceID
}

func (s *service) Start() error {
	go s.updater()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.proxyHTTPPort),
		Handler: s.mux,
	}

	t := time.NewTicker(5 * time.Second)
	go func() {
		for range t.C {
			if err := s.reload(); err != nil {
				logrus.Errorf("proxy: %s", err)
			}
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(errors.Wrap(err, "proxy"))
		}
	}()

	return nil
}

func (s *service) containerd() (*containerd.Client, error) {
	return stellar.DefaultContainerd(s.containerdAddr, s.namespace)
}

func (s *service) client() (*client.Client, error) {
	peer, err := s.agent.LocalNode()
	if err != nil {
		return nil, err
	}
	return client.NewClient(peer.Addr)
}

func (s *service) peerAddr() (string, error) {
	peer, err := s.agent.LocalNode()
	if err != nil {
		return "", err
	}

	return peer.Addr, nil
}

func (s *service) nodeName() string {
	return s.agent.Config().NodeName
}

func (s *service) nodeClient(id string) (*client.Client, error) {
	c, err := s.client()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	nodes, err := c.Cluster().Nodes()
	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		if node.Name == id {
			return client.NewClient(node.Addr)
		}
	}

	return nil, fmt.Errorf("node %s not found in cluster", id)
}

func newLB() (*roundrobin.RoundRobin, error) {
	// TODO: log separately?
	l := logrus.New()
	l.Out = ioutil.Discard
	fwd, err := forward.New(forward.Logger(l))
	if err != nil {
		return nil, err
	}
	lb, err := roundrobin.New(fwd, roundrobin.RoundRobinLogger(l))
	if err != nil {
		return nil, err
	}

	return lb, nil
}

func (s *service) getUpdateAction(id string) updateAction {
	if _, ok := s.currentServers[id]; ok {
		return updateActionUpdate
	}

	return updateActionAdd
}

func (s *service) generateServerID(v interface{}, extra ...interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write(data)
	for _, x := range extra {
		d, err := json.Marshal(x)
		if err != nil {
			return "", err

		}
		h.Write(d)
	}
	r := hex.EncodeToString(h.Sum(nil))[:24]
	return r, nil
}

func (s *service) pruneServers(next map[string]*backend) {
	for k, b := range s.currentServers {
		if _, exists := next[k]; !exists {
			logrus.Debugf("proxy: %s not found; removing", k)
			s.updateCh <- &proxyUpdate{
				backend: b,
				action:  updateActionRemove,
			}
		}
	}
}