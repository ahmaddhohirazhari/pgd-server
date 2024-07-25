package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"pgd-server.com/helpers"
)

type AuthMiddleware struct{}

func (am *AuthMiddleware) GetJWTSecret() string {
	ev := helpers.Environtment()
	return ev.JwtSecretKey
}

func (am *AuthMiddleware) JwtTokenCheck(c *gin.Context) {
	success, _, _ := am.TokenCheckMiddleware(c)
	if !success {
		c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.GeneralResponse{
			Message: "Token tidak valid",
			Status:  false,
		})
		return
	}
	c.Next()
}

func (am *AuthMiddleware) TokenCheckMiddleware(c *gin.Context) (bool, jwt.Claims, error) {
	jwtToken, err := am.ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		return false, nil, err
	}

	token, err := am.ParseToken(jwtToken)
	if err != nil {
		return false, nil, err
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		return false, nil, err
	}

	c.Set("userId", claims["userId"])
	c.Set("roleId", claims["roleId"])
	c.Set("companyId", claims["companyId"])
	c.Set("branchId", claims["branch_id"])
	c.Set("subbranchId", claims["subbranch_id"])
	c.Set("affiliate_code", claims["affiliate_code"])

	return true, claims, err
}

func (am *AuthMiddleware) ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func (am *AuthMiddleware) ParseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(am.GetJWTSecret()), nil
	})

	if err != nil {
		return nil, fmt.Errorf("mengurai token tidak valid")
	}

	return token, nil
}
