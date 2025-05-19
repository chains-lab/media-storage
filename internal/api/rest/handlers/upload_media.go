package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/api/rest/ape"
	"github.com/chains-lab/media-storage/internal/api/rest/requests"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (h Handler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    "UploadMediaHandler",
		"request_id": requestID.String(),
	})

	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		ape.BadTokenData(w, requestID)
		log.WithError(err).Errorf("Error getting token data")

		return
	}

	req, file, fileHeader, err := requests.UploadMedia(r)
	if err != nil {
		ape.BadRequest(w, requestID, "Error parsing request body")
		log.WithError(err).Errorf("Error parsing request body")

		return
	}

	requestToApp := app.UploadMediaRequest{
		FileHeader: fileHeader,
		File:       file,
		UserID:     user.AccountID,
		UserRole:   user.Role,
		Category:   req.Data.Attributes.Category,
		Resource:   req.Data.Attributes.Resource,
		ResourceID: req.Data.Attributes.ResourceId,
	}

	res, err := h.app.UploadMedia(r.Context(), requestToApp)
	if err != nil {
		ape.ApplicationError(w, requestID, err)
		log.WithError(err).Errorf("Error uploading media")

		return
	}

	log.Infof("Media %s successfully uploaded by user: %s", res.ID, user.AccountID)
	httpkit.Render(w, responses.Media(res))
}
