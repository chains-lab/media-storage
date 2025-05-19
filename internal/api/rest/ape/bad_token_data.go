package ape

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/google/uuid"
)

func BadTokenData(w http.ResponseWriter, requestID uuid.UUID) {
	httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
		Status:    http.StatusUnauthorized,
		RequestID: requestID.String(),
	})...)
}
