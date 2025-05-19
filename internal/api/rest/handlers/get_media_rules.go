package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/rest/ape"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const GetMediaRulesHandlerName = "GetMediaRulesHandler"

func (h Handler) GetMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    GetMediaRulesHandlerName,
		"request_id": requestID.String(),
	})

	ruleID := chi.URLParam(r, "resource")

	res, err := h.app.GetMediaRules(r.Context(), ruleID)
	if err != nil {
		ape.ApplicationError(w, requestID, err)
		log.WithError(err).Errorf("Error getting media rules %s", ruleID)

		return
	}

	log.Debugf("Media rules %s retrieved successfully %s", ruleID, requestID.String())
	httpkit.Render(w, responses.MediaRules(res))
}
