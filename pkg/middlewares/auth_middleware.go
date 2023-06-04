package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/pkg/helpers"
	"github.com/ocintnaf/fameforce/pkg/http"
	"github.com/ocintnaf/fameforce/types"
)

func NewAuthMiddleware(verifier types.IDTokenVerifier) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		verifiedIDToken, err := getAndVerifyIDToken(ctx, verifier)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(http.NewErrorResponse(err))
		}

		ctx.Locals("UserID", verifiedIDToken.UID)

		return ctx.Next()
	}
}

func getAndVerifyIDToken(ctx *fiber.Ctx, verifier types.IDTokenVerifier) (*types.BaseIDToken, error) {
	idToken, err := helpers.GetBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	return verifier.VerifyIDToken(ctx.Context(), idToken)
}
