package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/enums"
)

type Media struct {
	ID     uuid.UUID `db:"id"`
	Folder string    `db:"folder"`
	Ext    string    `db:"extension"`
	Size   int64     `db:"size"`
	URL    string    `db:"url"`

	//Name of resource
	ResourceType enums.ResourceType `db:"resource_type"`
	//Filename of resource
	ResourceID uuid.UUID `db:"resource_id"`
	//MediaType of resource
	MediaType enums.MediaType `db:"media_type"`

	//Owner ID of resource who
	OwnerID   uuid.UUID `db:"owner_id"`
	CreatedAt time.Time `db:"created_at"`
}
