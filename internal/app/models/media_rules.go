package models

import (
	"time"

	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

type MediaRules struct {
	ResourceType string
	ExitSize     []enums.ExitSize
	Roles        []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
