package service

import (
	"github.com/SolidShake/dividend-calendar-go/pkg/repository"
	"github.com/SolidShake/dividend-calendar-go/pkg/repository/models"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
