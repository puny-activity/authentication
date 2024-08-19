package errs

import (
	"errors"
	"net/http"
)

var (
	// Authentication
	EmptyAccessToken    = errors.New("empty access token")
	WrongPassword       = errors.New("wrong password")
	EmptyRefreshToken   = errors.New("empty refresh token")
	InvalidRefreshToken = errors.New("invalid access token")

	// Request data
	FailedToDecodeRequestBody = errors.New("failed to decode request body")
	NotProvidedEmail          = errors.New("not provided email")
	NotProvidedNickname       = errors.New("not provided nickname")
	NotProvidedPassword       = errors.New("not provided password")
	InvalidEmail              = errors.New("invalid email")
	InvalidPassword           = errors.New("invalid password")
	InvalidAPIVersion         = errors.New("invalid api version")

	// Conflict
	EmailAlreadyTaken    = errors.New("email already taken")
	NicknameAlreadyTaken = errors.New("nickname already taken")

	// Database
	DatabaseFailedToExecuteQuery = errors.New("failed to execute query")

	// Entities
	DatabaseUndefinedID       = errors.New("undefined id")
	DatabaseFailedToParseRole = errors.New("failed to parse role")
	RefreshTokenNotFound      = errors.New("refresh token not found")

	// Unexpected
	Unexpected = errors.New("unexpected error")
)

var unexpectedError = internalError{
	error:            Unexpected,
	localizationCode: "SRV-1",
	httpStatusCode:   http.StatusInternalServerError,
}

var errorList = []internalError{
	{
		error:            EmptyAccessToken,
		localizationCode: "ATH-1",
		httpStatusCode:   http.StatusUnauthorized,
	},
	{
		error:            WrongPassword,
		localizationCode: "ATH-2",
		httpStatusCode:   http.StatusUnauthorized,
	},
	{
		error:            EmptyRefreshToken,
		localizationCode: "ATH-3",
		httpStatusCode:   http.StatusUnauthorized,
	},
	{
		error:            InvalidRefreshToken,
		localizationCode: "ATH-4",
		httpStatusCode:   http.StatusUnauthorized,
	},

	{
		error:            FailedToDecodeRequestBody,
		localizationCode: "RDT-1",
		httpStatusCode:   http.StatusBadRequest,
	},

	{
		error:            NotProvidedEmail,
		localizationCode: "RDT-2",
		httpStatusCode:   http.StatusBadRequest,
	},
	{
		error:            NotProvidedPassword,
		localizationCode: "RDT-3",
		httpStatusCode:   http.StatusBadRequest,
	},

	{
		error:            InvalidEmail,
		localizationCode: "RDT-4",
		httpStatusCode:   http.StatusBadRequest,
	},
	{
		error:            InvalidPassword,
		localizationCode: "RDT-5",
		httpStatusCode:   http.StatusBadRequest,
	},
	{
		error:            NotProvidedNickname,
		localizationCode: "RDT-6",
		httpStatusCode:   http.StatusBadRequest,
	},
	{
		error:            InvalidAPIVersion,
		localizationCode: "RDT-7",
		httpStatusCode:   http.StatusBadRequest,
	},

	{
		error:            EmailAlreadyTaken,
		localizationCode: "CFL-1",
		httpStatusCode:   http.StatusBadRequest,
	},
	{
		error:            NicknameAlreadyTaken,
		localizationCode: "CFL-2",
		httpStatusCode:   http.StatusBadRequest,
	},

	{
		error:            DatabaseFailedToExecuteQuery,
		localizationCode: "DTB-1",
		httpStatusCode:   http.StatusInternalServerError,
	},

	{
		error:            DatabaseFailedToParseRole,
		localizationCode: "ENT-1",
		httpStatusCode:   http.StatusInternalServerError,
	},
	{
		error:            DatabaseUndefinedID,
		localizationCode: "ENT-2",
		httpStatusCode:   http.StatusInternalServerError,
	},
	{
		error:            RefreshTokenNotFound,
		localizationCode: "ENT-3",
		httpStatusCode:   http.StatusInternalServerError,
	},
}
