package controller

import (
	"encoding/json"
	"github.com/puny-activity/authentication/internal/entity"
	"github.com/puny-activity/authentication/internal/interr"
	"github.com/puny-activity/authentication/pkg/werr"
	"net/http"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) error {
	var signUpRequest SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&signUpRequest)
	if err != nil {
		return werr.WrapEE(interr.FailedToDecodeRequestBody, err)
	}

	accountToCreate := entity.AccountCreateRequest{
		Account: entity.Account{
			Username: signUpRequest.Username,
		},
		Password: signUpRequest.Password,
	}
	err = c.app.AccountUseCase.SignUp(r.Context(), accountToCreate)
	if err != nil {
		return err
	}

	return c.responseWriter.Write(w, http.StatusCreated, nil)
}
