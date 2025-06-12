package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/rest/requests"
	"github.com/chains-lab/media-storage/internal/rest/responses"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const UpdateMediaRulesHandlerName = "UpdateMediaRulesHandler"

func (h Handler) UpdateMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	ruleID := chi.URLParam(r, "resource")

	req, err := requests.UpdateMediaRules(r)
	if err != nil {
		h.presenter.InvalidPointer(w, requestID, err)
		return
	}

	if ruleID != req.Data.Id {
		h.presenter.MismatchIdentification(w, requestID, "resource", req.Data.Id)
		return
	}

	var updateReq app.UpdateMediaRulesRequest
	if req.Data.Attributes.Roles != nil {
		curRoles, err := parseRoles(req.Data.Attributes.Roles)
		if err != nil {
			h.presenter.InvalidPointer(w, requestID, err)
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

	res, appErr := h.app.UpdateMediaRules(r.Context(), ruleID, updateReq)
	if appErr != nil {
		h.presenter.AppError(w, requestID, appErr)
		return
	}

	h.log.WithField("request_id", requestID).Infof("Media rule %s updated successfully %s", ruleID, requestID.String())
	httpkit.Render(w, responses.MediaRules(res))
}
