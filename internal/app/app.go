package app

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/app/models"
	"github.com/hs-zavet/media-storage/internal/app/validator_media"
	"github.com/hs-zavet/media-storage/internal/config"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/media-storage/internal/repo"
	"github.com/hs-zavet/tokens"
	"github.com/sirupsen/logrus"
)

type repoMedia interface {
	AddMedia(ctx context.Context, reader io.Reader, input repo.AddMediaInput) (repo.MediaModel, error)
	ListMedia(ctx context.Context, folder string, limit, offset uint) ([]repo.MediaModel, error)
	DeleteMedia(ctx context.Context, folder string, id uuid.UUID, ext string) error
	DeleteFromFolder(ctx context.Context, folder string) error
}

type App struct {
	repo           repoMedia
	mediaValidator validator_media.Validator
}

func NewApp(cfg config.Config, log *logrus.Logger) (App, error) {
	data, err := repo.NewMedia(cfg)
	if err != nil {
		return App{}, err
	}

	mediaValidator := validator_media.NewValidator()

	return App{
		repo:           data,
		mediaValidator: mediaValidator,
	}, nil
}

type UploadMediaRequest struct {
	ResourceType enums.ResourceType
	ResourceID   uuid.UUID
	MediaType    enums.MediaType
	OwnerID      *uuid.UUID
}

func (a App) UploadMedia(ctx context.Context, user tokens.AccountData, file io.Reader, fileHeader *multipart.FileHeader, request UploadMediaRequest) (models.Media, error) {
	fileID := uuid.New()
	createdAt := time.Now().UTC()

	folder, err := a.mediaValidator.ValidateUpdate(request.MediaType, fileHeader, request.ResourceID, user.Role)
	if err != nil {
		return models.Media{}, fmt.Errorf("validate media: %w", err)
	}

	repoInput := repo.AddMediaInput{
		Folder:       folder,
		Filename:     fileID,
		Ext:          filepath.Ext(fileHeader.Filename),
		ResourceType: request.ResourceType,
		ResourceID:   request.ResourceID,
		MediaType:    request.MediaType,
		CreatedAt:    createdAt,
	}

	if request.OwnerID != nil {
		repoInput.OwnerID = request.OwnerID
	}

	media, err := a.repo.AddMedia(ctx, file, repoInput)
	if err != nil {
		return models.Media{}, fmt.Errorf("add media: %w", err)
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
		Size:         fileHeader.Size,
	}, nil
}

func createMediaModel(media repo.MediaModel) models.Media {
	res := models.Media{
		ID:           media.ID,
		Folder:       media.Folder,
		Ext:          media.Ext,
		Size:         media.Size,
		URL:          media.URL,
		ResourceType: media.ResourceType,
		ResourceID:   media.ResourceID,
		MediaType:    media.MediaType,
		CreatedAt:    media.CreatedAt,
	}

	if media.OwnerID != nil {
		res.OwnerID = media.OwnerID
	}

	return res
}
