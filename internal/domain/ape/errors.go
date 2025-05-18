package ape

import "fmt"

type Error struct {
	Code    string
	Title   string
	Details string
	cause   error
}

func (e *Error) Unwrap() error {
	return e.cause
}

func ErrorMediaNotFoundByPath(path string, err error) *Error {
	return &Error{
		Code:    CodeMediaNotFound,
		Title:   "Media Not Found",
		Details: fmt.Sprintf("The requested media %s does not exist.", path),
		cause:   err,
	}
}

func ErrorMediaRulesNotFound(id string, err error) *Error {
	return &Error{
		Code:    CodeMediaRulesNotFound,
		Title:   "Media Rules Not Found",
		Details: fmt.Sprintf("The requested media rules %s does not exist.", id),
		cause:   err,
	}
}

func ErrorMediaRulesAlreadyExists(cause error) *Error {
	return &Error{
		Code:    CodeMediaRulesAlreadyExists,
		Title:   "Media Rules Already Exists",
		Details: "Media rules already exists",
		cause:   cause,
	}
}

func ErrorFileToLarge(cause error) *Error {
	return &Error{
		Code:    CodeFileToLarge,
		Title:   "File Too Large",
		Details: "The file is too large.",
		cause:   cause,
	}
}

func ErrorFileFormatNotAllowed(cause error) *Error {
	return &Error{
		Code:    CodeFileFormatNotAllowed,
		Title:   "File Format Not Allowed",
		Details: "The file format is not allowed.",
		cause:   cause,
	}
}

func ErrorUserNotAllowedToUploadMedia(cause error) *Error {
	return &Error{
		Code:    CodeUserNotAllowedToUploadMedia,
		Title:   "User Not Allowed To Upload Media",
		Details: "The user is not allowed to upload media.",
		cause:   cause,
	}
}

func ErrorUserNotAllowedToDeleteMedia(cause error) *Error {
	return &Error{
		Code:    CodeUserNotAllowedToDeleteMedia,
		Title:   "User Not Allowed To Delete Media",
		Details: "The user is not allowed to delete media.",
		cause:   cause,
	}
}

func ErrorMediaAlreadyExists(cause error) *Error {
	return &Error{
		Code:    CodeMediaAlreadyExists,
		Title:   "Media Already Exists",
		Details: "Media already exists",
		cause:   cause,
	}
}
