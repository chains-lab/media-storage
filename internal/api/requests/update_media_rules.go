package requests

import (
	"encoding/json"
	"net/http"

	"github.com/chains-lab/gatekit/jsonkit"
	"github.com/chains-lab/media-storage/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func UpdateMediaRules(r *http.Request) (req resources.UpdateMediaRules, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = jsonkit.NewDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In(resources.UpdateMediaRulesType)),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	if err = errs.Filter(); err != nil {
		return
	}

	return req, errs.Filter()
}
