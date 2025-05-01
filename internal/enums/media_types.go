package enums

import "fmt"

type MediaType string

const (
	ArticleIcon     MediaType = "article_icon"
	ArticleImage    MediaType = "article_image"
	ArticleVideo    MediaType = "article_video"
	ArticleAudio    MediaType = "article_audio"
	ArticleDocument MediaType = "article_document"

	AuthorIcon MediaType = "author_icon"

	TagIcon MediaType = "tag_icon"

	UserIcon MediaType = "user_icon"
)

func ParseMediaType(contentType string) (MediaType, error) {
	switch contentType {
	case string(ArticleIcon):
		return ArticleIcon, nil
	case string(ArticleImage):
		return ArticleImage, nil
	case string(ArticleVideo):
		return ArticleVideo, nil
	case string(ArticleAudio):
		return ArticleAudio, nil
	case string(ArticleDocument):
		return ArticleDocument, nil
	case string(AuthorIcon):
		return AuthorIcon, nil
	case string(TagIcon):
		return TagIcon, nil
	case string(UserIcon):
		return UserIcon, nil
	default:
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}
}
