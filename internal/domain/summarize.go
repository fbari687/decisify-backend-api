package domain

type SummarizeRequest struct {
	Notes  string `json:"notes"`
	Length string `json:"length"`
}
type SummarizeResponse struct {
	Summary string `json:"summary"`
}
