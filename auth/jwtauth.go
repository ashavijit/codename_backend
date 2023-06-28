package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SECRET = []byte("codename_backend")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string , error){
	claims:= &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000, // 15 seconds
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	Token_string, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}
	return Token_string, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	token , err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil , err
	}
	return claims, nil
}

func JWT_BASIC_AUTH() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		tokenString := authHeader[7:] // remove "Bearer " from the header
		claims, err := VerifyJWT(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Header("username", claims.Username)
		ctx.Next()
	}
}