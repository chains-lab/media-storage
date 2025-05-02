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

	return resources.MediaRules{
		Data: resources.MediaRulesData{
			Type: resources.MediaRulesCollectionType,
			Id:   string(mediaRules.MediaType),
			Attributes: resources.MediaRulesAttributes{
				MaxSize:      mediaRules.MaxSize,
				AllowedExits: mediaRules.AllowedExits,
				Folder:       mediaRules.Folder,
				Roles:        roles,
			},
		},
	}
}
