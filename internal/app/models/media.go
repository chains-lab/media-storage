package models

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID         uuid.UUID
	Format     string
	Extension  string
	Size       int64
	Url        string
	Resource   string
	ResourceID string
	Category   string
	OwnerID    uuid.UUID
	CreatedAt  time.Time
}
