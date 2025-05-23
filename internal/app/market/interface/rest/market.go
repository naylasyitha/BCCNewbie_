package rest

import (
	"hackfest-uc/internal/app/market/usecase"
	"hackfest-uc/internal/domain/dto"
	"hackfest-uc/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MarketHandler struct {
	MarketUsecase usecase.MarketUsecaseItf
	Middleware    middleware.MiddlewareItf
}

func NewMarketHandler(routerGroup fiber.Router, marketUsecase usecase.MarketUsecaseItf, middleware middleware.MiddlewareItf) {
	MarketHandler := MarketHandler{
		MarketUsecase: marketUsecase,
		Middleware:    middleware,
	}

	marketGroup := routerGroup.Group("/markets")
	marketGroup.Post("/", MarketHandler.Middleware.Authentication, MarketHandler.CreateProduct)
}

func (h *MarketHandler) CreateProduct(ctx *fiber.Ctx) error {
	var product dto.CreateProduct

	// Parse form data including file upload
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse form data",
		})
	}

	// Parse other fields
	product.ProductName = form.Value["product_name"][0]
	product.ProductPrice, _ = strconv.ParseUint(form.Value["product_price"][0], 10, 64)
	product.ProductWeight, _ = strconv.ParseUint(form.Value["product_weight"][0], 10, 64)
	product.ProductType = dto.Jenis(form.Value["product_type"][0])
	product.ProductUsage = dto.Kegunaan(form.Value["product_usage"][0])
	product.Composition = dto.Composition(form.Value["composition"][0])
	product.Description = form.Value["description"][0]

	// Get uploaded file
	if len(form.File["photo_img"]) > 0 {
		product.PhotoIMG = form.File["photo_img"][0]
	}

	userId := ctx.Locals("userId")
	if userId == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	_, err = h.MarketUsecase.CreateProduct(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create product",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Product created successfully",
	})
}
