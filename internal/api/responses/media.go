package responses

import (
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/resources"
)

func Media(media app.MediaModels) resources.Media {
	attributes := resources.MediaAttributes{
		Format:     media.Format,
		Extension:  media.Extension,
		Size:       media.Size,
		Url:        media.Url,
		Resource:   media.Resource,
		ResourceId: media.ResourceID,
		Category:   media.Category,
		OwnerId:    media.OwnerID.String(),
		CreatedAt:  media.CreatedAt,
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
