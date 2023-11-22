package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//Arrange
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:3000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})

	t.Run("invalid amount", func(t *testing.T) {
		//Arrange
		amount := "abc"

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(nil, fiber.ErrBadRequest)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//http://localhost:3000/calculate?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("not found", func(t *testing.T) {
		//Arrange
		amount := 50

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(0, fiber.ErrNotFound)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)

		//Assert
		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	})
}
