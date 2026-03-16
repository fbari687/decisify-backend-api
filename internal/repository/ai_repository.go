package repository

import (
	"bytes"
	"decisify-backend-api/internal/domain"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type AIRepository interface {
	Summarize(notes string, paragraphLengthMax string) (string, error)
	KeyPoints(notes string, keyPointsMax string) (*domain.KeyPointsResponse, error)
	GenerateQuiz(notes string, questionMax string) (*domain.QuizResponse, error)
}

type aiRepository struct {
	apiKey string
}

func NewAIRepository(apiKey string) AIRepository {
	return &aiRepository{
		apiKey: apiKey,
	}
}

func (o *aiRepository) callAI(prompt string) (string, error) {

	payload := map[string]interface{}{
		"model":       "maia/gemini-2.5-flash",
		"temperature": 0.3,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.maiarouter.ai/v1/chat/completions",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+o.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to call AI API")
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Choices[0].Message.Content, nil
}

func (o *aiRepository) Summarize(notes string, paragraphLengthMax string) (string, error) {

	prompt := "You are an assistant that summarizes notes clearly.\n\n" +
		"Summarize the following notes into a short and clear " + paragraphLengthMax + " paragraph.\n\n" +
		"Notes:\n\n" + notes

	return o.callAI(prompt)
}

func (o *aiRepository) KeyPoints(notes string, keyPointsMax string) (*domain.KeyPointsResponse, error) {

	prompt := `
Extract the key points from the following notes.

Return ONLY raw JSON in this format:

{
  "key_points": [
    "point",
    "point",
    "point"
  ]
}

Generate at most ` + keyPointsMax + ` key points.

Do not include markdown or explanations.

Notes:
` + notes

	result, err := o.callAI(prompt)
	if err != nil {
		return nil, err
	}

	result = strings.TrimSpace(result)
	result = strings.TrimPrefix(result, "```json")
	result = strings.TrimPrefix(result, "```")
	result = strings.TrimSuffix(result, "```")
	result = strings.TrimSpace(result)

	var keyPoints domain.KeyPointsResponse

	err = json.Unmarshal([]byte(result), &keyPoints)
	if err != nil {
		return nil, err
	}

	return &keyPoints, nil
}

func (o *aiRepository) GenerateQuiz(notes string, questionMax string) (*domain.QuizResponse, error) {

	prompt := `
You are an expert educator.

Generate ` + questionMax + ` multiple-choice quiz questions based on the notes.

Requirements:
- Each question must test understanding of the key ideas.
- Each question must have exactly 4 answer choices labeled A, B, C, and D.
- Only one answer must be correct.
- The incorrect answers (distractors) must be plausible and realistic but clearly incorrect.
- Avoid obvious wrong answers.

Return ONLY raw JSON in this format:

{
  "quiz": [
    {
      "question": "string",
      "choices": [
        { "id": "A", "text": "choice text" },
        { "id": "B", "text": "choice text" },
        { "id": "C", "text": "choice text" },
        { "id": "D", "text": "choice text" }
      ],
      "correct_answer": "A",
      "explanation": "short explanation why the answer is correct"
    }
  ]
}

Do not include markdown or extra text.

Notes:
` + notes

	result, err := o.callAI(prompt)
	if err != nil {
		return nil, err
	}

	var quiz domain.QuizResponse

	err = json.Unmarshal([]byte(result), &quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}
