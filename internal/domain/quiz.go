package domain

type Choice struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type QuizItem struct {
	Question      string   `json:"question"`
	Choices       []Choice `json:"choices"`
	CorrectAnswer string   `json:"correct_answer"`
	Explanation   string   `json:"explanation"`
}

type QuizResponse struct {
	Quiz []QuizItem `json:"quiz"`
}
