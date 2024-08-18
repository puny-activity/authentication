package credential

import (
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
)

type Credential struct {
	Email    email.Email
	Password password.Password
}
