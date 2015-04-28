package main

import (
	"github.com/gilankpam/osin"
)

type ClaimManager struct {
	UserStorage *TestStorage
}

func (c *ClaimManager) GetClaims(scope string, user interface{}) osin.Claims {
	var claims osin.Claims = make(osin.Claims)
	u := user.(*osin.DefaultUser)
	us, _ := c.UserStorage.GetLocalUser(u.GetSub())
	switch scope {
	case "profile":
		claims["name"] = us.Name
		claims["gender"] = us.Gender
	case "email":
		claims["email"] = us.Email
		claims["email_verified"] = us.EmailVerified
	default:
		claims["name"] = us.Name
		claims["gender"] = us.Gender
	}
	return claims
}

func (c *ClaimManager) AvailableScope() []string {
	return nil
}
