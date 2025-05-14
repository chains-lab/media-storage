package handlers

import (
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/api/requests"
	"github.com/chains-lab/media-storage/internal/api/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/chains-lab/media-storage/internal/app/ape"
)

func (h *Handler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status: http.StatusUnauthorized,
			Title:  "unauthorized",
		})...)
		return
	}

	req, file, fileHeader, err := requests.UploadMedia(r)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status: http.StatusBadRequest,
			Error:  err,
		})...)
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
		switch {
		case errors.Is(err, ape.ErrMediaNotFound):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusNotFound,
				Title:  "Media not found",
			})...)
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusNotFound,
				Title:  "Media rules not found",
			})...)
		case errors.Is(err, ape.ErrFileToLarge):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusBadRequest,
				Title:  "file to large",
			})...)
		case errors.Is(err, ape.ErrFileFormatNotAllowed):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusBadRequest,
				Title:  "file format is not allowed",
			})...)
		case errors.Is(err, ape.ErrUserNotAllowedToUploadMedia):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusForbidden,
				Title:  "User not allowed to upload media",
			})...)
		default:
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusInternalServerError,
			})...)
		}

		h.log.WithError(err).Errorf("error uploading media")
		return
	}

	h.log.Infof("MediaModels %s successfully uploaded by user: %s", res.ID, user.AccountID)

	httpkit.Render(w, responses.Media(res))
}
