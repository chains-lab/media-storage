package models

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID           uuid.UUID `db:"id"`
	Ext          string    `db:"extension"`
	Size         int64     `db:"size"`
	ResourceID   uuid.UUID `db:"resource_id"`
	ResourceType string    `db:"resource_type"`
	URL          string    `db:"url"`
	OwnerID      uuid.UUID `db:"owner_id"`
	CreatedAt    time.Time `db:"created_at"`
}
