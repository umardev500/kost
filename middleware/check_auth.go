package auth

import (
	"context"
	"os"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/ekost/constants"
	"github.com/umardev500/kost/utils"
)

type checkAuthMiddleware struct {
	storage fiber.Storage
}

// NewCheckAuth creates a new checkAuthMiddleware instance.
//
// It takes a fiber.Storage parameter and returns a pointer to a checkAuthMiddleware struct.
func NewCheckAuth(s fiber.Storage) *checkAuthMiddleware {
	return &checkAuthMiddleware{
		storage: s,
	}
}

// Start returns a fiber.Handler that is responsible for handling the authentication middleware.
//
// It creates a new JWT middleware with the provided configuration, including a success handler
// and a signing key derived from the "SECRET" environment variable. The returned handler can
// be used to authenticate requests.
//
// Return:
//   - fiber.Handler: The authentication middleware handler.
func (cam *checkAuthMiddleware) Start() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: cam.CheckAuthSuccess,
		SigningKey: jwtware.SigningKey{
			Key: []byte(os.Getenv("SECRET")),
		},
	})
}

// CheckAuthSuccess checks the authorization of the request and proceeds with the next handler if authorization is successful.
//
// Parameters:
// - c: The fiber.Ctx object representing the current HTTP request context.
//
// Return:
// - err: An error object, if any error occurs during the authorization check.
func (cam *checkAuthMiddleware) CheckAuthSuccess(c *fiber.Ctx) (err error) {
	bearer := c.Get("Authorization")
	claims := utils.GetMapclaims(c)
	username := claims["username"].(string)
	id := claims["id"].(string)
	ctx := context.WithValue(c.Context(), constants.CtxKeyTx, bearer)
	err = cam.CheckAuthService(ctx, username)
	if err != nil {
		return utils.ErrorHandler(c, err)
	}

	c.Locals("username", username)
	c.Locals("id", id)

	return c.Next()
}

// CheckAuthService checks the authentication service for a given user.
//
// It takes a context.Context object and a string parameter representing the user.
// It returns an error.
func (cam *checkAuthMiddleware) CheckAuthService(ctx context.Context, u string) (err error) {
	log.Debug().Msgf("Starting check auth for user=%s", u)

	bearer := ctx.Value(constants.CtxKeyTx).(string)
	currentToken := strings.Split(bearer, " ")[1]

	token, err := cam.storage.Get(u)
	if err != nil {
		id := uuid.New()
		log.Error().Msgf("Error getting token [%s]", id)
		return
	}

	if string(token) != currentToken {
		id := uuid.New()
		newErr := utils.NewError()
		newErr.StatusCode = fiber.StatusUnauthorized
		newErr.Message = fiber.ErrUnauthorized.Message
		log.Error().Msgf("Can not access resources, invalid token user=%s [%s]", u, id)
		return newErr
	}

	if token == nil {
		id := uuid.New()
		newErr := utils.NewError()
		newErr.StatusCode = fiber.StatusNotFound
		newErr.Message = fiber.ErrNotFound.Message
		log.Error().Msgf("Token not found user=%s [%s]", u, id)
		return newErr
	}

	log.Debug().Msgf("Check auth success for user=%s", u)

	return
}
