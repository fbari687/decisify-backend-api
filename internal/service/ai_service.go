package service

import (
	"decisify-backend-api/internal/domain"
	"decisify-backend-api/internal/repository"
)

type AIService interface {
	SummarizeNotes(notes string, paragraphLengthMax string) (string, error)
	GetKeyPoints(notes string, keyPointsMax string) (*domain.KeyPointsResponse, error)
	GenerateQuiz(notes string, questionMax string) (*domain.QuizResponse, error)
}

type aiService struct {
	repo repository.AIRepository
}

func NewAIService(repo repository.AIRepository) AIService {
	return &aiService{repo: repo}
}

func (s *aiService) SummarizeNotes(notes string, paragraphLengthMax string) (string, error) {
	return s.repo.Summarize(notes, paragraphLengthMax)
}

func (s *aiService) GetKeyPoints(notes string, keyPointsMax string) (*domain.KeyPointsResponse, error) {
	return s.repo.KeyPoints(notes, keyPointsMax)
}

func (s *aiService) GenerateQuiz(notes string, questionMax string) (*domain.QuizResponse, error) {
	return s.repo.GenerateQuiz(notes, questionMax)
}
