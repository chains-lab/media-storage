package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/tokens/roles"
)

func (a App) validateCreate(resourceType string, size int64, filename string, role roles.Role) error {
	rule, err := a.repoRules.Get(context.TODO(), resourceType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaRulesNotFound
		default:
			return fmt.Errorf("get media rules: %w", err)
		}
	}

	access := false
	ext := filepath.Ext(filename)
	for _, e := range rule.ExitSize {
		if e.Exit == ext && e.Size >= size {
			access = true
			break
		}
	}

	if !access {
		return ape.ErrFileFormatNotAllowed
	}

	access = false
	for _, r := range rule.Roles {
		if r == role {
			access = true
			break
		}
	}

	if !access {
		return ape.ErrUserNotAllowedToUploadMedia
	}

	return nil
}

func (a App) validateDelete(ownerID, initiatorID uuid.UUID, role roles.Role) error {
	if ownerID != initiatorID && roles.CompareRolesUser(role, roles.Admin) < 0 {
		return ape.ErrUserNotAllowedToDeleteMedia
	}
	return nil
}
