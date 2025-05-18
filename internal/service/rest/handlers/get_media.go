package handlers

import (
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/responses"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h Handler) GetMedia(w http.ResponseWriter, r *http.Request) {
	mediaId, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status:   http.StatusBadRequest,
			Title:    "error parsing media_id",
			Error:    err,
			Parametr: "media_id",
		})...)
		return
	}

	media, err := h.app.GetMedia(r.Context(), mediaId)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaNotFound):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusNotFound,
				Title:  "Media not found",
			})...)
		default:
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusInternalServerError,
			})...)
		}

		h.log.WithError(err).Error("error getting media")
		return
	}

	httpkit.Render(w, responses.Media(media))
}
