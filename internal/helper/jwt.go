package helper

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super-secret-key")

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int64) (string, error) {
	// payload
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyJWT(token string) (int64, error) {
	claims := &Claims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, jwt.ErrTokenUnverifiable
		}

		return secretKey, nil
	})

	if err != nil {
		log.Print("error parsing token")
		return 0, err
	}

	if !parsedToken.Valid {
		log.Print("parsed token invalid")
		return 0, jwt.ErrTokenSignatureInvalid
	}

	if claims.UserId == 0 {
		return 0, jwt.ErrTokenInvalidClaims
	}

	return claims.UserId, nil
}
