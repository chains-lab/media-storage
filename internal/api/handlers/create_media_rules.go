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
	"github.com/hs-zavet/tokens/roles"
)

func (h *Handler) CreateMediaRules(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "resource-category")

	req, err := requests.CreateMediaRules(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if ruleID != req.Data.Id {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"resource-category": validation.NewError("resource-category", "invalid resource category"),
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

	res, err := h.app.CreateMediaRules(r.Context(), app.CreateMediaRulesRequest{
		ID:         ruleID,
		Extensions: req.Data.Attributes.Extensions,
		MaxSize:    req.Data.Attributes.MaxSize,
		Roles:      curRoles,
	})
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesAlreadyExists):
			httpkit.RenderErr(w, problems.Forbidden("media rules already exists"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Errorf("error create media rule %s", err)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}

func parseRoles(r []string) ([]roles.Role, error) {
	parsedRoles := make([]roles.Role, len(r)) // length = len(r)
	for i, str := range r {
		pr, err := roles.ParseRole(str)
		if err != nil {
			return nil, err
		}
		parsedRoles[i] = pr
	}
	return parsedRoles, nil
}
