package app

// App represents an app resource
type App struct {
	Name   string
	User   string // TODO: array? team?
	Config map[string]string
}

// AppRepository represents an app interface
type AppRepository interface {
	Create(app *App) (string, error)
	Get(name string) (*App, error)
	GetConfig(app, key string) (map[string]string, error)
	SetConfig(app, key, value string) (map[string]string, error)
}
