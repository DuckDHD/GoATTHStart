package repositories

import (
	"GoATTHStart/internal/models"
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

func NewUserRepository(gormDB *gorm.DB, sqlDB *sql.DB) *userRepository {
	return &userRepository{
		gormDB: gormDB,
		sqlDB:  sqlDB,
	}
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

func (ur *userRepository) CreateUser(user *models.User) error {
	transaction := ur.gormDB.Begin()
	err := transaction.Create(user).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit().Error
}

func (ur *userRepository) GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := ur.gormDB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(user *models.User) error {
	transaction := ur.gormDB.Begin()
	err := transaction.Save(user).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit().Error
}

func (ur *userRepository) DeleteUser(user *models.User) error {
	transaction := ur.gormDB.Begin()
	err := transaction.Delete(user).Error
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit().Error
}
