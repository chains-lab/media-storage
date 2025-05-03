package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/requests"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/media-storage/resources"
	"github.com/hs-zavet/tokens/roles"
)

func (h *Handler) CreateMediaRules(w http.ResponseWriter, r *http.Request) {
	req, err := requests.CreateMediaRules(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resourceType := chi.URLParam(r, "media_resource_type")

	if resourceType != req.Data.Id {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"media_resource_type": validation.NewError("media_resource_type", "invalid media resource id"),
		})...)
		return
	}

	curRoles, err := parseRoles(req.Data.Attributes.Roles)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"roles": validation.NewError("roles", "invalid role"),
		})...)
		return
	}

	res, err := h.app.CreateMediaRules(r.Context(), resourceType, app.CreateMediaRulesRequest{
		ExtSize: parseExtSize(req.Data.Attributes.ExitSize),
		Roles:   curRoles,
	})
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesAlreadyExists):
			httpkit.RenderErr(w, problems.Forbidden("media rules already exists"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Errorf("error create media rule %s", resourceType)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}

func parseRoles(r []string) ([]roles.Role, error) {
	parsedRoles := make([]roles.Role, 0, len(r))
	for i, role := range r {
		parsedRole, err := roles.ParseRole(role)
		if err != nil {
			return nil, err
		}
		parsedRoles[i] = parsedRole
	}
	return parsedRoles, nil
}

func parseExtSize(extSize []resources.ExitSizeInner) []enums.ExitSize {
	parsedExtSize := make([]enums.ExitSize, 0, len(extSize))
	for i, el := range extSize {
		parsedSize := enums.ExitSize{
			Size: el.Size,
			Exit: el.Exit,
		}
		parsedExtSize[i] = parsedSize
	}
	return parsedExtSize
}
