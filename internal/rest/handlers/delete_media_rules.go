package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const DeleteMediaRulesHandlerName = "DeleteMediaRulesHandler"

func (h Handler) DeleteMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	ruleID := chi.URLParam(r, "resource")

	err := h.app.DeleteMediaRules(r.Context(), ruleID)
	if err != nil {
		h.presenter.AppError(w, requestID, err)
		return
	}

	h.log.WithField("request_id", requestID).Infof("Media rule %s deleted successfully %s", ruleID, requestID.String())
	httpkit.Render(w, http.StatusNoContent)
}
