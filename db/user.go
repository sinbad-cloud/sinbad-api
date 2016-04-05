package db

// User represents an user resource
type User struct {
	ID       string `gorethink:"id,omitempty"`
	Email    string `gorethink:"email,omitempty"`
	Name     string `gorethink:"name"`
	Password string `gorethink:"password"`
}
