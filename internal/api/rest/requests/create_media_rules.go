package requests

import (
	"encoding/json"
	"net/http"

	"github.com/chains-lab/gatekit/jsonkit"
	"github.com/chains-lab/media-storage/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CreateMediaRules(r *http.Request) (req resources.CreateMediaRules, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = jsonkit.NewDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/id":                    validation.Validate(req.Data.Id, validation.Required, validation.Length(1, 255)),
		"data/type":                  validation.Validate(req.Data.Type, validation.Required, validation.In(resources.CreateMediaRulesType)),
		"data/attributes":            validation.Validate(req.Data.Attributes, validation.Required),
		"data/attributes/extensions": validation.Validate(req.Data.Attributes.Extensions, validation.Required, validation.Length(1, 255)),
		"data/attributes/max_size":   validation.Validate(req.Data.Attributes.MaxSize, validation.Required, validation.Min(1)),
		"data/attributes/roles":      validation.Validate(req.Data.Attributes.Roles, validation.Required),
	}
	if err = errs.Filter(); err != nil {
		return
	}

	return req, errs.Filter()
}
