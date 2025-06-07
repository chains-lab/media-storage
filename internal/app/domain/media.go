package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/repo"
	"github.com/google/uuid"
)

type repoMedia interface {
	UploadMedia(ctx context.Context, reader io.Reader, input repo.AddMediaInput) (repo.MediaModel, error)
	GetMedia(ctx context.Context, mediaID uuid.UUID) (repo.MediaModel, error)
	DeleteMedia(ctx context.Context, mediaID uuid.UUID) error
	DeleteFilesByResourceAndCategory(ctx context.Context, resource, category string) error
}

type Media struct {
	repo repoMedia
}

func NewMedia(cfg config.Config) (Media, error) {
	mediaRepo, err := repo.NewMedia(cfg)
	if err != nil {
		return Media{}, fmt.Errorf("create media repo: %w", err)
	}

	return Media{
		repo: mediaRepo,
	}, nil
}

type UploadMediaRequest struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
	UserID     uuid.UUID
	Resource   string
	ResourceID string
	Category   string
}

func (m Media) Upload(ctx context.Context, request UploadMediaRequest) (models.Media, *ape.Error) {
	createdAt := time.Now().UTC()

	repoInput := repo.AddMediaInput{
		Filename:   request.FileHeader.Filename,
		Resource:   request.Resource,
		ResourceID: request.ResourceID,
		Category:   request.Category,
		OwnerID:    request.UserID,
		CreatedAt:  createdAt,
	}

	media, err := m.repo.UploadMedia(ctx, request.File, repoInput)
	if err != nil {
		switch {
		default:
			return models.Media{}, ape.ErrorInternal(fmt.Errorf("upload media: %w", err))
		}
	}

	return createMediaModel(media), nil
}

func (m Media) Get(ctx context.Context, mediaID uuid.UUID) (models.Media, *ape.Error) {
	media, err := m.repo.GetMedia(ctx, mediaID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.Media{}, ape.ErrorMediaNotFound(err)
		default:
			return models.Media{}, ape.ErrorInternal(err)
		}
	}

	return createMediaModel(media), nil
}

type DeleteMediaRequest struct {
	MediaID       uuid.UUID
	InitiatorRole roles.Role
	InitiatorID   uuid.UUID
}

func (m Media) Delete(ctx context.Context, request DeleteMediaRequest) *ape.Error {
	media, err := m.repo.GetMedia(ctx, request.MediaID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrorMediaNotFound(err)
		default:
			return ape.ErrorInternal(err)
		}
	}

	if media.OwnerID != request.InitiatorID && roles.CompareRolesUser(request.InitiatorRole, roles.Admin) < 0 {
		return ape.ErrorUserNotAllowedToDeleteMedia(err)
	}

	err = m.repo.DeleteMedia(ctx, media.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrorMediaNotFound(err)
		default:
			return ape.ErrorInternal(err)
		}
	}

	return nil
}

func createMediaModel(media repo.MediaModel) models.Media {
	return models.Media{
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
