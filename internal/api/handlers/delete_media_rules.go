package handlers

import (
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) DeleteMediaRules(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "resource-category")

	err := h.app.DeleteMediaRules(r.Context(), ruleID)
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
		h.log.WithError(err).Error("Error deleting media rule")
		return
	}

	httpkit.Render(w, http.StatusNoContent)
}
