package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/requests"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/media-storage/internal/enums"
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

	resourcesType, err := enums.ParseResourceType(chi.URLParam(r, "media_resource_type"))
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"media_resource_type": validation.NewError("media_resource_type", "invalid resource type"),
		})...)
		return
	}

	mediaType, err := enums.ParseMediaType(req.Data.Attributes.MediaType)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"content_type": validation.NewError("media_type", "invalid content type"),
		})...)
		return
	}

	ResourcesID, err := uuid.Parse(req.Data.Attributes.ResourceId)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"resource_id": validation.NewError("resource_id", "invalid UUID format"),
		})...)
		return
	}

	requestToApp := app.UploadMediaRequest{
		ResourceType: resourcesType,
		ResourceID:   ResourcesID,
		MediaType:    mediaType,
		User:         user,
		File:         file,
		FileHeader:   fileHeader,
	}

	res, err := h.app.UploadMedia(r.Context(), requestToApp)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaNotFound):
			httpkit.RenderErr(w, problems.NotFound("media not found"))
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			//TODO: check if this is the right error
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

	h.log.Infof("Media %s successfully uploaded by user: %s", res.ID, user.AccountID)

	httpkit.Render(w, responses.Media(res))
}
