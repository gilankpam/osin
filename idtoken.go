package osin

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// generate id token for user U with Client C
// func GenerateIDToken(user *User, client *Client) string {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	token.Claims["iss"] =
// }

// http://openid.net/specs/openid-connect-core-1_0.html#IDToken
func (s *Server) generateIDToken(ar *AccessRequest) (string, error) {
	user := ar.UserData.(*DefaultUser)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["iss"] = s.Config.Issuer
	token.Claims["sub"] = user.GetSub()
	token.Claims["aud"] = ar.Client.GetId()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["exp"] = time.Now().Add(time.Duration(s.Config.IDTokenExpiration) * time.Second).Unix()
	return token.SignedString(s.Config.JWTKey)
}
