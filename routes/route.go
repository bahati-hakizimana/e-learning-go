package routes

import (
	"github.com/bahati-hakizimana/e-learning-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register",controllers.Register)
}