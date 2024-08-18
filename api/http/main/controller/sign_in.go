package controller

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/entity/account/credential"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/base/headerbase"
	"github.com/puny-activity/authentication/pkg/werr"
	"net/http"
)

func (c Controller) SignIn(w http.ResponseWriter, r *http.Request) error {
	version := r.Header.Get(headerbase.APIVersion)
	switch version {
	case "1":
		return c.signInV1(w, r)
	default:
		return errs.InvalidAPIVersion
	}
}

type signInV1Request struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DeviceName  string `json:"deviceName"`
	Fingerprint string `json:"fingerprint"`
}

type signInV1Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c Controller) signInV1(w http.ResponseWriter, r *http.Request) error {
	var signInRequest signInV1Request

	err := json.NewDecoder(r.Body).Decode(&signInRequest)
	if err != nil {
		return werr.WrapES(errs.FailedToDecodeRequestBody, err.Error())
	}

	signInEmail, err := email.New(signInRequest.Email)
	if err != nil {
		return werr.WrapSE("failed to construct email", err)
	}

	signInPassword, err := password.New(signInRequest.Password)
	if err != nil {
		return werr.WrapSE("failed to construct email", err)
	}

	credentials := credential.Credential{
		Email:    signInEmail,
		Password: signInPassword,
	}

	sourceDevice := device.Device{
		Name:        signInRequest.DeviceName,
		Fingerprint: signInRequest.Fingerprint,
	}

	tokenPair, err := c.app.AccountUseCase.SignIn(r.Context(), credentials, sourceDevice)
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
