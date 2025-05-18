package controllers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/google/uuid"
)

func (c Controller) ResultFromApp(w http.ResponseWriter, requestID uuid.UUID, appErr *ape.Error) {
	errorID := uuid.New()
	c.log.WithField("request_id", requestID).
		WithField("error_id", errorID).
		WithField("code", appErr.Code).
		WithError(appErr.Unwrap()).
		Error("error from app")

	base := httpkit.ResponseErrorInput{
		Code:      appErr.Code,
		Title:     appErr.Title,
		Detail:    appErr.Details,
		RequestID: requestID.String(),
		ErrorID:   errorID.String(),
	}

	switch appErr.Code {
	// resource not found
	case ape.CodeMediaRulesNotFound,
		ape.CodeMediaNotFound:
		base.Status = http.StatusNotFound

	// conflict / already exists
	case ape.CodeMediaAlreadyExists,
		ape.CodeMediaRulesAlreadyExists:
		base.Status = http.StatusConflict

	// bad request
	case ape.CodeFileToLarge,
		ape.CodeFileFormatNotAllowed:
		base.Status = http.StatusBadRequest

	// forbidden
	case ape.CodeUserNotAllowedToUploadMedia,
		ape.CodeUserNotAllowedToDeleteMedia:
		base.Status = http.StatusForbidden

	// internal
	case ape.CodeInternal:
		base.Status = http.StatusInternalServerError

	// catch-all
	default:
		base.Status = http.StatusInternalServerError
		base.Code = ape.CodeInternal
		base.Title = "Internal server error"
		base.Detail = "An unexpected error occurred"
	}

	httpkit.RenderErr(w, httpkit.ResponseError(base)...)
}
