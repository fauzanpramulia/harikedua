package middleware

import (
	"net/http"
	"harikedua/helpers"
	"github.com/labstack/echo/v4"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error{
			verifyToken, err := helpers.VerifyToken(ctx)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":	"Unauthenticated",
					"message": 	err.Error(),
				})
			}
			// Log or print verifyToken for debugging
			// fmt.Println(verifyToken["id"|)|
			ctx.Set ("userData", verifyToken)
			return next(ctx)
		}
	}
}