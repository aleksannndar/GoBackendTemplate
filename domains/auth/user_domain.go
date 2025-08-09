package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Id       string
	Email    string
	Password *string
}

func CreateNewuser(email string, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("email can't be empty")
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error while hashing password: %w", err)
	}

	user := &User{
		Id:       uuid.New().String(),
		Email:    strings.ToLower(email),
		Password: &hashedPassword,
	}

	return user, nil
}

func (u *User) ToEntity() *UserEntity {
	return &UserEntity{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}
