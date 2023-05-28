package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/pkg/http"
	"github.com/ocintnaf/fameforce/usecases"
)

type influencerController struct {
	router            fiber.Router
	influencerUsecase usecases.InfluencerUsecase
}

type InfluencerController interface {
	GetAll(ctx *fiber.Ctx) error
}

func NewInfluencerController(
	router fiber.Router,
	influencerUsecase usecases.InfluencerUsecase,
) *influencerController {
	return &influencerController{
		router:            router,
		influencerUsecase: influencerUsecase,
	}
}

func (ic *influencerController) GetAll(ctx *fiber.Ctx) error {
	influencerDTOs, err := ic.influencerUsecase.GetAll(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(http.NewHttpResponse(nil, err))
	}

	return ctx.Status(fiber.StatusOK).JSON(http.NewHttpResponse(influencerDTOs, nil))
}

func (ic *influencerController) RegisterRoutes() {
	ic.router.Get("/", ic.GetAll)
}
