package db

import (
	"errors"

	"bitbucket.org/jtblin/kigo-api/pkg/domain/user"

	r "github.com/dancannon/gorethink"
	"golang.org/x/crypto/bcrypt"
)

const userTable = "user"
const hashCost = 9

// UserModel represents an user resource
type UserModel struct {
	ID       string `gorethink:"id"`
	Name     string `gorethink:"name"`
	Password []byte `gorethink:"password"`
	Token    string `gorethink:"token"`
}

type userRepo struct {
	*RethinkClient
}

// Get returns a user
func (ur *userRepo) Get(email string) (*user.User, error) {
	cursor, err := r.Table(userTable).Get(email).Run(ur.session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	if cursor.IsNil() {
		return nil, errors.New("no record found")
	}
	var model UserModel
	if err = cursor.One(&model); err != nil {
		return nil, err
	}
	return &user.User{
		Email:    model.ID,
		Name:     model.Name,
		Password: string(model.Password),
		Token:    model.Token,
	}, nil
}

// Create creates a new user
func (ur *userRepo) Create(usr *user.User) (string, error) {
	if usr.Email == "" || usr.Password == "" {
		return "", errors.New("missing required field")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), hashCost)
	_, err = r.Table(userTable).Insert(UserModel{
		ID:       usr.Email,
		Name:     usr.Name,
		Password: hash,
		Token:    usr.Token,
	}).RunWrite(ur.session)
	if err != nil {
		return "", err
	}
	return usr.Email, nil
}

// NewUserRepository is an implementation for a UserRepository
func NewUserRepository(c *RethinkClient) user.UserRepository {
	return &userRepo{
		RethinkClient: c,
	}
}

func init() {
	tables = append(tables, userTable)
}
