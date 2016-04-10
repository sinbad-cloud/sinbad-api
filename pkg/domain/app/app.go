package app

// App represents an app resource
type App struct {
	Name    string
	Config  map[string]string
	RepoURL string
	User    string // TODO: array? team?
}

// AppRepository represents an app interface
type AppRepository interface {
	Create(app *App) (string, error)
	Get(name string) (*App, error)
	GetConfig(name, key string) (map[string]string, error)
	SetConfig(name, key, value string) (map[string]string, error)
}
