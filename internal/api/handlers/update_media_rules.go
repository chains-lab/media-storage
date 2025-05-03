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
)

func (h *Handler) UpdateMediaRules(w http.ResponseWriter, r *http.Request) {
	resourceType := chi.URLParam(r, "resource_type")

	req, err := requests.UpdateMediaRules(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if resourceType != req.Data.Id {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"resource_type": validation.NewError("resource_type", "invalid media resource id"),
		})...)
		return
	}

	var updateReq app.UpdateMediaRulesRequest
	if req.Data.Attributes.Roles != nil {
		curRoles, err := parseRoles(req.Data.Attributes.Roles)
		if err != nil {
			h.log.WithError(err).Warn("error parsing request")
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"roles": validation.NewError("roles", "invalid role"),
			})...)
			return
		}
		updateReq.Roles = &curRoles
	}
	if req.Data.Attributes.ExitSize != nil {
		extSize := parseExtSize(req.Data.Attributes.ExitSize)
		updateReq.ExtSize = &extSize
	}

	res, err := h.app.UpdateMediaRules(r.Context(), resourceType, updateReq)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			httpkit.RenderErr(w, problems.NotFound("media resource not found"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Errorf("error updating media rule %s", resourceType)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}
