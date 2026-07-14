package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

const AuthCookieName = "auth_token"
const CookieMaxAge = 2 * time.Hour

func SetAuthCookie(c fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     AuthCookieName,
		Value:    token,
		HTTPOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(CookieMaxAge),
		SameSite: "Lax",
	})
}

func ClearAuthCookie(c fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     AuthCookieName,
		Value:    "",
		HTTPOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})
}

func GetAuthToken(c fiber.Ctx) string {
	token := c.Cookies(AuthCookieName)
	if token != "" {
		return token
	}

	auth := c.Get("Authorization")
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}

	return ""
}
