package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserCredentials struct {
	ID         uuid.UUID
	Email      string
	Password   string
	CreactedAt time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
