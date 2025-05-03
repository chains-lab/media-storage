package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
)

func (h *Handler) DeleteMediaRules(w http.ResponseWriter, r *http.Request) {
	mediaResourceType := chi.URLParam(r, "resource_type")

	err := h.app.DeleteMediaRules(r.Context(), mediaResourceType)
	if err != nil {
		switch {
		//TODO: add more specific errors add validation for delete resources
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Error("Error deleting media rule")
		return
	}

	httpkit.Render(w, http.StatusNoContent)
}
