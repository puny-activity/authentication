package controller

import (
	"encoding/json"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/base/headerbase"
	"github.com/puny-activity/authentication/pkg/werr"
	"net/http"
)

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) error {
	version := r.Header.Get(headerbase.APIVersion)
	switch version {
	case "1":
		return c.signUpV1(w, r)
	default:
		return errs.InvalidAPIVersion
	}
}

type signUpV1Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Controller) signUpV1(w http.ResponseWriter, r *http.Request) error {
	var signUpRequest signUpV1Request

	err := json.NewDecoder(r.Body).Decode(&signUpRequest)
	if err != nil {
		return werr.WrapES(errs.FailedToDecodeRequestBody, err.Error())
	}

	accountToCreate := account.ToCreate{
		User: account.User{
			Username: signUpRequest.Username,
		},
		Password: signUpRequest.Password,
	}
	err = c.app.AccountUseCase.SignUp(r.Context(), accountToCreate)
	if err != nil {
		return werr.WrapSE("failed to sign up", err)
	}

	return c.responseWriter.Write(w, http.StatusCreated, nil)
}
