package middleware

import (
	"strings"

	"github.com/go/mini_market/src/model"
	"github.com/go/mini_market/src/server"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	NO_AUTH_NEEDE = []string{
		"login",
		"signup",
		"home",
	}
)

func shouldCheckToke(route string) bool {
	for _, p := range NO_AUTH_NEEDE {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			if !shouldCheckToke(c.Path()) {
				next(c)
			}
			tokenString := strings.TrimSpace(c.Request().Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, &model.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				return echo.ErrUnauthorized
			}
			return next(c)
		})
	}
}
