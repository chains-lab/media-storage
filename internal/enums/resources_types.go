package enums

import "fmt"

type ResourceType string

const (
	ArticleResource ResourceType = "article"
	AuthorResource  ResourceType = "author"
	TagResource     ResourceType = "tag"

	UserResource ResourceType = "user"
)

func ParseResourceType(resourceType string) (ResourceType, error) {
	switch resourceType {
	case string(ArticleResource):
		return ArticleResource, nil
	case string(AuthorResource):
		return AuthorResource, nil
	case string(TagResource):
		return TagResource, nil
	case string(UserResource):
		return UserResource, nil
	default:
		return "", fmt.Errorf("invalid resource type: %s", resourceType)
	}
}
