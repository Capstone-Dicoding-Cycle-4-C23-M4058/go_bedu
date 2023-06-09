package middlewares

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Create JWTCookieService
func CreateCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "bEDUCookie"
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Path = "/"
	c.SetCookie(cookie)
}

// DeleteCookie deletes the JWT cookie
func DeleteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "bEDUCookie"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	cookie.Path = "/"

	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}
