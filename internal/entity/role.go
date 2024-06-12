package entity

import "errors"

type Role struct {
	code string
	name string
}

var (
	RoleUndefined = Role{code: "RoleUndefined", name: "undefined"}
	RoleUser      = Role{code: "RoleUser", name: "user"}
	RoleAdmin     = Role{code: "RoleAdmin", name: "admin"}
)

var roleByName = map[string]Role{
	RoleUser.name:  RoleUser,
	RoleAdmin.name: RoleAdmin,
}

func NewRole(name string) (Role, error) {
	role, ok := roleByName[name]
	if !ok {
		return RoleUndefined, errors.New("unknown role")
	}

	return role, nil
}

func (e Role) Code() string {
	return e.code
}

func (e Role) Name() string {
	return e.name
}
