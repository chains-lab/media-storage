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

const GetMediaHandlerName = "GetMediaHandler"

func (h Handler) GetMedia(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    GetMediaHandlerName,
		"request_id": requestID.String(),
	})

	mediaId, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		ape.BadRequest(w, requestID, "Invalid media ID parameter")
		log.WithError(err).Errorf("Invalid media_id format parameter")

		return
	}

	media, err := h.app.GetMedia(r.Context(), mediaId)
	if err != nil {
		ape.ApplicationError(w, requestID, err)
		log.WithError(err).Errorf("Error getting media %s", mediaId)

		return
	}

	log.Debugf("Media %s retrieved successfully", mediaId)
	httpkit.Render(w, responses.Media(media))
}
