package models

import (
	"GoATTHStart/internal/domain"
	"time"

	uuid "github.com/google/uuid"
)

type UserCredentials struct {
	ID         uuid.UUID `gorm:"type:varchar(36);primary_key"`
	Email      string    `gorm:"type:varchar(255);not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	CreactedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  time.Time `gorm:"autoUpdateTime"`
}

func (uc *UserCredentials) ToDomain() domain.UserCredentials {
	return domain.UserCredentials{
		ID:         uc.ID,
		Email:      uc.Email,
		Password:   uc.Password,
		CreactedAt: uc.CreactedAt,
		UpdatedAt:  uc.UpdatedAt,
		DeletedAt:  uc.DeletedAt,
	}
}

func FromDomain(uc domain.UserCredentials) UserCredentials {
	return UserCredentials{
		ID:         uc.ID,
		Email:      uc.Email,
		Password:   uc.Password,
		CreactedAt: uc.CreactedAt,
		UpdatedAt:  uc.UpdatedAt,
		DeletedAt:  uc.DeletedAt,
	}
}
