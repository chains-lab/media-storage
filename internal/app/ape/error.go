package ape

import "fmt"

var (
	ErrMediaRulesNotFound            = fmt.Errorf("media rules not found")
	ErrMediaRulesAlreadyExists       = fmt.Errorf("media rules already exists")
	ErrFileToLarge                   = fmt.Errorf("file too large")
	ErrFileFormatNotAllowed          = fmt.Errorf("file format not allowed")
	ErrUserNotAllowedToUploadMedia   = fmt.Errorf("user not allowed to upload media")
	ErrUserNotAllowedToDownloadMedia = fmt.Errorf("user not allowed to download media")
	ErrUserNotAllowedToDeleteMedia   = fmt.Errorf("user not allowed to delete media")
	ErrMediaNotFound                 = fmt.Errorf("media not found")
)
