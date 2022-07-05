package internal

import "context"

type UserRepository interface {
	Save(ctx context.Context, user User) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

type User struct {
	id       string
	name     string
	email    string
	password string
}

func NewUser(id, name, email, password string) User {
	return User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}

func (u User) ID() string {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}
