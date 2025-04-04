package tokenverify

import (
	"fmt"
	"testing"

	"github.com/eos/protocol/constant"
	"github.com/golang-jwt/jwt/v4"
)

var secret = "EOS_server"

func Test_ParseToken(t *testing.T) {
	fmt.Print("111")
	original_claims := BuildClaims("12345678",constant.AndroidPadPlatformID,10)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, original_claims)
	tokenString,err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatal(err)
	}
	now_claims ,err := getClaimsFromToken(tokenString,secretFun())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(now_claims)
	t.Log(fmt.Sprintf("uid %s platformID %d",now_claims.UserID,now_claims.PlatformID))
}
func secretFun() jwt.Keyfunc {
	return func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}
}