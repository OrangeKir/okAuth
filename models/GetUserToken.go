package models

import "okAuth/types"

type AuthInfo struct {
	Login    string
	Password string
}

type CreateUserInfo struct {
	Login    string
	Password string
	Role     RoleType.Role
}

type ChangeUserRole struct {
	Login string
	role  RoleType.Role
}

type ChangeUserPassword struct {
	Login       string
	OldPassword string
	NewPassword string
}
