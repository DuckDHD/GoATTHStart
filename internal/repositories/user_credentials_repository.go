package repositories

import (
	"GoATTHStart/internal/models"
	"database/sql"

	"gorm.io/gorm"
)

type userCredentialsRepository struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

type UserCredentialsRepository interface {
	CreateUserCredentials(userCredentials *models.UserCredentials) error
	GetUserCredentials(email string) (*models.UserCredentials, error)
	UpdateUserCredentials(userCredentials *models.UserCredentials) error
}

func NewUserCredentialsRepository(gormDB *gorm.DB, sqlDB *sql.DB, mode string) *userCredentialsRepository {
	return &userCredentialsRepository{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

func (ucr *userCredentialsRepository) CreateUserCredentials(userCredentials *models.UserCredentials) error {
	transaction := ucr.gormDB.Begin()
	err := transaction.Create(userCredentials).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit().Error
}

func (ucr *userCredentialsRepository) GetUserCredentials(email string) (*models.UserCredentials, error) {
	var userCredentials models.UserCredentials
	err := ucr.gormDB.Where("email = ?", email).First(&userCredentials).Error
	if err != nil {
		return nil, err
	}

	return &userCredentials, nil
}

func (ucr *userCredentialsRepository) UpdateUserCredentials(userCredentials *models.UserCredentials) error {
	transaction := ucr.gormDB.Begin()
	err := transaction.Save(userCredentials).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit().Error
}

func (ucr *userCredentialsRepository) DeleteUserCredentials(userCredentials *models.UserCredentials) error {
	transaction := ucr.gormDB.Begin()
	err := transaction.Delete(userCredentials).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit().Error
}
