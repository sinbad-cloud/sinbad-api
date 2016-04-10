package db

import (
	"errors"

	"bitbucket.org/jtblin/kigo-api/pkg/domain/app"

	r "github.com/dancannon/gorethink"
)

const appTable = "app"

// AppModel represents an app resource
type AppModel struct {
	ID     string            `gorethink:"id,omitempty"`
	User   string            `gorethink:"user,omitempty"` // TODO: array? team?
	Config map[string]string `gorethink:"config,omitempty"`
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
func (ar *appRepo) Get(name string) (*app.App, error) {
	cursor, err := r.Table(appTable).Get(name).Run(ar.session)
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
		Name:   model.ID,
		User:   model.User,
		Config: model.Config,
	}, nil
}

// Create creates a new app
func (ar *appRepo) Create(app *app.App) (string, error) {
	if app.Name == "" || app.User == "" {
		return "", errors.New("missing required field")
	}
	_, err := r.Table(appTable).Insert(AppModel{
		ID:     app.Name,
		User:   app.User,
		Config: app.Config,
	}).RunWrite(ar.session)
	if err != nil {
		return "", err
	}
	return app.Name, nil
}

// GetConfig returns an app config
func (ar *appRepo) GetConfig(name, key string) (map[string]string, error) {
	app, err := ar.Get(name)
	if err != nil {
		return nil, err
	}
	if app.Config != nil {
		value, ok := app.Config[key]
		if ok {
			config := make(map[string]string, 1)
			config[key] = value
			return config, nil
		}
	}
	return nil, errors.New("no config found")
}

// SetConfig sets an app config
func (ar *appRepo) SetConfig(name, key, value string) (map[string]string, error) {
	app, err := ar.Get(name)
	if err != nil {
		return nil, err
	}
	if app.Config == nil {
		app.Config = make(map[string]string)
	}
	app.Config[key] = value

	if _, err = r.Table(appTable).Get(name).Update(AppModel{Config: app.Config}).RunWrite(ar.session); err != nil {
		return nil, err
	}

	config := make(map[string]string, 1)
	config[key] = value
	return config, nil
}

func init() {
	tables = append(tables, appTable)
}
