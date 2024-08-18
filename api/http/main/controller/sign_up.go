package controller

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/entity/role"
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
	Email    string  `json:"email"`
	Nickname *string `json:"nickname"`
	Password string  `json:"password"`
}

func (c Controller) signUpV1(w http.ResponseWriter, r *http.Request) error {
	var signUpRequest signUpV1Request

	err := json.NewDecoder(r.Body).Decode(&signUpRequest)
	if err != nil {
		return werr.WrapES(errs.FailedToDecodeRequestBody, err.Error())
	}

	signUpEmail, err := email.New(signUpRequest.Email)
	if err != nil {
		return werr.WrapSE("failed to construct email", err)
	}

	signUpNickname := signUpEmail.String()
	if signUpRequest.Nickname != nil {
		signUpNickname = *signUpRequest.Nickname
	}

	accountToCreate := account.Account{
		Email:     signUpEmail,
		Nickname:  signUpNickname,
		Role:      role.Undefined,
		CreatedAt: carbon.Now(),
	}

	passwordToCreate, err := password.New(signUpRequest.Password)
	if err != nil {
		return werr.WrapSE("failed to construct password", err)
	}

	err = c.app.AccountUseCase.Create(r.Context(), accountToCreate, passwordToCreate)
	if err != nil {
		return werr.WrapSE("failed to sign up", err)
	}

	return c.responseWriter.Write(w, http.StatusCreated, nil)
}
