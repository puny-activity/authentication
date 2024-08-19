package controller

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/base/headerbase"
	"github.com/puny-activity/authentication/pkg/werr"
	"net/http"
)

func (c Controller) Refresh(w http.ResponseWriter, r *http.Request) error {
	version := r.Header.Get(headerbase.APIVersion)
	switch version {
	case "1":
		return c.refreshV1(w, r)
	default:
		return errs.InvalidAPIVersion
	}
}

type refreshV1Request struct {
	RefreshToken *string `json:"refreshToken"`
}

type refreshV1Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c Controller) refreshV1(w http.ResponseWriter, r *http.Request) error {
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

	tokenPair, err := c.app.AccountUseCase.Refresh(r.Context(), *refreshToken)
	if err != nil {
		return werr.WrapSE("failed to sign up", err)
	}

	response := signInV1Response{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Expires:  carbon.Now().AddSeconds(c.cfg.App.RefreshToken.TTLSecond()).ToStdTime(),
		HttpOnly: true,
		Secure:   true,
	})

	return c.responseWriter.Write(w, http.StatusCreated, response)
}
