package controllers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/google/uuid"
)

func (c Controller) ParameterFromURL(w http.ResponseWriter, requestID uuid.UUID, err error, parameter string) {
	errorID := uuid.New()

	c.log.WithField("request_id", requestID).
		WithField("error_id", errorID).
		WithError(err).
		Errorf("error getting %s from url", parameter)

	httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
		Status:    http.StatusUnauthorized,
		Code:      ape.CodeInvalidRequestHeader,
		Detail:    err.Error(),
		RequestID: requestID.String(),
		ErrorID:   errorID.String(),
	})...)
}
