package app

import (
	"context"

	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/chains-lab/media-storage/internal/app/domain"
	"github.com/chains-lab/media-storage/internal/app/models"
)

type CreateMediaRulesRequest struct {
	ID         string
	Extensions []string
	MaxSize    int64
	Roles      []roles.Role
}

func (a App) CreateMediaRules(ctx context.Context, request CreateMediaRulesRequest) (models.MediaRules, *ape.Error) {
	return a.rules.Create(ctx, domain.CreateMediaRulesRequest{
		ID:         request.ID,
		Extensions: request.Extensions,
		MaxSize:    request.MaxSize,
		Roles:      request.Roles,
	})
}

func (a App) GetMediaRules(ctx context.Context, ID string) (models.MediaRules, *ape.Error) {
	return a.rules.Get(ctx, ID)
}

type UpdateMediaRulesRequest struct {
	Extensions   *[]string     `db:"extensions"`
	MaxSize      *int64        `db:"max_size"`
	AllowedRoles *[]roles.Role `db:"allowed_roles"`
}

func (a App) UpdateMediaRules(ctx context.Context, ID string, request UpdateMediaRulesRequest) (models.MediaRules, *ape.Error) {
	return a.rules.Update(ctx, ID, domain.UpdateMediaRulesRequest{
		Extensions:   request.Extensions,
		MaxSize:      request.MaxSize,
		AllowedRoles: request.AllowedRoles,
	})
}

func (a App) DeleteMediaRules(ctx context.Context, ID string) *ape.Error {
	return a.rules.Delete(ctx, ID)
}
