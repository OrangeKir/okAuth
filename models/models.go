package models

import (
	"github.com/golang-jwt/jwt"
	"okAuth/types"
)

type AuthInfo struct {
	Login    string
	Password string
}

type CreateUserInfoRequest struct {
	Login    string
	Password string
	Role     RoleType.Role
}

type ChangeUserRoleRequest struct {
	Login    string
	Password string
	Role     RoleType.Role
}

type ChangeUserPasswordRequest struct {
	Login       string
	OldPassword string
	NewPassword string
}

type ValidateTokenRequest struct {
	Token string
}

type ValidateTokenResponse struct {
	IsValid bool
}

type TokenClaims struct {
	jwt.StandardClaims
	Login string
	Role  RoleType.Role
}
