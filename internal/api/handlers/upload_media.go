package handlers

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/requests"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/tokens"
)

func (h *Handler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	req, file, fileHeader, err := requests.UploadMedia(r)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
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
			httpkit.RenderErr(w, problems.NotFound("media not found"))
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			httpkit.RenderErr(w, problems.NotFound("media rules for this media type not found"))
		case errors.Is(err, ape.ErrFileToLarge):
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"upload_data": validation.NewError("file", "file too large"),
			})...)
		case errors.Is(err, ape.ErrFileFormatNotAllowed):
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"upload_data": validation.NewError("file", "file format not allowed"),
			})...)
		case errors.Is(err, ape.ErrUserNotAllowedToUploadMedia):
			httpkit.RenderErr(w, problems.Forbidden("user role not allowed to upload this type media"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}

		h.log.WithError(err).Errorf("error uploading media")
		return
	}

	h.log.Infof("MediaModels %s successfully uploaded by user: %s", res.ID, user.AccountID)

	httpkit.Render(w, responses.Media(res))
}
