package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/internal/config"
	"github.com/chains-lab/media-storage/internal/repo"
)

type mediaRulesRepo interface {
	Create(ctx context.Context, input repo.CreateMediaRulesInput) (repo.MediaRulesModel, error)
	Get(ctx context.Context, id string) (repo.MediaRulesModel, error)
	Update(ctx context.Context, id string, input repo.MediaRulesUpdateInput) error
	Delete(ctx context.Context, id string) error
}

func NewMediaRules(cfg config.Config) (MediaRules, error) {
	rulesRepo, err := repo.NewMediaRulesRepo(cfg)
	if err != nil {
		return MediaRules{}, err
	}

	return MediaRules{
		repo: rulesRepo,
	}, nil
}

type MediaRules struct {
	repo mediaRulesRepo
}

func (r MediaRules) Get(ctx context.Context, ID string) (models.MediaRules, *ape.Error) {
	rules, err := r.repo.Get(ctx, ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrorMediaRulesNotFound(err)
		default:
			return models.MediaRules{}, ape.ErrorInternal(err)
		}
	}

	return createMediaRulesModel(rules), nil
}

type CreateMediaRulesRequest struct {
	ID         string
	Extensions []string
	MaxSize    int64
	Roles      []roles.Role
}

func (r MediaRules) Create(ctx context.Context, request CreateMediaRulesRequest) (models.MediaRules, *ape.Error) {
	now := time.Now().UTC()

	_, appErr := r.Get(context.TODO(), request.ID)
	if appErr != nil {
		return models.MediaRules{}, appErr
	}

	repoInput := repo.CreateMediaRulesInput{
		ID:           request.ID,
		Extensions:   request.Extensions,
		MaxSize:      request.MaxSize,
		AllowedRoles: request.Roles,
		UpdatedAt:    now,
		CreatedAt:    now,
	}

	res, err := r.repo.Create(ctx, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrorMediaRulesAlreadyExists(err)
		default:
			return models.MediaRules{}, ape.ErrorInternal(err)
		}
	}

	return createMediaRulesModel(res), nil
}

type UpdateMediaRulesRequest struct {
	Extensions   *[]string     `db:"extensions"`
	MaxSize      *int64        `db:"max_size"`
	AllowedRoles *[]roles.Role `db:"allowed_roles"`
}

func (r MediaRules) Update(ctx context.Context, ID string, request UpdateMediaRulesRequest) (models.MediaRules, *ape.Error) {
	_, appErr := r.Get(ctx, ID)
	if appErr != nil {
		return models.MediaRules{}, appErr
	}

	now := time.Now().UTC()
	updated := false

	var repoInput repo.MediaRulesUpdateInput
	if request.Extensions != nil {
		repoInput.Extensions = request.Extensions
		updated = true
	}
	if request.MaxSize != nil {
		repoInput.MaxSize = request.MaxSize
		updated = true
	}
	if request.AllowedRoles != nil {
		repoInput.AllowedRoles = request.AllowedRoles
		updated = true
	}
	repoInput.UpdatedAt = now

	//for idempotency
	if !updated {
		return r.Get(ctx, ID)
	}

	err := r.repo.Update(ctx, ID, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.MediaRules{}, ape.ErrorMediaRulesNotFound(err)
		default:
			return models.MediaRules{}, ape.ErrorInternal(err)
		}
	}

	return r.Get(ctx, ID)
}

func (r MediaRules) Delete(ctx context.Context, ID string) *ape.Error {
	_, appErr := r.Get(ctx, ID)
	if appErr != nil {
		return appErr
	}

	arr := strings.Split(ID, "-")
	if len(arr) != 2 {
		return ape.ErrorInvalidRequestQuery(fmt.Errorf("invalid media rules ID format: %s", ID))
	}
	//resource := arr[0]
	//category := arr[1]

	//err := r.repo.DeleteFilesByResourceAndCategory(ctx, resource, category)
	//if err != nil {
	//	switch {
	//	case errors.Is(err, sql.ErrNoRows):
	//		return ErrMediaNotFound
	//	default:
	//		return fmt.Errorf("delete media: %w", err)
	//	}
	//}

	err := r.repo.Delete(ctx, ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ape.ErrorMediaRulesNotFound(err)
		default:
			return ape.ErrorInternal(err)
		}
	}

	return nil
}

func createMediaRulesModel(mediaRules repo.MediaRulesModel) models.MediaRules {
	return models.MediaRules{
		ID:           mediaRules.ID,
		Extensions:   mediaRules.Extensions,
		MaxSize:      mediaRules.MaxSize,
		AllowedRoles: mediaRules.AllowedRoles,
		UpdatedAt:    mediaRules.UpdatedAt,
		CreatedAt:    mediaRules.CreatedAt,
	}
}
