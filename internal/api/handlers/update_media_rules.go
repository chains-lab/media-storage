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
	ruleID := chi.URLParam(r, "resource-category")

	req, err := requests.UpdateMediaRules(r)
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
		updateReq.AllowedRoles = &curRoles
	}
	if req.Data.Attributes.Extensions != nil {
		updateReq.Extensions = &req.Data.Attributes.Extensions
	}
	if req.Data.Attributes.MaxSize != nil {
		updateReq.MaxSize = req.Data.Attributes.MaxSize
	}

	res, err := h.app.UpdateMediaRules(r.Context(), ruleID, updateReq)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			httpkit.RenderErr(w, problems.NotFound("media resource not found"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Errorf("error updating media rule %s", ruleID)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}
