package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const DeleteMediaHandlerName = "DeleteMediaHandler"

func (h Handler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	user, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		h.presenter.InvalidToken(w, requestID, err)
		return
	}

	mediaID, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		h.presenter.InvalidParameter(w, requestID, err, "media_id")
		return
	}

	appErr := h.app.DeleteMedia(r.Context(), app.DeleteMediaRequest{
		MediaID:       mediaID,
		InitiatorRole: user.Role,
		InitiatorID:   user.UserID,
	})
	if appErr != nil {
		h.presenter.AppError(w, requestID, appErr)
		return
	}

	h.log.WithField("request_id", requestID).Debugf("Media deleted successfully by user %s", user.UserID)
	httpkit.Render(w, http.StatusNoContent)
}
