package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/api/rest/ape"
	"github.com/chains-lab/media-storage/internal/api/rest/requests"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const CreateMediaRylesHandlerName = "CreateMediaRulesHandler"

func (h Handler) CreateMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()
	log := h.log.WithFields(logrus.Fields{
		"handler":    CreateMediaRylesHandlerName,
		"request_id": requestID.String(),
	})

	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		ape.BadRequest(w, requestID, "Error getting token data")
		log.WithError(err).Errorf("Error getting token data")
		
		return
	}

	resource := chi.URLParam(r, "resource")

	req, err := requests.CreateMediaRules(r)
	if err != nil {
		ape.BadRequest(w, requestID, err.Error())
		log.WithError(err).Errorf("Error parsing request body")

		return
	}

	if resource != req.Data.Id {
		ape.BadRequest(w, requestID, "Resource ID does not match the request ID")
		log.Errorf("Resource ID %s does not match the request ID %s", resource, req.Data.Id)

		return
	}

	rolesInReq, err := parseRoles(req.Data.Attributes.Roles)
	if err != nil {
		ape.BadRequest(w, requestID, "Error parsing roles")
		log.WithError(err).Errorf("Error parsing roles")

		return
	}

	res, err := h.app.CreateMediaRules(r.Context(), app.CreateMediaRulesRequest{
		ID:         req.Data.Id,
		Extensions: req.Data.Attributes.Extensions,
		MaxSize:    req.Data.Attributes.MaxSize,
		Roles:      rolesInReq,
	})
	if err != nil {
		ape.ApplicationError(w, requestID, err)
		log.WithError(err).Errorf("Error creating media rules")

		return
	}

	log.Infof("Created media rules %s by user: %s", req.Data.Id, user.AccountID)
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
