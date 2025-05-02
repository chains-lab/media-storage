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
}

type MediaRulesRepo interface {
	Create(ctx context.Context, input repo.CreateMediaRulesInput) (repo.MediaRulesModel, error)
	Get(ctx context.Context, mType enums.MediaType) (repo.MediaRulesModel, error)
	Update(ctx context.Context, input repo.MediaRulesUpdateInput) error
	Delete(ctx context.Context, mType enums.MediaType) error
}

type App struct {
	repoMedia repoMedia
	repoRules MediaRulesRepo
	//mediaValidator validator.Validator
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
	ResourceType enums.ResourceType
	ResourceID   uuid.UUID
	MediaType    enums.MediaType
	User         tokens.AccountData
	File         io.Reader
	FileHeader   *multipart.FileHeader
}

func (a App) UploadMedia(ctx context.Context, request UploadMediaRequest) (models.Media, error) {
	fileID := uuid.New()
	createdAt := time.Now().UTC()

	folder, err := a.validateCreate(
		request.MediaType,
		request.FileHeader.Size,
		request.FileHeader.Filename,
		request.ResourceID,
		request.User.Role,
	)
	if err != nil {
		return models.Media{}, err
	}

	repoInput := repo.AddMediaInput{
		Folder:       folder,
		Filename:     fileID,
		Ext:          filepath.Ext(request.FileHeader.Filename),
		ResourceType: request.ResourceType,
		ResourceID:   request.ResourceID,
		MediaType:    request.MediaType,
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
		Folder:       media.Folder,
		Ext:          media.Ext,
		ResourceType: media.ResourceType,
		ResourceID:   media.ResourceID,
		MediaType:    media.MediaType,
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
	Role        roles.Role
	ResourceID  uuid.UUID
	InitiatorID uuid.UUID
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

	if err = a.validateDelete(media.OwnerID, request.InitiatorID, request.Role); err != nil {
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

type AddMediaRulesRequest struct {
	MediaType    enums.MediaType
	MaxSize      int64
	AllowedExits []string
	Folder       string
	Roles        []roles.Role
}

func (a App) AddMediaRules(ctx context.Context, request AddMediaRulesRequest) (models.MediaRules, error) {
	repoInput := repo.CreateMediaRulesInput{
		MediaType:    request.MediaType,
		MaxSize:      request.MaxSize,
		AllowedExits: request.AllowedExits,
		Folder:       request.Folder,
		Roles:        request.Roles,
	}

	rules, err := a.repoRules.Create(ctx, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesAlreadyExists
		default:
			return models.MediaRules{}, fmt.Errorf("add media rules in repo %s", err)
		}
	}

	return models.MediaRules{
		MediaType:    rules.MediaType,
		MaxSize:      rules.MaxSize,
		AllowedExits: rules.AllowedExits,
		Folder:       rules.Folder,
		Roles:        rules.Roles,
	}, nil
}

func (a App) GetMediaRules(ctx context.Context, mediaType enums.MediaType) (models.MediaRules, error) {
	rules, err := a.repoRules.Get(context.TODO(), mediaType)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesNotFound
		default:
			return models.MediaRules{}, fmt.Errorf("get media rules: %w", err)
		}
	}

	return models.MediaRules{
		MediaType:    rules.MediaType,
		MaxSize:      rules.MaxSize,
		AllowedExits: rules.AllowedExits,
		Folder:       rules.Folder,
		Roles:        rules.Roles,
	}, nil
}

type UpdateMediaRulesRequest struct {
	MediaType    enums.MediaType
	MaxSize      int64
	AllowedExits []string
	Folder       string
	Roles        []roles.Role
}

func (a App) UpdateMediaRules(ctx context.Context, mType enums.MediaType, request UpdateMediaRulesRequest) (models.MediaRules, error) {
	repoInput := repo.MediaRulesUpdateInput{
		MaxSize:      &request.MaxSize,
		AllowedExits: &request.AllowedExits,
		Folder:       &request.Folder,
		Roles:        &request.Roles,
	}

	err := a.repoRules.Update(ctx, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrMediaRulesNotFound
		default:
			return models.MediaRules{}, fmt.Errorf("update media rules: %w", err)
		}
	}

	return models.MediaRules{
		MaxSize:      request.MaxSize,
		AllowedExits: request.AllowedExits,
		Folder:       request.Folder,
		Roles:        request.Roles,
	}, nil
}

func (a App) DeleteMediaRules(ctx context.Context, mType enums.MediaType) error {
	err := a.repoRules.Delete(ctx, mType)
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
		Folder:       media.Folder,
		Ext:          media.Ext,
		Size:         media.Size,
		URL:          media.URL,
		OwnerID:      media.OwnerID,
		ResourceType: media.ResourceType,
		ResourceID:   media.ResourceID,
		MediaType:    media.MediaType,
		CreatedAt:    media.CreatedAt,
	}
}
