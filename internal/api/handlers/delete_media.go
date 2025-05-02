package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/tokens"
)

func (h *Handler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resourceID, err := uuid.Parse(chi.URLParam(r, "media_resource_type"))
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"media_resource_type": validation.NewError("media_resource_type", "invalid UUID format"),
		})...)
		return
	}

	err = h.app.DeleteMedia(r.Context(), app.DeleteMediaRequest{
		Role:        user.Role,
		ResourceID:  resourceID,
		InitiatorID: user.AccountID,
	})
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaNotFound):
			httpkit.RenderErr(w, problems.NotFound("media not found"))
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			//TODO: check if this is the right error
			httpkit.RenderErr(w, problems.NotFound("media rules for this media type not found"))
		case errors.Is(err, ape.ErrUserNotAllowedToDeleteMedia):
			httpkit.RenderErr(w, problems.Forbidden("user not allowed to delete media"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}

		h.log.WithError(err).Error("Error deleting media")
		return
	}

	httpkit.Render(w, http.StatusNoContent)
}
