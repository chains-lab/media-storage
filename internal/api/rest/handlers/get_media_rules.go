package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const GetMediaRulesHandlerName = "GetMediaRulesHandler"

func (h Handler) GetMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	ruleID := chi.URLParam(r, "resource")

	res, err := h.app.GetMediaRules(r.Context(), ruleID)
	if err != nil {
		h.presenter.AppError(w, requestID, err)
		return
	}

	h.log.WithField("request_id", requestID).Debugf("Media rules %s retrieved successfully", ruleID)
	httpkit.Render(w, responses.MediaRules(res))
}
