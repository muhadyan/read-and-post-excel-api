package service

import (
	"excel-read/db"
	"excel-read/model"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckLogin(c echo.Context, username, password string) (bool, error) {
	db := db.DbManager()
	users := model.Users{}

	db.First(&users, "username = ?", username)

	if username != users.Username {
		return false, c.String(http.StatusNotFound, "Username not found")
	}

	match, err := CheckPasswordHash(password, users.Password)
	if !match {
		return false, c.String(http.StatusNotFound, "Hash and password doesn't match")
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func SignUp(c echo.Context, username string, password string) error {
	db := db.DbManager()
	users := model.Users{}

	db.First(&users, "username = ?", username)

	if username == users.Username {
		return c.String(http.StatusConflict, "Username already exist")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.Users{Username: username, Password: string(bytes)}

	db.Create(&user)

	return c.String(http.StatusOK, "Signup Success")
}