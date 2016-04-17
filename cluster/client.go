package cluster

import (
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

// TODO: merge cluster and job?

type Client struct {
	*client.Client
}

// NewClient returns a cluster client
func NewClient(host, username, token string) (*Client, error) {
	var c *client.Client
	var err error
	if host != "" && username != "" && token != "" {
		config := &restclient.Config{
			Host:        host,
			Username:    username,
			BearerToken: token,
		}
		c, err = client.New(config)
	} else {
		c, err = client.NewInCluster()
	}
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}
