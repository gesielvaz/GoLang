package service

import (
	"context"

	"github.com/disturb/inventory/internal/model"
	"github.com/disturb/inventory/internal/repository"
)

//go:generate monckey --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}
type Serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &Serv{
		repo: repo,
	}
}
