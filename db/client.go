package db

import (
	"time"

	log "github.com/Sirupsen/logrus"
	r "github.com/dancannon/gorethink"
)

// TODO: friendly errors

// RethinkClient represents a DB client implementation for rethinkdb
type rethinkClient struct {
	session *r.Session
}

func get() {

}

func (c *rethinkClient) getConnection(dbAddress string) {
	session, err := NewConnection(dbAddress)
	if err != nil {
		time.Sleep(time.Duration(1) * time.Second)
		c.getConnection(dbAddress)
		return
	}
	c.session = session
	// TODO: move to public method and trigger from executor
	if err = c.watchDeployments(); err != nil {
		log.Warnf("Error setting deployments watch: %s", err.Error())
		c.getConnection(dbAddress)
		return
	}
	if err = c.watchBuilds(); err != nil {
		log.Warnf("Error setting builds watch: %s", err.Error())
		c.getConnection(dbAddress)
		return
	}
}

// NewClient returns a db client
func NewClient(dbAddress string) *rethinkClient {
	client := &rethinkClient{session: &r.Session{}}
	go client.getConnection(dbAddress)
	return client
}
