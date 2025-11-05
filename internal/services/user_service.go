package services

import (
	"GoATTHStart/internal/domain"
	"GoATTHStart/internal/models"
	"GoATTHStart/internal/repositories"
	"log/slog"

	"github.com/google/uuid"
)

type userService struct {
	userRepo repositories.UserRepository
	logger   *slog.Logger
}

func NewUserService(userRepo repositories.UserRepository, logger *slog.Logger) *userService {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

type UserService interface {
	CreateUser(user *domain.User) error
	GetUser(id uuid.UUID) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(user *domain.User) error
}

func (u *userService) CreateUser(user *domain.User) error {
	userModel := models.FromDomainUser(*user)
	err := u.userRepo.CreateUser(&userModel)
	if err != nil {
		u.logger.Error("Failed to create user", "error", err)
		return err
	}
	return nil
}

func (u *userService) GetUser(id uuid.UUID) (*domain.User, error) {
	userModel, err := u.userRepo.GetUser(id)
	if err != nil {
		u.logger.Error("Failed to get user", "error", err)
		return nil, err
	}
	user := userModel.ToDomain()
	return &user, nil
}
