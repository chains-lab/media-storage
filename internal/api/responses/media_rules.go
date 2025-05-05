package responses

import (
	"github.com/hs-zavet/media-storage/internal/app"
	"github.com/hs-zavet/media-storage/resources"
)

func MediaRules(mediaRules app.MediaRulesModel) resources.MediaRules {
	roles := make([]string, 0, len(mediaRules.AllowedRoles))
	for _, role := range mediaRules.AllowedRoles {
		roles = append(roles, string(role))
	}

	return resources.MediaRules{
		Data: resources.MediaRulesData{
			Type: resources.MediaRulesCollectionType,
			Id:   mediaRules.ID,
			Attributes: resources.MediaRulesAttributes{
				Extensions: mediaRules.Extensions,
				MaxSize:    mediaRules.MaxSize,
				Roles:      roles,
				UpdatedAt:  mediaRules.UpdatedAt,
				CreatedAt:  mediaRules.CreatedAt,
			},
		},
	}
}
