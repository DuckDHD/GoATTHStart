package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	CredID     uuid.UUID
	FirstName  string
	LastName   string
	Email      string
	CreactedAt time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
