package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
)

func (h *Handler) DeleteMediaRules(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "resource-category")

	err := h.app.DeleteMediaRules(r.Context(), ruleID)
	if err != nil {
		switch {
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Error("Error deleting media rule")
		return
	}

	httpkit.Render(w, http.StatusNoContent)
}
