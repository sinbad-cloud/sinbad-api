package user

// User represents a user resource
type User struct {
	Email    string
	Name     string
	Password string
	Token    string
}

// UserRepository represents a user interface
type UserRepository interface {
	Create(app *User) (string, error)
	Get(email string) (*User, error)
}
