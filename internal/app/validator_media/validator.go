package validator_media

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"regexp"

	"github.com/google/uuid"
	"github.com/hs-zavet/media-storage/internal/enums"
	"github.com/hs-zavet/tokens/roles"
)

var placeholderRE = regexp.MustCompile(`\{[^}]+\}`)

type MediaRules struct {
	MaxSize           int64
	AllowedExits      []string
	Folder            string
	RolesAccessUpdate []roles.Role
}

type Validator struct {
	rules map[enums.MediaType]MediaRules
}

func NewValidator() Validator {
	return Validator{
		rules: map[enums.MediaType]MediaRules{
			enums.ArticleIcon: {
				MaxSize:           5 << 20,
				AllowedExits:      []string{".png", ".jpg", ".jpeg"},
				Folder:            "news/articles/{article_id}/icon",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.ArticleImage: {
				MaxSize:           10 << 20,
				AllowedExits:      []string{".png", ".jpg", ".jpeg", ".webp"},
				Folder:            "news/articles/{article_id}/images",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.ArticleVideo: {
				MaxSize:           100 << 20,
				AllowedExits:      []string{".mp4", ".mov"},
				Folder:            "news/articles/{article_id}/video",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.ArticleAudio: {
				MaxSize:           50 << 20,
				AllowedExits:      []string{".mp3", ".wav"},
				Folder:            "news/articles/{article_id}/audio",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.ArticleDocument: {
				MaxSize:           20 << 20,
				AllowedExits:      []string{".pdf", ".docx", ".txt"},
				Folder:            "news/articles/{article_id}/documents",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.AuthorIcon: {
				MaxSize:           5 << 20,
				AllowedExits:      []string{".png", ".jpg", ".jpeg"},
				Folder:            "news/authors/{author_id}/icon",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.TagIcon: {
				MaxSize:           5 << 20,
				AllowedExits:      []string{".png", ".jpg", ".jpeg"},
				Folder:            "news/tags/{tag_id}/icon",
				RolesAccessUpdate: []roles.Role{roles.Admin, roles.SuperUser},
			},

			enums.UserIcon: {
				MaxSize:           5 << 20,
				AllowedExits:      []string{".png", ".jpg", ".jpeg"},
				Folder:            "users/{user_id}/icon",
				RolesAccessUpdate: []roles.Role{roles.User, roles.Admin, roles.SuperUser},
			},
		},
	}
}

func (v Validator) ValidateUpdate(mediaType enums.MediaType, fileHeader *multipart.FileHeader, resourceID uuid.UUID, role roles.Role) (string, error) {
	rule, ok := v.rules[mediaType]
	if !ok {
		return "", fmt.Errorf("unknown media type: %s", mediaType)
	}

	if fileHeader.Size > rule.MaxSize {
		return "", fmt.Errorf("file too large (max %d)", rule.MaxSize)
	}

	ext := filepath.Ext(fileHeader.Filename)
	if !contains(rule.AllowedExits, ext) {
		return "", fmt.Errorf("extension %s not allowed for %s", ext, mediaType)
	}

	access := false
	for _, r := range rule.RolesAccessUpdate {
		if r == role {
			access = true
			break
		}
	}
	if !access {
		return "", fmt.Errorf("role %s not allowed for %s", role, mediaType)
	}

	folder := placeholderRE.ReplaceAllString(rule.Folder, resourceID.String())
	return folder, nil
}

func contains(exits []string, exit string) bool {
	for _, v := range exits {
		if v == exit {
			return true
		}
	}
	return false
}
