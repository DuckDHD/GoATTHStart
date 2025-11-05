package services

import (
	"GoATTHStart/internal/domain"
	"GoATTHStart/internal/models"
	"GoATTHStart/internal/repositories"
	"GoATTHStart/internal/toolbox"
	"log/slog"
)

type userCredentialsService struct {
	userCredRepo repositories.UserCredentialsRepository
	logger       *slog.Logger
}

type UserCredentialsService interface {
	CreateUserCredentials(userCredentials *domain.UserCredentials) error
	GetUserCredentials(email string) (*domain.UserCredentials, error)
	UpdateUserCredentials(userCredentials *domain.UserCredentials) error
	DeleteUserCredentials(userCredentials *domain.UserCredentials) error
}

func NewUserCredentialsService(userCredRepo repositories.UserCredentialsRepository, logger *slog.Logger) *userCredentialsService {
	return &userCredentialsService{
		userCredRepo: userCredRepo,
		logger:       logger,
	}
}

func (ucs *userCredentialsService) CreateUserCredentials(userCredentials *domain.UserCredentials) error {
	if userCredentials == nil {
		return nil
	}

	if userCredentials.Email == "" {
		return nil
	}

	if userCredentials.Password == "" {
		return nil
	}

	hashedPassword, err := toolbox.HashPassword(userCredentials.Password)
	if err != nil {
		return err
	}
	userCredentials.Password = hashedPassword

	_, err = ucs.userCredRepo.GetUserCredentials(userCredentials.Email)
	if err == nil {
		ucs.logger.Info("User already exists")
		return domain.ErrUserCredAalreadyExists
	}

	userCredModel := models.FromDomainUC(*userCredentials)

	err = ucs.userCredRepo.CreateUserCredentials(&userCredModel)
	if err != nil {
		ucs.logger.Error("Failed to create user credentials", "error", err)
		return domain.ErrFailedTocreateCreds
	}

	return nil
}
