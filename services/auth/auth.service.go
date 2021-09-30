package authService

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/nimerfarahty/go-rest/config"
	"github.com/nimerfarahty/go-rest/database"
	"github.com/nimerfarahty/go-rest/models"
)

func Login(loginInput *models.LoginInput) (*models.LoginResponse, error) {

	var user *models.User

	tx := database.Conn.Where("email = ? or phone = ?", loginInput.Identity, loginInput.Identity).First(&user)

	if tx.RowsAffected != 1 {
		return nil, errors.New("no user found")
	}
	println(*user.Password, loginInput.Password)
	if !user.CheckPassword(loginInput.Password) {
		return nil, errors.New("incorrect password")
	}

	return successLogin(user)
}

func Logout() {

}

func RefresTokens() {

}

func successLogin(user *models.User) (*models.LoginResponse, error) {
	conf := config.Keys.Security

	refreshHandle := uuid.NewString()
	refreshToken, err := createToken(
		refreshHandle,
		conf.RefreshToken.Secret,
		conf.RefreshToken.Expiry,
	)

	if err != nil {
		return nil, err
	}

	accessToken, err := createToken(
		strconv.Itoa(user.ID),
		conf.AccessToken.Secret,
		conf.AccessToken.Expiry,
	)

	if err != nil {
		return nil, err
	}

	user.Token = &refreshHandle

	tx := database.Conn.Save(&user)

	if tx.RowsAffected != 1 {
		return nil, errors.New("error while updating refresh handle")
	}

	return &models.LoginResponse{Tokens: &models.Tokens{AccessToken: *accessToken, RefreshToken: *refreshToken}, User: user}, nil

}

func createToken(sub string, secret string, expiry string) (*string, error) {

	duration, err := time.ParseDuration(expiry)

	if err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(duration).Unix()

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}
