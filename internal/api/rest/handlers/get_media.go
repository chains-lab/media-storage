package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const GetMediaHandlerName = "GetMediaHandler"

func (h Handler) GetMedia(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	mediaId, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		h.presenter.InvalidParameter(w, requestID, err, "media_id")
		return
	}

	media, appErr := h.app.GetMedia(r.Context(), mediaId)
	if appErr != nil {
		h.presenter.AppError(w, requestID, appErr)
		return
	}

	h.log.WithField("request_id", requestID).Debugf("Media %s retrieved successfully", mediaId)
	httpkit.Render(w, responses.Media(media))
}
