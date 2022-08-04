package common

import (
	"commerce-app/config"
	"commerce-app/domain"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GenerateToken(ID int) string {
	info := jwt.MapClaims{}
	info["ID"] = ID
	info["Role"] = "admin"
	info["exp "] = time.Now().UnixNano()
	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	token, err := auth.SignedString([]byte(config.SECRET))
	if err != nil {
		log.Fatal("cannot generate key")
		return ""
	}
	return token
}

func ExtractData(c echo.Context) domain.User {
	var userdata domain.User

	head := c.Request().Header
	token := strings.Split(head.Get("Authorization"), " ")

	res, _ := jwt.Parse(token[len(token)-1], func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})
	if res.Valid {
		resClaim := res.Claims.(jwt.MapClaims)
		parseID := resClaim["ID"].(float64)
		userdata.ID = int(parseID)
		userdata.Role = resClaim["Role"].(string)
		return userdata
	}
	return domain.User{}
}

func ExtractData_Admin(c echo.Context) (int, string) {
	head := c.Request().Header
	token := strings.Split(head.Get("Authorization"), " ")

	res, _ := jwt.Parse(token[len(token)-1], func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})
	if res.Valid {
		resClaim := res.Claims.(jwt.MapClaims)
		parseID := resClaim["ID"].(float64)
		parseRole := resClaim["Role"].(string)
		return int(parseID), parseRole
	}
	return -1, ""
}

func UseJWT(secret []byte) middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    secret,
	}
}
