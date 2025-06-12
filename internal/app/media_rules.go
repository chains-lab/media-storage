package app

import (
	"context"

	"github.com/chains-lab/media-storage/internal/aws"
)

func (a App) RequestUploadLink(ctx context.Context, resource, category) (*aws.MediaURL, error) {
	rules, err := a.rules.GetMediaRules(ctx, resource, category)
	if err != nil {
		return nil, err
	}

	return a.rules.GenerateMediaPutURL(ctx, folder, originalFilename, contentType)
}
