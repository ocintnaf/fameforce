package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/dtos"
	"github.com/ocintnaf/fameforce/pkg/helpers"
	"github.com/ocintnaf/fameforce/pkg/http"
	"github.com/ocintnaf/fameforce/usecases"
)

type userController struct {
	userUsecase usecases.UserUsecase
}

// UserController is the interface for the user controller.
type UserController interface {
	Me(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

// NewUserController returns a new user controller.
func NewUserController(userUsecase usecases.UserUsecase) *userController {
	return &userController{userUsecase: userUsecase}
}

func (uc *userController) Me(ctx *fiber.Ctx) error {
	userID, err := helpers.GetCtxLocal[string](ctx, "UserID")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(http.NewErrorResponse(err))
	}

	userDTO, err := uc.userUsecase.GetByID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(http.NewErrorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(http.NewSuccessResponse(userDTO))
}

func (uc *userController) Create(ctx *fiber.Ctx) error {
	userDTO := dtos.UserDTO{}

	if err := ctx.BodyParser(&userDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.NewFailResponse("invalid payload"))
	}

	if validationErrs := helpers.Validate(userDTO); validationErrs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.NewFailResponse(validationErrs))
	}

	userID := ctx.Locals("UserID").(string)
	userDTO.ID = userID

	createdUserDTO, err := uc.userUsecase.Create(userDTO)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(http.NewErrorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(http.NewSuccessResponse(createdUserDTO))
}
