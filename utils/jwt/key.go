package jwt

import (
	"codename_backend/models"
	// "go/token"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


var JWT_SECRET = os.Getenv("JWT_SECRET")

func GenerateAllTokens(email string, uid string) (signedToken string, signedRefreshToken string, err error){
	claims := models.signedDetails{
		Email: email,
		uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := models.signedDetails{
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Panic(err)
		return
	}
	refreshToken,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SECRET))
	if err != nil {
		log.Panic(err)
		return
	}
	return token,refreshToken,nil

}