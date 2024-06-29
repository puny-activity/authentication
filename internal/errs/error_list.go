package errs

import (
	"errors"
	"net/http"
)

var (
	// Authentication
	EmptyAccessToken = errors.New("empty access token")

	// Request data
	FailedToDecodeRequestBody = errors.New("failed to decode request body")
	NotProvidedUsername       = errors.New("not provided username")
	NotProvidedNickname       = errors.New("not provided nickname")
	NotProvidedPassword       = errors.New("not provided password")
	InvalidUsername           = errors.New("invalid username")
	InvalidPassword           = errors.New("invalid password")
	InvalidAPIVersion         = errors.New("invalid api version")

	// Conflict
	UsernameAlreadyTaken = errors.New("username already taken")

	// Database
	DatabaseFailedToExecuteQuery = errors.New("failed to execute query")

	// Entities
	DatabaseFailedToParseRole = errors.New("failed to parse role")

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
		error:            FailedToDecodeRequestBody,
		localizationCode: "RDT-1",
		httpStatusCode:   http.StatusBadRequest,
	},

	{
		error:            NotProvidedUsername,
		localizationCode: "RDT-2",
		httpStatusCode:   http.StatusBadRequest,
	},
	{
		error:            NotProvidedPassword,
		localizationCode: "RDT-3",
		httpStatusCode:   http.StatusBadRequest,
	},

	{
		error:            InvalidUsername,
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
		error:            UsernameAlreadyTaken,
		localizationCode: "CFL-1",
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
}
