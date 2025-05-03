package responses

import (
	"github.com/hs-zavet/media-storage/internal/app/models"
	"github.com/hs-zavet/media-storage/resources"
)

func MediaRules(mediaRules models.MediaRules) resources.MediaRules {
	roles := make([]string, 0, len(mediaRules.Roles))
	for _, role := range mediaRules.Roles {
		roles = append(roles, string(role))
	}

	exitSize := make([]resources.ExitSizeInner, 0, len(mediaRules.ExitSize))
	for _, el := range mediaRules.ExitSize {
		exitSize = append(exitSize, resources.ExitSizeInner{
			Exit: el.Exit,
			Size: el.Size,
		})
	}

	return resources.MediaRules{
		Data: resources.MediaRulesData{
			Type: resources.MediaRulesCollectionType,
			Id:   mediaRules.ResourceType,
			Attributes: resources.MediaRulesAttributes{
				ExitSize:  exitSize,
				Roles:     roles,
				UpdatedAt: mediaRules.UpdatedAt,
				CreatedAt: mediaRules.CreatedAt,
			},
		},
	}
}
