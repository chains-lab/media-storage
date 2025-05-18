package ape

const (
	//General error codes

	CodeInvalidRequestBody   = "INVALID_REQUEST_BODY"
	CodeInvalidRequestQuery  = "INVALID_REQUEST_QUERY"
	CodeInvalidRequestHeader = "INVALID_REQUEST_HEADER"
	CodeInvalidRequestPath   = "INVALID_REQUEST_PATH"
	CodeInvalidRequestMethod = "INVALID_REQUEST_METHOD"
	UnauthorizedError        = "UNAUTHORIZED"

	//Specific error codes

	CodeMediaRulesNotFound          = "MEDIA_RULES_NOT_FOUND"
	CodeMediaRulesAlreadyExists     = "MEDIA_RULES_ALREADY_EXISTS"
	CodeFileToLarge                 = "FILE_TOO_LARGE"
	CodeFileFormatNotAllowed        = "FILE_FORMAT_NOT_ALLOWED"
	CodeUserNotAllowedToUploadMedia = "USER_NOT_ALLOWED_TO_UPLOAD_MEDIA"
	CodeUserNotAllowedToDeleteMedia = "USER_NOT_ALLOWED_TO_DELETE_MEDIA"
	CodeMediaNotFound               = "MEDIA_DOES_NOT_FOUND"
	CodeMediaAlreadyExists          = "MEDIA_ALREADY_EXISTS"
	CodeInternal                    = "INTERNAL_SERVER_ERROR"
)
