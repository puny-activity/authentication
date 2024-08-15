package role

import "errors"

type Role struct {
	localizationCode string
	name             string
}

var (
	Undefined = Role{localizationCode: "RoleUndefined", name: "undefined"}
	God       = Role{localizationCode: "RoleGod", name: "god"}
	Admin     = Role{localizationCode: "RoleAdmin", name: "admin"}
	User      = Role{localizationCode: "RoleUser", name: "user"}
	Guest     = Role{localizationCode: "RoleGuest", name: "guest"}
)

var roleByName = map[string]Role{
	God.name:   God,
	Admin.name: Admin,
	User.name:  User,
	Guest.name: Guest,
}

func Parse(name string) (Role, error) {
	role, ok := roleByName[name]
	if !ok {
		return Undefined, errors.New("unknown role")
	}

	return role, nil
}

func (e Role) IsGod() bool {
	return e == God
}

func (e Role) IsAdmin() bool {
	return e == Admin
}

func (e Role) IsUser() bool {
	return e == User
}

func (e Role) IsGuest() bool {
	return e == Guest
}

func (e Role) Code() string {
	return e.localizationCode
}

func (e Role) Name() string {
	return e.name
}
