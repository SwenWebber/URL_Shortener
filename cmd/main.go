package main

import (
	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	

	
}

func GenerateString(size int) string {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, size)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)

}
