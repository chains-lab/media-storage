package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/api/rest/requests"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/google/uuid"
)

func (h Handler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	user, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		h.presenter.InvalidToken(w, requestID, err)
		return
	}

	req, file, fileHeader, err := requests.UploadMedia(r)
	if err != nil {
		h.presenter.InvalidPointer(w, requestID, err)
		return
	}

	requestToApp := app.UploadMediaRequest{
		FileHeader: fileHeader,
		File:       file,
		UserID:     user.UserID,
		UserRole:   user.Role,
		Category:   req.Data.Attributes.Category,
		Resource:   req.Data.Attributes.Resource,
		ResourceID: req.Data.Attributes.ResourceId,
	}

	res, appErr := h.app.UploadMedia(r.Context(), requestToApp)
	if appErr != nil {
		h.presenter.AppError(w, requestID, appErr)
		return
	}

	h.log.WithField("request_id", requestID).Infof("Media %s successfully uploaded by user: %s", res.ID, user.UserID)
	httpkit.Render(w, responses.Media(res))
}
