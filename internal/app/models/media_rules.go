package models

import (
	"time"

	"github.com/chains-lab/gatekit/roles"
)

type MediaRules struct {
	ID           string
	Extensions   []string
	MaxSize      int64
	AllowedRoles []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
