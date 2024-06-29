package role

import "errors"

type Role struct {
	localizationCode string
	name             string
}

var (
	Undefined = Role{localizationCode: "", name: "undefined"}
	Admin     = Role{localizationCode: "RoleAdmin", name: "admin"}
	User      = Role{localizationCode: "RoleUser", name: "user"}
)

var roleByName = map[string]Role{
	Admin.name: Admin,
	User.name:  User,
}

func Parse(name string) (Role, error) {
	role, ok := roleByName[name]
	if !ok {
		return Undefined, errors.New("unknown role")
	}

	return role, nil
}

func (e Role) IsAdmin() bool {
	return e == Admin
}

func (e Role) IsUser() bool {
	return e == User
}

func (e Role) Code() string {
	return e.localizationCode
}

func (e Role) Name() string {
	return e.name
}
