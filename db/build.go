package db

import (
	"errors"
	"time"

	"bitbucket.org/jtblin/kigo-api/pkg/domain/build"

	r "github.com/dancannon/gorethink"
)

const buildTable = "build"

// BuildModel represents an build resource
type BuildModel struct {
	App          string    `gorethink:"app,omitempty"`
	Author       string    `gorethink:"author,omitempty"`
	Commit       string    `gorethink:"commit,omitempty"`
	ID           string    `gorethink:"id,omitempty"`
	Organization string    `gorethink:"organization,omitempty"`
	Origin       string    `gorethink:"origin,omitempty"`
	Repository   string    `gorethink:"repository,omitempty"`
	Timestamp    time.Time `gorethink:"timestamp,omitempty"`
}

type buildRepo struct {
	*rethinkClient
}

// WatchBuilds watch builds and send to queue
func (c *rethinkClient) watchBuilds() error {
	cursor, err := r.Table(buildTable).Changes().Run(c.session)
	if err != nil {
		return err
	}
	go func() {
		var model BuildModel
		// TODO: use channel and cursor.Listen instead?
		// TODO: check if new or updated or deleted
		// i.e. newValue != nil, oldValue != nil
		for cursor.Next(&model) {
			build.Enqueue(&build.Build{
			// TODO: fill this in
			})
		}
	}()
	return nil
}

// NewBuildRepository is an implementation for a BuildRepository
func NewBuildRepository(c *rethinkClient) build.BuildRepository {
	return &buildRepo{
		rethinkClient: c,
	}
}

// Get returns a build
func (br *buildRepo) Get(ID string) (*build.Build, error) {
	cursor, err := r.Table(buildTable).Get(ID).Run(br.session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if cursor.IsNil() {
		return nil, errors.New("no record found")
	}
	var model BuildModel
	if err = cursor.One(&model); err != nil {
		return nil, err
	}
	return &build.Build{
		App:          model.App,
		Author:       model.Author,
		Commit:       model.Commit,
		ID:           model.ID,
		Organization: model.Organization,
		Origin:       model.ID,
		Repository:   model.Repository,
		Timestamp:    model.Timestamp,
	}, nil
}

// Create creates a new build
func (br *buildRepo) Create(build *build.Build) (string, error) {
	if build.App == "" || build.Origin == "" || build.Organization == "" || build.Repository == "" {
		return "", errors.New("missing required field")
	}
	_, err := r.Table(buildTable).Insert(BuildModel{
		App:          build.App,
		Author:       build.Author,
		Commit:       build.Commit,
		ID:           build.ID,
		Organization: build.Organization,
		Origin:       build.ID,
		Repository:   build.Repository,
		Timestamp:    build.Timestamp,
	}).RunWrite(br.session)
	if err != nil {
		return "", err
	}
	return build.ID, nil
}

func init() {
	tables = append(tables, buildTable)
}
