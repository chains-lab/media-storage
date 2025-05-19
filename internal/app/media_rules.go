package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/repo"
)

type MediaRulesModel struct {
	ID           string
	Extensions   []string
	MaxSize      int64
	AllowedRoles []roles.Role
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type CreateMediaRulesRequest struct {
	ID         string
	Extensions []string
	MaxSize    int64
	Roles      []roles.Role
}

func (a App) CreateMediaRules(ctx context.Context, request CreateMediaRulesRequest) (MediaRulesModel, error) {
	now := time.Now().UTC()

	_, err := a.GetMediaRules(context.TODO(), request.ID)
	if !errors.Is(err, ErrMediaRulesNotFound) {
		if err == nil {
			return MediaRulesModel{}, ErrMediaRulesAlreadyExists
		}
		return MediaRulesModel{}, err
	}

	repoInput := repo.CreateMediaRulesInput{
		ID:           request.ID,
		Extensions:   request.Extensions,
		MaxSize:      request.MaxSize,
		AllowedRoles: request.Roles,
		UpdatedAt:    now,
		CreatedAt:    now,
	}

	res, err := a.repoRules.Create(ctx, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaRulesModel{}, ErrMediaRulesAlreadyExists
		default:
			return MediaRulesModel{}, fmt.Errorf("add media rules in repo %s", err)
		}
	}

	return createMediaRulesModel(res)
}

func (a App) GetMediaRules(ctx context.Context, ID string) (MediaRulesModel, error) {
	rules, err := a.repoRules.Get(ctx, ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaRulesModel{}, ErrMediaRulesNotFound
		default:
			return MediaRulesModel{}, fmt.Errorf("get media rules: %w", err)
		}
	}

	return createMediaRulesModel(rules)
}

type UpdateMediaRulesRequest struct {
	Extensions   *[]string     `db:"extensions"`
	MaxSize      *int64        `db:"max_size"`
	AllowedRoles *[]roles.Role `db:"allowed_roles"`
}

func (a App) UpdateMediaRules(ctx context.Context, ID string, request UpdateMediaRulesRequest) (MediaRulesModel, error) {
	_, err := a.GetMediaRules(context.TODO(), ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaRulesModel{}, ErrMediaRulesNotFound
		default:
			return MediaRulesModel{}, fmt.Errorf("get media rules: %w", err)
		}
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
		return a.GetMediaRules(ctx, ID)
	}

	err = a.repoRules.Update(ctx, ID, repoInput)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaRulesModel{}, ErrMediaRulesNotFound
		default:
			return MediaRulesModel{}, fmt.Errorf("update media rules: %w", err)
		}
	}

	rules, err := a.repoRules.Get(ctx, ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return MediaRulesModel{}, ErrMediaRulesNotFound
		default:
			return MediaRulesModel{}, fmt.Errorf("get media rules: %w", err)
		}
	}

	return createMediaRulesModel(rules)
}

func (a App) DeleteMediaRules(ctx context.Context, ID string) error {
	_, err := a.GetMediaRules(context.TODO(), ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrMediaRulesNotFound
		default:
			return fmt.Errorf("get media rules: %w", err)
		}
	}

	arr := strings.Split(ID, "-")
	if len(arr) != 2 {
		return fmt.Errorf("invalid media id: %s", ID)
	}
	resource := arr[0]
	category := arr[1]

	err = a.repoMedia.DeleteFilesByResourceAndCategory(ctx, resource, category)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrMediaNotFound
		default:
			return fmt.Errorf("delete media: %w", err)
		}
	}

	err = a.repoRules.Delete(ctx, ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrMediaRulesNotFound
		default:
			return fmt.Errorf("delete media rules: %w", err)
		}
	}

	return nil
}

func createMediaRulesModel(mediaRules repo.MediaRulesModel) (MediaRulesModel, error) {
	return MediaRulesModel{
		ID:           mediaRules.ID,
		Extensions:   mediaRules.Extensions,
		MaxSize:      mediaRules.MaxSize,
		AllowedRoles: mediaRules.AllowedRoles,
		UpdatedAt:    mediaRules.UpdatedAt,
		CreatedAt:    mediaRules.CreatedAt,
	}, nil
}
