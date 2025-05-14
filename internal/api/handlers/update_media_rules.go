package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/requests"
	"github.com/chains-lab/media-storage/internal/api/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) UpdateMediaRules(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "resource-category")

	req, err := requests.UpdateMediaRules(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status: http.StatusBadRequest,
			Error:  err,
		})...)
		return
	}

	if ruleID != req.Data.Id {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status:   http.StatusBadRequest,
			Title:    "resource-category does not match id",
			Error:    fmt.Errorf("resource-category %s does not match id %s", ruleID, req.Data.Id),
			Parametr: "id",
			Pointer:  "/data/id",
		})...)
		return
	}

	var updateReq app.UpdateMediaRulesRequest
	if req.Data.Attributes.Roles != nil {
		curRoles, err := parseRoles(req.Data.Attributes.Roles)
		if err != nil {
			h.log.WithError(err).Warn("error parsing request")
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status:  http.StatusBadRequest,
				Title:   "error parsing roles",
				Error:   err,
				Pointer: "/data/attributes/roles",
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
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusNotFound,
				Title:  "Media rules not found",
			})...)
		default:
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusInternalServerError,
			})...)
		}
		h.log.WithError(err).Errorf("error updating media rule %s", ruleID)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}
