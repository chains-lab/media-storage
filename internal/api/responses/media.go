package responses

import (
	"github.com/hs-zavet/media-storage/internal/app/models"
	"github.com/hs-zavet/media-storage/resources"
)

func Media(media models.Media) resources.Media {
	attributes := resources.MediaAttributes{
		Format:    media.Ext,
		Size:      media.Size,
		Url:       media.URL,
		CreatedAt: media.CreatedAt,
	}

	res := resources.Media{
		Data: resources.MediaData{
			Id:         media.ID.String(),
			Type:       resources.MediaType,
			Attributes: attributes,
		},
	}

	return res
}
