package ape

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/uuid"
)

func BadRequest(w http.ResponseWriter, requestID uuid.UUID, details string) {
	httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
		Status:    http.StatusBadRequest,
		RequestID: requestID.String(),
		Detail:    details,
	})...)
}
