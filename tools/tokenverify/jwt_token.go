package tokenverify

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/eos/tools/errs"
)

const HoursOneDay = 24
const secondBefore  = 5


type Claims struct {
	UserID string
	PlatformID int
	jwt.RegisteredClaims
}

func BuildClaims(uid string ,paltformID int,ttl int64) Claims {
	now := time.Now()
	before := now.Add(-time.Second * time.Duration(secondBefore))
	return Claims{
		UserID: uid,
		PlatformID: paltformID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl * HoursOneDay) * time.Hour)),
			IssuedAt: jwt.NewNumericDate(before),
		},
	}
}
func getClaimsFromToken(tokenString string,secretFunc jwt.Keyfunc) (*Claims,error) {
	token,err  := jwt.ParseWithClaims(tokenString, &Claims{}, secretFunc)
	if err == nil {
		if claims,ok := token.Claims.(*Claims); ok && token.Valid {
			return claims,nil
		}
		return nil,errs.WrapMsg(errs.ErrTokenUnknown,"claims unknown","token",tokenString)
	}
	if validError,ok := err.(*jwt.ValidationError) ;ok {
		return nil,errs.WrapMsg(mapValidationError(validError),"jwt parse error","token",tokenString)
	}
	return nil,errs.WrapMsg(err,"jwt parse error","token",tokenString)
	
}

func mapValidationError(err *jwt.ValidationError) error {
	if err.Errors & jwt.ValidationErrorMalformed != 0 {
		return	errs.ErrTokenMalformed 
	}else if err.Errors & jwt.ValidationErrorExpired != 0 {
		return errs.ErrTokenExpired
	}else if err.Errors & jwt.ValidationErrorNotValidYet != 0 {
		return errs.ErrTokenNotValidYet
	}
	return errs.NewCodeError(errs.TokenUnknownError,err.Error())
}