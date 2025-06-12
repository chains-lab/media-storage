package ape

const (
	//General error codes

	CodeInternal             = "INTERNAL_SERVER_ERROR"
	CodeInvalidRequestBody   = "INVALID_REQUEST_BODY"
	CodeInvalidRequestQuery  = "INVALID_REQUEST_QUERY"
	CodeInvalidRequestHeader = "INVALID_REQUEST_HEADER"
	CodeInvalidRequestPath   = "INVALID_REQUEST_PATH"
	UnauthorizedError        = "UNAUTHORIZED"

	//Specific error codes

	CodeMediaRulesNotFound       = "MEDIA_RULES_NOT_FOUND"
	CodeFileTooLarge             = "FILE_TOO_LARGE"
	CodeMediaExtensionNotAllowed = "MEDIA_EXTENSION_NOT_ALLOWED"

	//CodeFileFormatNotAllowed        = "FILE_FORMAT_NOT_ALLOWED"
	//CodeUserNotAllowedToUploadMedia = "USER_NOT_ALLOWED_TO_UPLOAD_MEDIA"
	//CodeUserNotAllowedToDeleteMedia = "USER_NOT_ALLOWED_TO_DELETE_MEDIA"
	//CodeMediaNotFound               = "MEDIA_DOES_NOT_FOUND"
	//CodeMediaAlreadyExists          = "MEDIA_ALREADY_EXISTS"
	//CodeInvalidMediaId              = "INVALID_MEDIA_ID"
)
