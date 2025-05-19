package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/rest/ape"
	"github.com/chains-lab/media-storage/internal/api/rest/requests"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const UpdateMediaRulesHandlerName = "UpdateMediaRulesHandler"

func (h Handler) UpdateMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    UpdateMediaRulesHandlerName,
		"request_id": requestID.String(),
	})

	ruleID := chi.URLParam(r, "resource")

	req, err := requests.UpdateMediaRules(r)
	if err != nil {
		ape.BadRequest(w, requestID, "Error parsing request body")
		log.WithError(err).Errorf("Error parsing request body")

		return
	}

	if ruleID != req.Data.Id {
		ape.BadRequest(w, requestID, "Resource ID does not match the request ID")
		log.Errorf("Resource ID %s does not match the request ID %s", ruleID, req.Data.Id)

		return
	}

	var updateReq app.UpdateMediaRulesRequest
	if req.Data.Attributes.Roles != nil {
		curRoles, err := parseRoles(req.Data.Attributes.Roles)
		if err != nil {
			ape.BadRequest(w, requestID, "Error parsing roles")
			log.WithError(err).Errorf("Error parsing roles")

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
		ape.ApplicationError(w, requestID, err)
		return
	}

	log.Infof("Media rule %s updated successfully %s", ruleID, requestID.String())
	httpkit.Render(w, responses.MediaRules(res))
}
