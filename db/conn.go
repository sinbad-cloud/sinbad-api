package db

import (
	log "github.com/Sirupsen/logrus"
	r "github.com/dancannon/gorethink"
)

const (
	db = "kigo"
)

var session *r.Session
var tables []string

func initTables() error {
	cursor, err := r.TableList().Run(session)
	if err != nil {
		return err
	}
	defer cursor.Close()

	var row string
	var existingTables = make(map[string]bool)

	for cursor.Next(&row) {
		existingTables[row] = true
	}

	for _, table := range tables {
		if _, ok := existingTables[table]; !ok {
			if err = r.TableCreate(table).Exec(session); err != nil {
				return err
			}
			log.Infof("Created table \"%s.%s\"", db, table)
		}
	}

	return nil
}

func initDB() error {
	cursor, err := r.DBList().Run(session)
	if err != nil {
		return err
	}
	defer cursor.Close()

	var row string
	var exists bool
	for cursor.Next(&row) {
		if row == db {
			exists = true
			break
		}
	}
	if !exists {
		if err = r.DBCreate(db).Exec(session); err != nil {
			return err
		}
		log.Infof("Created db \"%s\"", db)
	}
	return nil
}

// NewConnection returns a singleton rethinkdb connection
func NewConnection() (*r.Session, error) {
	var err error
	if session != nil {
		return session, nil
	}
	session, err = r.Connect(r.ConnectOpts{
		Addresses:     []string{"localhost:28015"},
		Database:      db,
		DiscoverHosts: true,
	})
	if err != nil {
		return nil, err
	}

	if err = initDB(); err != nil {
		return nil, err
	}

	if err = initTables(); err != nil {
		return nil, err
	}

	return session, nil
}
