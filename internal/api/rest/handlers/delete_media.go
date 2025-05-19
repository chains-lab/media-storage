package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/api/rest/ape"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const DeleteMediaHandlerName = "DeleteMediaHandler"

func (h Handler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    DeleteMediaHandlerName,
		"request_id": requestID.String(),
	})

	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		ape.BadTokenData(w, requestID)
		log.WithError(err).Errorf("Error getting token data")

		return
	}

	mediaID, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		ape.BadRequest(w, requestID, "Invalid media ID parameter")
		log.WithError(err).Errorf("Invalid media_id format parameter")

		return
	}

	err = h.app.DeleteMedia(r.Context(), app.DeleteMediaRequest{
		MediaID:       mediaID,
		InitiatorRole: user.Role,
		InitiatorID:   user.AccountID,
	})
	if err != nil {
		ape.ApplicationError(w, requestID, err)

		return
	}

	log.Debugf("Media deleted successfully by user %s", user.AccountID)
	httpkit.Render(w, http.StatusNoContent)
}
