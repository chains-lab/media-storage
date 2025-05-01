package handlers

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/requests"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app"
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
		h.log.WithError(err).Warn("error parsing request1")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resourcesType, err := enums.ParseResourceType(req.Data.Attributes.ResourceType)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"resource_type": validation.NewError("resource_type", "invalid resource type"),
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
	}

	if req.Data.Attributes.OwnerId != nil {
		ownerID, err := uuid.Parse(*req.Data.Attributes.OwnerId)
		if err != nil {
			h.log.WithError(err).Warn("error parsing request")
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		requestToApp.OwnerID = &ownerID
	}

	res, err := h.app.UploadMedia(r.Context(), user, file, fileHeader, requestToApp)
	if err != nil {
		switch {
		case errors.Is(err, nil):
			httpkit.RenderErr(w, problems.NotFound("media not found"))
		default:
			h.log.WithError(err).Errorf("error uploading media")
			httpkit.RenderErr(w, problems.InternalError())
		}
		return
	}

	h.log.Infof("Media %s successfully uploaded by user: %s", res.ID, user.AccountID)

	httpkit.Render(w, responses.UploadMedia(res))
}
