package app

// App represents an app resource
type App struct {
	ID   string
	Name string
	User string // TODO: array? team?
}

type AppRepository interface {
	Create(app *App) (string, error)
	Get(name string) (*App, error)
}
