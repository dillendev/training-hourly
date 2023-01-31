package main

import (
	"strings"

	"github.com/labstack/echo/v4"

	hourly "github.com/dillendev/training-hourly"
)

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if ctx.Request().URL.Path != "/api/auth/tokens" {
				auth := ctx.Request().Header.Get("Authorization")
				if !strings.HasPrefix(auth, "Bearer ") {
					return ctx.JSON(400, hourly.Error{Message: "invalid authorization header"})
				}

				token := auth[7:]
				if err := verifyToken(token); err != nil {
					return ctx.JSON(401, hourly.Error{Message: err.Error()})
				}
			}

			return next(ctx)
		}
	})

	hourly.RegisterHandlers(e, &server{})

	if err := e.Start(":8989"); err != nil {
		panic(err)
	}
}
