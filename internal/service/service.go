package service

import (
	"context"

	"github.com/Jenietor2/app-inventory/internal/models"
	"github.com/Jenietor2/app-inventory/internal/repository"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct{
	repo repository.Repository
}

func New(repo repository.Repository)Service{
	return &serv{
		repo: repo,
	}
}