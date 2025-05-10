package middleware

import (
	"PetAi/pkg/apperror"
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"strings"

	"PetAi/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type Middleware struct {
	publicKey *rsa.PublicKey
}

func InitJwtMiddleware(publicKeyPass string) *Middleware {
	return NewMiddleware(publicKeyPass)
}

func NewMiddleware(publicKeyPath string) *Middleware {
	keyData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read public key: %v")
	}

	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid public key format: %v")
	}

	return &Middleware{publicKey: parsedKey}
}

func (m *Middleware) AuthRequired(roles ...auth.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr, err := extractToken(c)
		if err != nil {
			return apperror.Unauthorized(err)
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, apperror.InternalServerError(errors.New("unexpected signing method"))
			}
			return m.publicKey, nil
		})

		if err != nil || !token.Valid {
			return apperror.Unauthorized(err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return apperror.Unauthorized(err)
		}

		userID, _ := claims["user_id"].(string)
		roleStr, _ := claims["role"].(string)
		role := auth.Role(roleStr)

		c.Locals("user_id", userID)
		c.Locals("role", role)

		if len(roles) > 0 && !roleAllowed(role, roles) {
			return fiber.ErrForbidden
		}

		return c.Next()
	}
}

func extractToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("missing bearer token")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func roleAllowed(role auth.Role, allowed []auth.Role) bool {
	for _, r := range allowed {
		if r == role {
			return true
		}
	}
	return false
}
