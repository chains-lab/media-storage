package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/media-storage/internal/api/requests"
	"github.com/hs-zavet/media-storage/internal/api/responses"
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/internal/app/ape"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

func (h *Handler) UpdateMediaRules(w http.ResponseWriter, r *http.Request) {
	mediaResourceType, err := enums.ParseMediaType(chi.URLParam(r, "media_resource_type"))
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	req, err := requests.UpdateMediaRules(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if string(mediaResourceType) == req.Data.Id {
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"media_resource_id": validation.NewError("media_resource_id", "invalid media resource id"),
		})...)
		return
	}

	var updateReq app.UpdateMediaRulesRequest
	if req.Data.Attributes.AllowedExits != nil {
		updateReq.AllowedExits = req.Data.Attributes.AllowedExits
	}
	if req.Data.Attributes.Folder != nil {
		updateReq.Folder = *req.Data.Attributes.Folder
	}
	if req.Data.Attributes.MaxSize != nil {
		updateReq.MaxSize = *req.Data.Attributes.MaxSize
	}
	if req.Data.Attributes.Roles != nil {
		newRoles := make([]roles.Role, 0, len(req.Data.Attributes.Roles))
		for _, role := range req.Data.Attributes.Roles {
			curRole, err := roles.ParseRole(role)
			if err != nil {
				h.log.WithError(err).Warn("error parsing request")
				httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
					"roles": validation.NewError("roles", "invalid role"),
				})...)
				return
			}
			newRoles = append(newRoles, curRole)
		}
		updateReq.Roles = newRoles
	}

	res, err := h.app.UpdateMediaRules(r.Context(), mediaResourceType, updateReq)
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesNotFound):
			httpkit.RenderErr(w, problems.NotFound("media resource not found"))
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Errorf("error updating media rule %s", mediaResourceType)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}
