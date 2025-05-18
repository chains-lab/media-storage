package controllers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
)

func (c Controller) BadRequestValidation(w http.ResponseWriter, err error) {
	httpkit.RenderErr(w, httpkit.ResponseError(httpkit.ResponseErrorInput{
		Status: http.StatusBadRequest,
		Error:  err,
	})...)
	return
}
