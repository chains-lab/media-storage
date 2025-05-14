package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/repo"
	"github.com/google/uuid"
)

type MediaModels struct {
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

type UploadMediaRequest struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
	UserID     uuid.UUID
	UserRole   roles.Role
	Resource   string
	ResourceID string
	Category   string
}

func (a App) UploadMedia(ctx context.Context, request UploadMediaRequest) (MediaModels, error) {
	createdAt := time.Now().UTC()

	ruleID := fmt.Sprintf("%s-%s", request.Resource, request.Category)

	rules, err := a.GetMediaRules(ctx, ruleID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaModels{}, ape.ErrMediaRulesNotFound
		default:
			return MediaModels{}, fmt.Errorf("get media rules in repo %s", err)
		}
	}

	allowedExtension := false
	for _, ext := range rules.Extensions {
		if ext == filepath.Ext(request.FileHeader.Filename) {
			allowedExtension = true
			break
		}
	}
	if !allowedExtension {
		return MediaModels{}, ape.ErrMediaExtensionNotAllowed
	}

	allowedSize := false
	if request.FileHeader.Size <= rules.MaxSize {
		allowedSize = true
	}
	if !allowedSize {
		return MediaModels{}, ape.ErrFileToLarge
	}

	allowedRole := false
	for _, role := range rules.AllowedRoles {
		if role == request.UserRole {
			allowedRole = true
			break
		}
	}
	if !allowedRole {
		return MediaModels{}, ape.ErrUserNotAllowedToUploadMedia
	}

	repoInput := repo.AddMediaInput{
		Filename:   request.FileHeader.Filename,
		Resource:   request.Resource,
		ResourceID: request.ResourceID,
		Category:   request.Category,
		OwnerID:    request.UserID,
		CreatedAt:  createdAt,
	}

	media, err := a.repoMedia.UploadMedia(ctx, request.File, repoInput)
	if err != nil {
		switch {
		default:
			return MediaModels{}, fmt.Errorf("add media in repo %s", err)
		}
	}

	return MediaModels{
		ID:         media.ID,
		Format:     media.Format,
		Extension:  media.Extension,
		Size:       media.Size,
		Url:        media.Url,
		Resource:   media.Resource,
		ResourceID: media.ResourceID,
		Category:   media.Category,
		OwnerID:    media.OwnerID,
		CreatedAt:  media.CreatedAt,
	}, nil
}

func (a App) GetMedia(ctx context.Context, mediaID uuid.UUID) (MediaModels, error) {
	media, err := a.repoMedia.GetMedia(ctx, mediaID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaModels{}, ape.ErrMediaNotFound
		default:
			return MediaModels{}, fmt.Errorf("get media: %w", err)
		}
	}

	return createMediaModel(media), nil
}

type DeleteMediaRequest struct {
	UserID      uuid.UUID
	UserRole    roles.Role
	MediaID     uuid.UUID
	InitiatorID uuid.UUID
}

func (a App) DeleteMedia(ctx context.Context, request DeleteMediaRequest) error {
	media, err := a.repoMedia.GetMedia(ctx, request.MediaID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaNotFound
		default:
			return fmt.Errorf("get media: %w", err)
		}
	}

	if media.OwnerID != request.UserID && roles.CompareRolesUser(request.UserRole, roles.Admin) < 0 {
		return ape.ErrUserNotAllowedToDeleteMedia
	}

	err = a.repoMedia.DeleteMedia(ctx, media.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaNotFound
		default:
			return fmt.Errorf("delete media: %w", err)
		}
	}

	return nil
}

func createMediaModel(media repo.MediaModel) MediaModels {
	return MediaModels{
		ID:         media.ID,
		Format:     media.Format,
		Extension:  media.Extension,
		Size:       media.Size,
		Url:        media.Url,
		Resource:   media.Resource,
		ResourceID: media.ResourceID,
		Category:   media.Category,
		OwnerID:    media.OwnerID,
		CreatedAt:  media.CreatedAt,
	}
}
