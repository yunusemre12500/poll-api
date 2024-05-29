package v1

import (
	"time"

	"github.com/google/uuid"
)

type CreatePollRequestBodyQuestionChoiceItem struct {
	Position uint   `json:"position"`
	Text     string `json:"text"`
}

type CreatePollRequestBodyQuestionItem struct {
	AllowMultipleChoices bool                                       `json:"allowMultipleChoices"`
	Choices              []*CreatePollRequestBodyQuestionChoiceItem `json:"choices"`
	Position             uint                                       `json:"position"`
	Text                 string                                     `json:"text"`
}

type CreatePollRequestBody struct {
	EndsAt    time.Time                            `json:"endsAt"`
	Questions []*CreatePollRequestBodyQuestionItem `json:"questions"`
	Title     string                               `json:"title"`
}

type CreatePollResponseBodyQuestionChoiceItem struct {
	Position uint   `json:"position"`
	Text     string `json:"text"`
}

type CreatePollResponseBodyQuestionItem struct {
	AllowMultipleChoices bool                                        `json:"allowMultipleChoices"`
	Choices              []*CreatePollResponseBodyQuestionChoiceItem `json:"choices"`
	Position             uint                                        `json:"position"`
	Text                 string                                      `json:"text"`
}

type CreatePollResponseBody struct {
	CreatedAt time.Time                             `json:"createdAt"`
	EndsAt    time.Time                             `json:"endsAt"`
	ID        uuid.UUID                             `json:"id"`
	Questions []*CreatePollResponseBodyQuestionItem `json:"questions"`
	Title     string                                `json:"title"`
}
