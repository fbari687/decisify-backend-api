package routes

import (
	"decisify-backend-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, aiHandler handler.AIHandler) {

	api := app.Group("/api")

	ai := api.Group("/ai")

	ai.Post("/summary", aiHandler.Summarize)
	ai.Post("/keypoints", aiHandler.KeyPoints)
	ai.Post("/quiz", aiHandler.GenerateQuiz)
}
