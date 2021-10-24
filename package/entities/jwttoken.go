package entities

import (
	"crypto/md5"
	"encoding/hex"
	"golang-ecommerce-practice/errs"
	"golang-ecommerce-practice/zapLog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	jwtKey string = "MY-LONG-KEY-FOR-JWT"
)

type ClaimsJwt struct {
	UserToken
	jwt.StandardClaims
}

func PasswordEncrypt(pass *string) string {
	hash := md5.Sum([]byte(*pass + "salt (optional)"))
	return hex.EncodeToString(hash[:])
}

func CreateJwt(usr *UserToken) (*string, *errs.AppError) {
	claims := ClaimsJwt{
		*usr,
		jwt.StandardClaims{
			ExpiresAt: int64(30 * time.Minute),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ret, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		zapLog.Error(err.Error())
		return nil, errs.NewAppError("internal server error", http.StatusInternalServerError)
	}

	return &ret, nil
}

func ValidateJwt(jwtstr *string) (*UserToken, *errs.AppError) {
	tkn, err := jwt.ParseWithClaims(*jwtstr, &ClaimsJwt{}, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if claims, success := tkn.Claims.(*ClaimsJwt); success && tkn.Valid {
		return &claims.UserToken, nil
	} else if fail, success := err.(*jwt.ValidationError); success {
		if fail.Errors&jwt.ValidationErrorMalformed != 0 {
			zapLog.Info("Token not even token")
			return nil, errs.NewAppError("token not valid", http.StatusUnauthorized)
		} else if fail.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			zapLog.Info("Token Expired")
			return nil, errs.NewAppError("Token Expired", http.StatusUnauthorized)
		} else {
			zapLog.Error("Token could't be handle" + err.Error())
			return nil, errs.NewAppError("token not valid", http.StatusUnauthorized)
		}
	} else {
		zapLog.Error("Token errror couldn't be handle" + err.Error())
		return nil, errs.NewAppError("token not valid", http.StatusUnauthorized)
	}
}
