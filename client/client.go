package client

import (
	clusterapi "github.com/ehazlett/stellar/api/services/cluster/v1"
	datastoreapi "github.com/ehazlett/stellar/api/services/datastore/v1"
	healthapi "github.com/ehazlett/stellar/api/services/health/v1"
	networkapi "github.com/ehazlett/stellar/api/services/network/v1"
	nodeapi "github.com/ehazlett/stellar/api/services/node/v1"
	versionapi "github.com/ehazlett/stellar/api/services/version/v1"
	"google.golang.org/grpc"
)

type Client struct {
	conn             *grpc.ClientConn
	versionService   versionapi.VersionClient
	healthService    healthapi.HealthClient
	nodeService      nodeapi.NodeClient
	clusterService   clusterapi.ClusterClient
	datastoreService datastoreapi.DatastoreClient
	networkService   networkapi.NetworkClient
}

func NewClient(addr string) (*Client, error) {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: c,
	}
	client.versionService = versionapi.NewVersionClient(c)
	client.healthService = healthapi.NewHealthClient(c)
	client.nodeService = nodeapi.NewNodeClient(c)
	client.clusterService = clusterapi.NewClusterClient(c)
	client.datastoreService = datastoreapi.NewDatastoreClient(c)
	client.networkService = networkapi.NewNetworkClient(c)

	return client, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Node() *node {
	return &node{
		client: c.nodeService,
	}
}

func (c *Client) Cluster() *cluster {
	return &cluster{
		client: c.clusterService,
	}
}

func (c *Client) Datastore() *datastore {
	return &datastore{
		client: c.datastoreService,
	}
}

func (c *Client) Network() *network {
	return &network{
		client: c.networkService,
	}
}

func (c *Client) Version() *version {
	return &version{
		client: c.versionService,
	}
}

func (c *Client) Health() *health {
	return &health{
		client: c.healthService,
	}
}

func (c *Client) VersionService() versionapi.VersionClient {
	return c.versionService
}

func (c *Client) HealthService() healthapi.HealthClient {
	return c.healthService
}

func (c *Client) NodeService() nodeapi.NodeClient {
	return c.nodeService
}

func (c *Client) ClusterService() clusterapi.ClusterClient {
	return c.clusterService
}

func (c *Client) DatastoreService() datastoreapi.DatastoreClient {
	return c.datastoreService
}

func (c *Client) NetworkService() networkapi.NetworkClient {
	return c.networkService
}
