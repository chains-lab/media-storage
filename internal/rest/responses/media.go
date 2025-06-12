package responses

import (
	"github.com/chains-lab/media-storage/internal/app/models"
	"github.com/chains-lab/media-storage/resources"
)

func Media(media models.Media) resources.Media {
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
