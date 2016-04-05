package db

import (
	"errors"

	"bitbucket.org/jtblin/kigo-api/pkg/domain/app"

	r "github.com/dancannon/gorethink"
)

const appTable = "app"

// App represents an app resource
type AppModel struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
	User string `gorethink:"user"` // TODO: array? team?
}

type appRepo struct {
	*RethinkClient
}

// NewAppRepository is an implementation for an AppRepository
func NewAppRepository(c *RethinkClient) app.AppRepository {
	return &appRepo{
		RethinkClient: c,
	}
}

// Get returns an app
func (ar *appRepo) Get(ID string) (*app.App, error) {
	cursor, err := r.Table(appTable).Get(ID).Run(ar.session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if cursor.IsNil() {
		return nil, errors.New("no record found")
	}
	var model AppModel
	if err = cursor.One(&model); err != nil {
		return nil, err
	}
	return &app.App{
		Name: model.Name,
		User: model.User,
	}, nil
}

// Create creates a new app
func (ar *appRepo) Create(app *app.App) (string, error) {
	if app.Name == "" || app.User == "" {
		return "", errors.New("missing required field")
	}
	response, err := r.Table(appTable).Insert(AppModel{
		Name: app.Name,
		User: app.User,
	}).RunWrite(ar.session)
	if err != nil {
		return "", err
	}
	return response.GeneratedKeys[0], nil
}

func init() {
	tables = append(tables, appTable)
}
