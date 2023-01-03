package service

import (
	"context"
	"errors"
	"testing"

	"github.com/disturb/inventory/encryption"
	"github.com/disturb/inventory/internal/model"
)

func (s *Serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return errors.New("user already exists")
	}
	//todo:hash the password
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass, _ := encryption.ToBase64(bb)
	return s.repo.SaveUser(ctx, email, name, pass)
}
func (s *Serv) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	//todo:descrypt password
	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decrypterPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}
	if string(decrypterPassword) != password {
		return nil, errors.New("invalid credentials")
	}
	return &model.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test01",
			Password:      "validpassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_AlreadyExists",
			Email:         "test01@test.com",
			UserName:      "test01",
			Password:      "validpassword",
			ExpectedError: errors.New("user already exists"),
		},
	}
}

ctx := context.Background()

for i := range testCases {
	tc := testCases[i]

	t.Run(tc.Name, func(t *testing.T) {
		t.Parallel()
		repo := &repository.MockRepository{}
		s := New(repo)
		err := s.RegisterUser(ctx, tc.email, tc.UserName, tc.Password )

		if err != tc.ExpectedError {
			t.errorf("Expected error %v, got %v ", tc.ExpectedError, err)
		}
	})
}

