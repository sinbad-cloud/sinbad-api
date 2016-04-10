package app

// App represents an app resource
type App struct {
	Name string
	User string // TODO: array? team?
}

// AppRepository represents an app interface
type AppRepository interface {
	Create(app *App) (string, error)
	Get(name string) (*App, error)
}
