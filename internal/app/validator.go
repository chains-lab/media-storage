package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

func (a App) validateCreate(
	mediaType enums.MediaType,
	size int64,
	filename string,
	resourceID uuid.UUID,
	role roles.Role,
) (string, error) {
	rule, err := a.repoRules.Get(context.TODO(), mediaType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", ape.ErrMediaRulesNotFound
		default:
			return "", fmt.Errorf("get media rules: %w", err)
		}
	}

	if size > rule.MaxSize {
		return "", ape.ErrFileToLarge
	}

	access := false
	ext := filepath.Ext(filename)
	for _, e := range rule.AllowedExits {
		if e == ext {
			access = true
			break
		}
	}

	if !access {
		return "", ape.ErrFileFormatNotAllowed
	}

	access = false
	for _, r := range rule.Roles {
		if r == role {
			access = true
			break
		}
	}

	if !access {
		return "", ape.ErrUserNotAllowedToUploadMedia
	}

	folder := fmt.Sprintf("%s/%s", rule.Folder, resourceID.String())
	return folder, nil
}

func (a App) validateDelete(ownerID, initiatorID uuid.UUID, role roles.Role) error {
	if ownerID != initiatorID && roles.CompareRolesUser(role, roles.Admin) < 0 {
		return ape.ErrUserNotAllowedToDeleteMedia
	}
	return nil
}
