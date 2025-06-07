package ape

// ваш тип
type Error struct {
	Code    string
	Title   string
	Details string
	cause   error
}

func (e *Error) Unwrap() error { return e.cause }

// general
func ErrorInvalidRequestBody(cause error) *Error {
	return &Error{
		Code:    CodeInvalidRequestBody,
		Title:   "Invalid request body",
		Details: cause.Error(),
		cause:   cause,
	}
}
func ErrorInvalidRequestQuery(cause error) *Error {
	return &Error{
		Code:    CodeInvalidRequestQuery,
		Title:   "Invalid request query",
		Details: cause.Error(),
		cause:   cause,
	}
}

// …и так далее для header, path, method, unauthorized…

// specific
func ErrorMediaRulesNotFound(cause error) *Error {
	return &Error{
		Code:    CodeMediaRulesNotFound,
		Title:   "Media rules not found",
		Details: "No media rules found for this resource",
		cause:   cause,
	}
}

func ErrorMediaRulesAlreadyExists(cause error) *Error {
	return &Error{
		Code:    CodeMediaRulesAlreadyExists,
		Title:   "Media rules already exists",
		Details: "Media rules for this resource already exist",
		cause:   cause,
	}
}

func ErrorFileTooLarge(cause error) *Error {
	return &Error{
		Code:    CodeFileTooLarge,
		Title:   "File too large",
		Details: "Uploaded file exceeds maximum allowed size",
		cause:   cause,
	}
}

func ErrorFileFormatNotAllowed(cause error) *Error {
	return &Error{
		Code:    CodeFileFormatNotAllowed,
		Title:   "File format not allowed",
		Details: "Uploaded file has an unsupported format",
		cause:   cause,
	}
}

func ErrorUserNotAllowedToUploadMedia(cause error) *Error {
	return &Error{
		Code:    CodeUserNotAllowedToUploadMedia,
		Title:   "User not allowed to upload media",
		Details: "You don’t have permission to upload media",
		cause:   cause,
	}
}

func ErrorUserNotAllowedToDeleteMedia(cause error) *Error {
	return &Error{
		Code:    CodeUserNotAllowedToDeleteMedia,
		Title:   "User not allowed to delete media",
		Details: "You don’t have permission to delete this media",
		cause:   cause,
	}
}

func ErrorMediaNotFound(cause error) *Error {
	return &Error{
		Code:    CodeMediaNotFound,
		Title:   "Media not found",
		Details: "Requested media does not exist",
		cause:   cause,
	}
}

func ErrorMediaAlreadyExists(cause error) *Error {
	return &Error{
		Code:    CodeMediaAlreadyExists,
		Title:   "Media already exists",
		Details: "Media with this identifier already exists",
		cause:   cause,
	}
}

func ErrorMediaExtensionNotAllowed(cause error) *Error {
	return &Error{
		Code:    CodeMediaExtensionNotAllowed,
		Title:   "Media extension not allowed",
		Details: "The file extension is not allowed for this resource",
		cause:   cause,
	}
}

func ErrorInternal(cause error) *Error {
	return &Error{
		Code:    CodeInternal,
		Title:   "Internal server error",
		Details: "An unexpected error occurred",
		cause:   cause,
	}
}
