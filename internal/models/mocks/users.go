package mocks

import (
	"github.com/pauloa.junior/mynotes/internal/models"
)

type FakeUserRepository struct{}

func (r *FakeUserRepository) Insert(name, email, password string) error {
	switch email {
	case "paulo@paulo.com":
		return models.ErrDuplicatedEmail
	default:
		return nil
	}
}

func (r *FakeUserRepository) Authenticate(email, password string) (int, error) {
	if email == "paulo@paulo.com" && password == "12345678" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (r *FakeUserRepository) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}
