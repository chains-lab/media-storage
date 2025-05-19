package ape

import (
	"errors"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/google/uuid"
)

const (
	//General error codes

	CodeInvalidRequestBody   = "INVALID_REQUEST_BODY"
	CodeInvalidRequestQuery  = "INVALID_REQUEST_QUERY"
	CodeInvalidRequestHeader = "INVALID_REQUEST_HEADER"
	CodeInvalidRequestPath   = "INVALID_REQUEST_PATH"
	CodeInvalidRequestMethod = "INVALID_REQUEST_METHOD"
	UnauthorizedError        = "UNAUTHORIZED"

	//Specific error codes

	CodeMediaRulesNotFound          = "MEDIA_RULES_NOT_FOUND"
	CodeMediaRulesAlreadyExists     = "MEDIA_RULES_ALREADY_EXISTS"
	CodeFileToLarge                 = "FILE_TOO_LARGE"
	CodeFileFormatNotAllowed        = "FILE_FORMAT_NOT_ALLOWED"
	CodeUserNotAllowedToUploadMedia = "USER_NOT_ALLOWED_TO_UPLOAD_MEDIA"
	CodeUserNotAllowedToDeleteMedia = "USER_NOT_ALLOWED_TO_DELETE_MEDIA"
	CodeMediaNotFound               = "MEDIA_DOES_NOT_FOUND"
	CodeMediaAlreadyExists          = "MEDIA_ALREADY_EXISTS"
	CodeInternal                    = "INTERNAL_SERVER_ERROR"
)

func ApplicationError(w http.ResponseWriter, requestID uuid.UUID, err error) {

	switch {
	case errors.Is(err, app.ErrMediaRulesNotFound):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusNotFound,
			Code:      CodeMediaRulesNotFound,
			Title:     "Media rules not found",
			Detail:    "Media rules not found",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrMediaRulesAlreadyExists):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusConflict,
			Code:      CodeMediaRulesAlreadyExists,
			Title:     "Media rules already exists",
			Detail:    "Media rules already exists",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrFileToLarge):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusRequestEntityTooLarge,
			Code:      CodeFileToLarge,
			Title:     "File too large",
			Detail:    "File too large",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrUserNotAllowedToUploadMedia):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusForbidden,
			Code:      CodeUserNotAllowedToUploadMedia,
			Title:     "User not allowed to upload media",
			Detail:    "User not allowed to upload media",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrUserNotAllowedToDeleteMedia):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusForbidden,
			Code:      CodeUserNotAllowedToDeleteMedia,
			Title:     "User not allowed to delete media",
			Detail:    "User not allowed to delete media",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrMediaNotFound):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusNotFound,
			Code:      CodeMediaNotFound,
			Title:     "Media not found",
			Detail:    "Media not found",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrMediaExtensionNotAllowed):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusBadRequest,
			Code:      CodeFileFormatNotAllowed,
			Title:     "File format not allowed",
			Detail:    "File format not allowed",
			RequestID: requestID.String(),
		})...)

	case errors.Is(err, app.ErrMediaAlreadyExists):
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusConflict,
			Code:      CodeMediaAlreadyExists,
			Title:     "Media already exists",
			Detail:    "Media already exists",
			RequestID: requestID.String(),
		})...)

	default:
		httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
			Status:    http.StatusInternalServerError,
			Code:      CodeInternal,
			RequestID: requestID.String(),
		})...)
	}
}
