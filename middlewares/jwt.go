package middlewares

import (
	"errors"
	"quiz3/commons"
	"quiz3/commons/responses"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := getTokenFromHeader(ctx)

		if err != nil {
			responses.GenerateUnauthorizedResponse(ctx, err.Error())

			return
		}

		token, tokenValidationError := verifyToken(tokenString)

		if tokenValidationError != nil {
			responses.GenerateUnauthorizedResponse(ctx, tokenValidationError.Error())

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); !ok {
			responses.GenerateUnauthorizedResponse(ctx, "invalid token")

			return
		} else {
			ctx.Set("user", claims)
		}

		ctx.Next()
	}
}

func CreateToken(id int, username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":      id,
			"username": username,
			"iss":      "booksApiServer",
			"aud":      "booksApiClient",
			"exp":      time.Now().Add(time.Hour).Unix(),
			"iat":      time.Now().Unix(),
		})

	token, err := claims.SignedString(commons.JWT_SECRET_KEY)

	if err != nil {
		return "", err
	}

	return token, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return commons.JWT_SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return token, nil
}

func getTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return authHeader, errors.New("authorization header is required")
	}

	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 || splitAuthHeader[0] != "Bearer" {
		return authHeader, errors.New("invalid authorization header format")
	}

	return splitAuthHeader[1], nil
}

// return is for : id, username, error
func GetClaims(ctx *gin.Context) (int, string, error) {
	claims, exists := ctx.Get("user")

	if !exists {
		return 0, "", errors.New("invalid authorization header format")
	}

	mapClaims := claims.(jwt.MapClaims)
	username := mapClaims["username"].(string)
	id := int(mapClaims["sub"].(float64))

	return id, username, nil
}
