package controller

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/base/headerbase"
	"github.com/puny-activity/authentication/pkg/werr"
	"net/http"
)

func (c Controller) SignOut(w http.ResponseWriter, r *http.Request) error {
	version := r.Header.Get(headerbase.APIVersion)
	switch version {
	case "1":
		return c.signOutV1(w, r)
	default:
		return errs.InvalidAPIVersion
	}
}

type signOutV1Request struct {
	RefreshToken *string `json:"refreshToken"`
}

type signOutV1Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c Controller) signOutV1(w http.ResponseWriter, r *http.Request) error {
	var refreshToken *string = nil
	refreshTokenFromCookie, err := r.Cookie("refresh_token")
	if err == nil {
		refreshToken = &refreshTokenFromCookie.Value
	} else {
		var refreshRequest refreshV1Request
		err := json.NewDecoder(r.Body).Decode(&refreshRequest)
		if err != nil {
			return werr.WrapES(errs.FailedToDecodeRequestBody, err.Error())
		}
		if refreshRequest.RefreshToken != nil {
			refreshToken = refreshRequest.RefreshToken
		}
	}

	if refreshToken == nil {
		return errs.EmptyRefreshToken
	}

	err = c.app.AccountUseCase.SignOut(r.Context(), *refreshToken)
	if err != nil {
		return werr.WrapSE("failed to sign up", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  carbon.Now().SubHour().ToStdTime(),
		HttpOnly: true,
		Secure:   true,
	})

	return c.responseWriter.Write(w, http.StatusOK, nil)
}
