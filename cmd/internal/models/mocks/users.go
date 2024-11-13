package mocks

import "snippetbox.art.net/cmd/internal/models"

type UserModel struct{}

func (m *UserModel) Insert(name, id, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (string, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return "UUID", nil
	}

	return "", models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id string) (bool, error) {
	switch id {
	case "UUID":
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(userId string) (models.User, error) {
	return models.User{}, nil
}
