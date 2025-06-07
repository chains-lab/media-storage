package app

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/app/domain"
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/google/uuid"
)

type UploadMediaRequest struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
	UserID     uuid.UUID
	UserRole   roles.Role
	Resource   string
	ResourceID string
	Category   string
}

func (a App) UploadMedia(ctx context.Context, request UploadMediaRequest) (models.Media, *ape.Error) {
	ruleID := fmt.Sprintf("%s-%s", request.Resource, request.Category)

	rule, appErr := a.rules.Get(ctx, ruleID)
	if appErr != nil {
		return models.Media{}, appErr
	}

	allowedExtension := false
	for _, ext := range rule.Extensions {
		if ext == filepath.Ext(request.FileHeader.Filename) {
			allowedExtension = true
			break
		}
	}
	if !allowedExtension {
		return models.Media{}, ape.ErrorMediaExtensionNotAllowed(fmt.Errorf("extension %s not allowed for resource %s and category %s", filepath.Ext(request.FileHeader.Filename), request.Resource, request.Category))
	}

	allowedSize := false
	if request.FileHeader.Size <= rule.MaxSize {
		allowedSize = true
	}
	if !allowedSize {
		return models.Media{}, ape.ErrorFileTooLarge(fmt.Errorf("file size %d exceeds allowed size %d", request.FileHeader.Size, rule.MaxSize))
	}

	allowedRole := false
	for _, role := range rule.AllowedRoles {
		if role == request.UserRole {
			allowedRole = true
			break
		}
	}
	if !allowedRole {
		return models.Media{}, ape.ErrorUserNotAllowedToUploadMedia(fmt.Errorf("user role %s not allowed to upload media for resource %s and category %s", request.UserRole, request.Resource, request.Category))
	}

	repoInput := domain.UploadMediaRequest{
		FileHeader: request.FileHeader,
		File:       request.File,
		UserID:     request.UserID,
		Resource:   request.Resource,
		ResourceID: request.ResourceID,
		Category:   request.Category,
	}

	return a.media.Upload(ctx, repoInput)
}

func (a App) GetMedia(ctx context.Context, mediaID uuid.UUID) (models.Media, *ape.Error) {
	return a.media.Get(ctx, mediaID)
}

type DeleteMediaRequest struct {
	MediaID       uuid.UUID
	InitiatorRole roles.Role
	InitiatorID   uuid.UUID
}

func (a App) DeleteMedia(ctx context.Context, request DeleteMediaRequest) *ape.Error {
	return a.media.Delete(ctx, domain.DeleteMediaRequest{
		MediaID:       request.MediaID,
		InitiatorRole: request.InitiatorRole,
		InitiatorID:   request.InitiatorID,
	})
}
