package db

import (
	r "github.com/dancannon/gorethink"
)

// TODO: friendly errors

// RethinkClient represents a DB client implementation for rethinkdb
type RethinkClient struct {
	session *r.Session
}

// NewClient returns a db client
func NewClient() (*RethinkClient, error) {
	session, err := NewConnection()
	// TODO: make resilient to database not started i.e. retry
	if err != nil {
		return nil, err
	}
	client := &RethinkClient{session: session}
	// TODO: move to public method and trigger from executor
	if err = client.watchDeployments(); err != nil {
		return client, err
	}
	return client, nil
}
