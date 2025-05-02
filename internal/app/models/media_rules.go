package models

import (
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

type MediaRules struct {
	MediaType    enums.MediaType `db:"media_type"`
	MaxSize      int64           `db:"max_size"`
	AllowedExits []string        `db:"allowed_exits"`
	Folder       string          `db:"folder"`
	Roles        []roles.Role    `db:"roles_access_update"`
}
