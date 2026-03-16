package handler

import (
	"decisify-backend-api/internal/domain"
	"decisify-backend-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AIHandler interface {
	Summarize(c *fiber.Ctx) error
	KeyPoints(c *fiber.Ctx) error
	GenerateQuiz(c *fiber.Ctx) error
}

type aiHandler struct {
	service service.AIService
}

func NewAIHandler(service service.AIService) AIHandler {
	return &aiHandler{service: service}
}

func (h *aiHandler) Summarize(c *fiber.Ctx) error {

	var req domain.SummarizeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	summary, err := h.service.SummarizeNotes(req.Notes, req.Length)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(domain.SummarizeResponse{
		Summary: summary,
	})
}

func (h *aiHandler) KeyPoints(c *fiber.Ctx) error {

	var req domain.SummarizeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	points, err := h.service.GetKeyPoints(req.Notes, req.Length)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(points)
}

func (h *aiHandler) GenerateQuiz(c *fiber.Ctx) error {

	var req domain.SummarizeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	quiz, err := h.service.GenerateQuiz(req.Notes, req.Length)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(quiz)
}
