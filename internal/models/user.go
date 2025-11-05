package models

import (
	"GoATTHStart/internal/domain"
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:varchar(36);primary_key"`
	CredID     uuid.UUID `gorm: foreignKey:CredID;type:varchar(36);not null"`
	FirstName  string    `gorm:"type:varchar(255);not null"`
	LastName   string    `gorm:"type:varchar(255);not null"`
	Email      string    `gorm:"type:varchar(255);not null"`
	CreactedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  time.Time `gorm:"autoUpdateTime"`
}

func FromDomainUser(u domain.User) User {
	return User{
		ID:         u.ID,
		CredID:     u.CredID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		CreactedAt: u.CreactedAt,
		UpdatedAt:  u.UpdatedAt,
		DeletedAt:  u.DeletedAt,
	}
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:         u.ID,
		CredID:     u.CredID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
		CreactedAt: u.CreactedAt,
		UpdatedAt:  u.UpdatedAt,
		DeletedAt:  u.DeletedAt,
	}
}
