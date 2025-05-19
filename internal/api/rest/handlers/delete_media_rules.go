package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/rest/ape"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const DeleteMediaRulesHandlerName = "DeleteMediaRulesHandler"

func (h Handler) DeleteMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    DeleteMediaRulesHandlerName,
		"request_id": requestID.String(),
	})

	ruleID := chi.URLParam(r, "resource")

	err := h.app.DeleteMediaRules(r.Context(), ruleID)
	if err != nil {
		ape.ApplicationError(w, requestID, err)
		log.WithError(err).Errorf("Error deleting media rule %s", ruleID)
		
		return
	}

	log.Infof("Media rule %s deleted successfully %s", ruleID, requestID.String())
	httpkit.Render(w, http.StatusNoContent)
}
