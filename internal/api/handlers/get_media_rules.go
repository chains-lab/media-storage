package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/media-storage/internal/enums"
)

func (h *Handler) GetMediaRules(w http.ResponseWriter, r *http.Request) {
	mediaResourceType, err := enums.ParseMediaType(chi.URLParam(r, "media_resource_type"))
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	res, err := h.app.GetMediaRules(r.Context(), mediaResourceType)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			httpkit.RenderErr(w, problems.NotFound("media rules not found"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Error("Error getting media rules")
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}
