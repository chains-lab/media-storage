package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/media-storage/internal/app/models"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/media-storage/internal/repo"
	"github.com/hs-zavet/tokens"
	"github.com/hs-zavet/tokens/roles"
)

type repoMedia interface {
	AddMedia(ctx context.Context, reader io.Reader, input repo.AddMediaInput) (repo.MediaModel, error)
	GetMedia(ctx context.Context, filename uuid.UUID) (repo.MediaModel, error)
	DeleteMedia(ctx context.Context, id uuid.UUID) error
	DeleteFilesByResourceType(ctx context.Context, prefix string) error
}

type MediaRulesRepo interface {
	Create(ctx context.Context, input repo.CreateMediaRulesInput) (repo.MediaRulesModel, error)
	Get(ctx context.Context, resourceType string) (repo.MediaRulesModel, error)
	Update(ctx context.Context, resourceType string, input repo.MediaRulesUpdateInput) error
	Delete(ctx context.Context, resourceType string) error
}

type App struct {
	repoMedia repoMedia
	repoRules MediaRulesRepo
}

func NewApp(cfg config.Config) (App, error) {
	mediaRepo, err := repo.NewMedia(cfg)
	if err != nil {
		return App{}, err
	}

	rulesRepo, err := repo.NewMediaRulesRepo(cfg)
	if err != nil {
		return App{}, err
	}

	return App{
		repoMedia: mediaRepo,
		repoRules: rulesRepo,
	}, nil
}

type UploadMediaRequest struct {
	User         tokens.AccountData
	FileHeader   *multipart.FileHeader
	File         multipart.File
	ResourceType string
	ResourceID   uuid.UUID
	ExitSize     []enums.ExitSize
	Roles        []roles.Role
	CreatedAt    time.Time
}

func (a App) UploadMedia(ctx context.Context, request UploadMediaRequest) (models.Media, error) {
	fileID := uuid.New()
	createdAt := time.Now().UTC()

	err := a.validateCreate(request.ResourceType, request.FileHeader.Size, request.FileHeader.Filename, request.User.Role)
	if err != nil {
		return models.Media{}, err
	}

	repoInput := repo.AddMediaInput{
		Filename:     fileID,
		Ext:          filepath.Ext(request.FileHeader.Filename),
		ResourceType: request.ResourceType,
		ResourceID:   request.ResourceID,
		OwnerID:      request.User.AccountID,
		CreatedAt:    createdAt,
	}

	media, err := a.repoMedia.AddMedia(ctx, request.File, repoInput)
	if err != nil {
		switch {
		default:
			return models.Media{}, fmt.Errorf("add media in repo %s", err)
		}
	}

	return models.Media{
		ID:           media.ID,
		Ext:          media.Ext,
		ResourceType: media.ResourceType,
		ResourceID:   media.ResourceID,
		CreatedAt:    media.CreatedAt,
		OwnerID:      media.OwnerID,
		Size:         request.FileHeader.Size,
	}, nil
}

func (a App) GetMedia(ctx context.Context, resourceID uuid.UUID) (models.Media, error) {
	media, err := a.repoMedia.GetMedia(ctx, resourceID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.Media{}, ape.ErrMediaNotFound
		default:
			return models.Media{}, fmt.Errorf("get media: %w", err)
		}
	}

	return createMediaModel(media), nil
}

type DeleteMediaRequest struct {
	User         tokens.AccountData
	ResourceType string
	ResourceID   uuid.UUID
	InitiatorID  uuid.UUID
}

func (a App) DeleteMedia(ctx context.Context, request DeleteMediaRequest) error {
	media, err := a.repoMedia.GetMedia(ctx, request.ResourceID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaNotFound
		default:
			return fmt.Errorf("get media: %w", err)
		}
	}

	if err = a.validateDelete(media.OwnerID, request.InitiatorID, request.User.Role); err != nil {
		return err
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

type CreateMediaRulesRequest struct {
	ExtSize []enums.ExitSize
	Roles   []roles.Role
}

func (a App) CreateMediaRules(ctx context.Context, resourceType string, request CreateMediaRulesRequest) (models.MediaRules, error) {
	now := time.Now().UTC()

	_, err := a.GetMediaRules(context.TODO(), resourceType)
	if !errors.Is(err, ape.ErrMediaRulesNotFound) {
		return models.MediaRules{}, ape.ErrMediaRulesAlreadyExists
	}

	repoInput := repo.CreateMediaRulesInput{
		ResourceType: resourceType,
		ExitSize:     request.ExtSize,
		Roles:        request.Roles,
		UpdatedAt:    now,
		CreatedAt:    now,
	}

	res, err := a.repoRules.Create(ctx, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesAlreadyExists
		default:
			return models.MediaRules{}, fmt.Errorf("add media rules in repo %s", err)
		}
	}

	return createMediaRulesModel(res), nil
}

func (a App) GetMediaRules(ctx context.Context, resourceType string) (models.MediaRules, error) {

	fmt.Printf("test \n")

	rules, err := a.repoRules.Get(context.TODO(), resourceType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesNotFound
		default:
			return models.MediaRules{}, fmt.Errorf("get media rules: %w", err)
		}
	}

	return createMediaRulesModel(rules), nil
}

type UpdateMediaRulesRequest struct {
	ExtSize *[]enums.ExitSize
	Roles   *[]roles.Role
}

func (a App) UpdateMediaRules(ctx context.Context, resourceType string, request UpdateMediaRulesRequest) (models.MediaRules, error) {
	now := time.Now().UTC()
	updated := false

	fmt.Printf("test \n")
	var repoInput repo.MediaRulesUpdateInput
	if request.ExtSize != nil {
		repoInput.ExitSize = request.ExtSize
		updated = true
	}
	if request.Roles != nil {
		repoInput.Roles = request.Roles
		updated = true
	}
	repoInput.UpdatedAt = now

	if !updated {
		return a.GetMediaRules(ctx, resourceType)
	}

	err := a.repoRules.Update(ctx, resourceType, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesNotFound
		default:
			return models.MediaRules{}, fmt.Errorf("update media rules: %w", err)
		}
	}

	rules, err := a.repoRules.Get(ctx, resourceType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesNotFound
		default:
			return models.MediaRules{}, fmt.Errorf("get media rules: %w", err)
		}
	}

	return createMediaRulesModel(rules), nil
}

func (a App) DeleteMediaRules(ctx context.Context, resourceType string) error {
	_, err := a.repoRules.Get(context.TODO(), resourceType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaRulesNotFound
		default:
			return fmt.Errorf("get media rules: %w", err)
		}
	}

	err = a.repoMedia.DeleteFilesByResourceType(ctx, resourceType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaNotFound
		default:
			return fmt.Errorf("delete media: %w", err)
		}
	}

	err = a.repoRules.Delete(ctx, resourceType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrMediaRulesNotFound
		default:
			return fmt.Errorf("delete media rules: %w", err)
		}
	}

	return nil
}

func createMediaModel(media repo.MediaModel) models.Media {
	return models.Media{
		ID:           media.ID,
		Ext:          media.Ext,
		Size:         media.Size,
		URL:          media.URL,
		OwnerID:      media.OwnerID,
		ResourceType: media.ResourceType,
		ResourceID:   media.ResourceID,
		CreatedAt:    media.CreatedAt,
	}
}

func createMediaRulesModel(mediaRules repo.MediaRulesModel) models.MediaRules {
	return models.MediaRules{
		ResourceType: mediaRules.ResourceType,
		ExitSize:     mediaRules.ExitSize,
		Roles:        mediaRules.Roles,
		CreatedAt:    mediaRules.CreatedAt,
		UpdatedAt:    mediaRules.UpdatedAt,
	}
}
