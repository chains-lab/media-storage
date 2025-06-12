package requests

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/chains-lab/gatekit/jsonkit"
	"github.com/chains-lab/media-storage/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UploadMedia парсит JSONAPI запрос из тела или из multipart-формы (поле upload_data)
// Возвращает структуру запроса и загруженный файл с метаданными
func UploadMedia(r *http.Request) (req resources.UploadMedia, file multipart.File, fileHeader *multipart.FileHeader, err error) {
	var raw []byte
	ct := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "multipart/form-data") {
		err = fmt.Errorf("invalid content type: %s", ct)
	}

	// multipart: JSON в поле upload_data + сам файл в поле file
	if err = r.ParseMultipartForm(32 << 20); err != nil {
		err = fmt.Errorf("parse multipart form: %w", err)
		return
	}
	raw = []byte(r.FormValue("upload_data"))
	fmt.Printf(">>> RAW UPLOAD_DATA = %q\n", raw)

	file, fileHeader, err = r.FormFile("file")
	if err != nil {
		err = fmt.Errorf("read uploaded file: %w", err)
		return
	}

	// декодируем JSONAPI
	if err = json.Unmarshal(raw, &req); err != nil {
		err = jsonkit.NewDecodeError("upload_data", err)
		return
	}

	// валидация полей JSONAPI
	errs := validation.Errors{
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In(resources.MediaUploadType)),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
		"data/attributes/resource": validation.Validate(req.Data.Attributes.Resource, validation.Required),
		"data/attributes/resource_id": validation.Validate(req.Data.Attributes.ResourceId, validation.Required),
		"data/attributes/category": validation.Validate(req.Data.Attributes.Category, validation.Required),
		"data/attributes/owner_id": validation.Validate(req.Data.Attributes.OwnerId, validation.Required),
	}
	if err = errs.Filter(); err != nil {
		return
	}

	return
}
