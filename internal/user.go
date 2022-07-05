package internal

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Save(ctx context.Context, user User) error
}

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.FromString(value)

	if err != nil {
		return UserID{}, err
	}
	return UserID{value: v.String()}, nil
}

type UserName struct {
	value string
}

func NewUserName(value string) (UserName, error) {
	if len(value) == 0 || value == "" {
		return UserName{}, errors.New("username is empty")
	}
	return UserName{value: value}, nil
}

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (UserEmail, error) {
	if len(value) == 0 || value == "" {
		return UserEmail{}, errors.New("email is empty")
	}
	return UserEmail{value: value}, nil
}

type UserPassword struct {
	value string
}

func NewUserPassword(value string) (UserPassword, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return UserPassword{}, err
	}

	return UserPassword{value: string(hashedPassword)}, nil
}

type User struct {
	id       UserID
	name     UserName
	email    UserEmail
	password UserPassword
}

func NewUser(id, name, email, password string) (User, error) {
	voId, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}

	voName, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	voEmail, err := NewUserEmail(email)
	if err != nil {
		return User{}, err
	}

	voPassword, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	return User{
		id:       voId,
		name:     voName,
		email:    voEmail,
		password: voPassword,
	}, nil
}

func (u User) ID() UserID {
	return u.id
}

func (u User) Name() UserName {
	return u.name
}

func (u User) Email() UserEmail {
	return u.email
}
func (u User) Password() UserPassword {
	return u.password
}
