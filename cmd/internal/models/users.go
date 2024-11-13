package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModelInterface interface {
	Insert(userId, name, email, password string) error
	Authenticate(email, password string) (string, error)
	Exists(id string) (bool, error)
	Get(userId string) (User, error)
}
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

func (um *UserModel) Insert(userId, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `INSERT INTO users(id, name, email, hash_password, created) VALUES (?, ?, ?, ?, UTC_TIMESTAMP())`

	_, err = um.DB.Exec(query, userId, name, email, string(hashedPassword))

	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}

		return err
	}
	return nil
}

func (um *UserModel) Authenticate(email, password string) (string, error) {
	var id string
	var hashedPassword []byte

	stmt := `SELECT id, hash_password FROM users WHERE email = ?`
	row := um.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrInvalidCredentials
		} else {
			return "", err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", ErrInvalidCredentials
		} else {
			return "", err
		}
	}

	return id, nil
}

func (um *UserModel) Exists(userId string) (bool, error) {
	var exists bool
	stmt := `SELECT EXISTS(SELECT true FROM users WHERE id = ?)`

	err := um.DB.QueryRow(stmt, userId).Scan(&exists)
	return exists, err
}

func (um *UserModel) Get(userId string) (User, error) {
	query := `SELECT id, name, email, hash_password, created FROM users WHERE id = ?`

	row := um.DB.QueryRow(query, userId)

	user := User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Hashedpassword, &user.Create)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrNoRecord
		} else {
			return User{}, err
		}
	}

	return user, nil
}
