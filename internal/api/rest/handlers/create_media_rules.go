package handlers

import (
	"net/http"

	"github.com/chains-lab/gatekit/httpkit"
	"github.com/chains-lab/gatekit/roles"
	"github.com/chains-lab/gatekit/tokens"
	"github.com/chains-lab/media-storage/internal/api/rest/requests"
	"github.com/chains-lab/media-storage/internal/api/rest/responses"
	"github.com/chains-lab/media-storage/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const CreateMediaRulesHandlerName = "CreateMediaRulesHandler"

func (h Handler) CreateMediaRules(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New()

	user, err := tokens.GetUserTokenData(r.Context())
	if err != nil {
		h.presenter.InvalidToken(w, requestID, err)
		return
	}

	resource := chi.URLParam(r, "resource")

	req, err := requests.CreateMediaRules(r)
	if err != nil {
		h.presenter.InvalidPointer(w, requestID, err)
		return
	}

	if resource != req.Data.Id {
		h.presenter.MismatchIdentification(w, requestID, "resource", req.Data.Id)
		return
	}

	rolesInReq, err := parseRoles(req.Data.Attributes.Roles)
	if err != nil {
		h.presenter.InvalidPointer(w, requestID, err)
		return
	}

	res, appErr := h.app.CreateMediaRules(r.Context(), app.CreateMediaRulesRequest{
		ID:         req.Data.Id,
		Extensions: req.Data.Attributes.Extensions,
		MaxSize:    req.Data.Attributes.MaxSize,
		Roles:      rolesInReq,
	})
	if appErr != nil {
		h.presenter.AppError(w, requestID, appErr)
		return
	}

	h.log.WithField("request_id", requestID).Infof("Created media rules %s by user: %s", req.Data.Id, user.UserID)
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
