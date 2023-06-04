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

type UserController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

func NewUserController(userUsecase usecases.UserUsecase) *userController {
	return &userController{userUsecase: userUsecase}
}

func (uc *userController) GetAll(ctx *fiber.Ctx) error {
	userDTOs, err := uc.userUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(http.NewErrorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(http.NewSuccessResponse(fiber.Map{
		"users": userDTOs,
	}))
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

	return ctx.Status(fiber.StatusOK).JSON(http.NewSuccessResponse(fiber.Map{
		"user": createdUserDTO,
	}))
}
