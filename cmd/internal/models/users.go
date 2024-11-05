package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id             string
	Name           string
	Email          string
	Hashedpassword []byte
	Create         time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (um *UserModel) Insert(name, email, password string) error {
	return nil
}

func (um *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (um *UserModel) Exists(userId string) (bool, error) {
	return true, nil
}
