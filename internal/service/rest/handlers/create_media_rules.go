package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/media-storage/internal/api/requests"
	"github.com/chains-lab/media-storage/internal/api/responses"
	"github.com/chains-lab/media-storage/internal/app/ape"
	"github.com/go-chi/chi/v5"
)

func (h Handler) CreateMediaRules(w http.ResponseWriter, r *http.Request) {
	ruleID := chi.URLParam(r, "resource-category")

	req, err := requests.CreateMediaRules(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status: http.StatusBadRequest,
			Error:  err,
		})...)
		return
	}

	if ruleID != req.Data.Id {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status:   http.StatusBadRequest,
			Title:    "resource-category does not match id",
			Error:    fmt.Errorf("resource-category %s does not match id %s", ruleID, req.Data.Id),
			Parametr: "id",
			Pointer:  "/data/id",
		})...)
		return
	}

	curRoles, err := parseRoles(req.Data.Attributes.Roles)
	if err != nil {
		h.log.WithError(err).Warn("error parsing request")
		httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
			Status:  http.StatusBadRequest,
			Title:   "error parsing roles",
			Error:   err,
			Pointer: "/data/attributes/roles",
		})...)
		return
	}

	res, err := h.app.CreateMediaRules(r.Context(), domain.CreateMediaRulesRequest{
		ID:         ruleID,
		Extensions: req.Data.Attributes.Extensions,
		MaxSize:    req.Data.Attributes.MaxSize,
		Roles:      curRoles,
	})
	if err != nil {
		switch {
		case errors.Is(err, ape.ErrMediaRulesAlreadyExists):
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusConflict,
				Title:  "Media rules already exists",
			})...)
		default:
			httpkit.ResponseError(w, httpkit.ResponseError(httpkit.ReponseErrorInput{
				Status: http.StatusInternalServerError,
			})...)
		}
		h.log.WithError(err).Errorf("error create media rule %s", err)
		return
	}

	httpkit.Render(w, responses.MediaRules(res))
}

func parseRoles(r []string) ([]roles.Role, error) {
	parsedRoles := make([]roles.Role, len(r))
	for i, str := range r {
		pr, err := roles.ParseRole(str)
		if err != nil {
			return nil, err
		}
		parsedRoles[i] = pr
	}
	return parsedRoles, nil
}
