package app

import "fmt"

var (
	ErrMediaRulesNotFound          = fmt.Errorf("media rules not found")
	ErrMediaRulesAlreadyExists     = fmt.Errorf("media rules already exists")
	ErrFileToLarge                 = fmt.Errorf("file too large")
	ErrUserNotAllowedToUploadMedia = fmt.Errorf("user not allowed to upload media")
	ErrUserNotAllowedToDeleteMedia = fmt.Errorf("user not allowed to delete media")
	ErrMediaAlreadyExists          = fmt.Errorf("media already exists")
	ErrMediaNotFound               = fmt.Errorf("media not found")
	ErrMediaExtensionNotAllowed    = fmt.Errorf("media extension not allowed")
)
