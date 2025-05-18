package handlers

import (
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h Handler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status: http.StatusUnauthorized,
			Title:  "unauthorized",
		})...)
		return
	}

	mediaID, err := uuid.Parse(chi.URLParam(r, "media_id"))
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status:   http.StatusBadRequest,
			Title:    "error parsing media_id",
			Error:    err,
			Parametr: "media_id",
		})...)
		return
	}

	err = h.app.DeleteMedia(r.Context(), domain.DeleteMediaRequest{
		UserID:      user.AccountID,
		UserRole:    user.Role,
		MediaID:     mediaID,
		InitiatorID: user.AccountID,
	})
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
		case errors.Is(err, ape.ErrUserNotAllowedToDeleteMedia):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusForbidden,
				Title:  "User not allowed to delete media",
			})...)
		default:
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusInternalServerError,
			})...)
		}

		h.log.WithError(err).Error("Error deleting media")
		return
	}

	httpkit.Render(w, http.StatusNoContent)
}
