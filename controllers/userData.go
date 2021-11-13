package controllers

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/nimerfarahty/go-rest/config"
	"github.com/nimerfarahty/go-rest/models"
)

func setUserData(c *fiber.Ctx) {
	auth := c.Get("authorization")

	arr := strings.Split(auth, " ")
	if len(arr) != 2 {
		c.Context().SetUserValue("status", "wrong access token or header format")
		return
	}

	tokenString := strings.TrimSpace(arr[1])

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong token format ")
		}
		return []byte(config.Keys.Security.AccessToken.Secret), nil
	})

	if err != nil {
		c.Context().SetUserValue("status", err.Error())
		return
	}

	if !token.Valid {
		c.Context().SetUserValue("status", "token is not valid")
		return
	}

	var user models.UserJWT
	claims := token.Claims.(jwt.MapClaims)
	if err := mapstructure.Decode(claims["data"], &user); err != nil {
		c.Context().SetUserValue("status", "error while decoding payload claim")
		return
	}

	//fmt.Printf("%+v\n", user)

	c.Context().SetUserValue("status", "ok")
	c.Context().SetUserValue("user", user)
}
