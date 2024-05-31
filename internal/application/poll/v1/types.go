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

type GetPollByIdResponseBodyQuestionChoiceItem struct {
	Position uint   `json:"position"`
	Text     string `json:"text"`
}

type GetPollByIdResponseBodyQuestionItem struct {
	AllowMultipleChoices bool                                         `json:"allowMultipleChoices"`
	Choices              []*GetPollByIdResponseBodyQuestionChoiceItem `json:"choices"`
	Position             uint                                         `json:"position"`
	Text                 string                                       `json:"text"`
}

type GetPollByIdResponseBody struct {
	CreatedAt time.Time                              `json:"createdAt"`
	EndsAt    time.Time                              `json:"endsAt"`
	ID        uuid.UUID                              `json:"id"`
	Questions []*GetPollByIdResponseBodyQuestionItem `json:"questions"`
	Title     string                                 `json:"title"`
}

type ListPollsResponseBodyQuestionChoiceItem struct {
	Position uint   `json:"position"`
	Text     string `json:"text"`
}

type ListPollsResponseBodyQuestionItem struct {
	AllowMultipleChoices bool                                       `json:"allowMultipleChoices"`
	Choices              []*ListPollsResponseBodyQuestionChoiceItem `json:"choices"`
	Position             uint                                       `json:"position"`
	Text                 string                                     `json:"text"`
}

type ListPollsResponseBody struct {
	CreatedAt time.Time                            `json:"createdAt"`
	EndsAt    time.Time                            `json:"endsAt"`
	ID        uuid.UUID                            `json:"id"`
	Questions []*ListPollsResponseBodyQuestionItem `json:"questions"`
	Title     string                               `json:"title"`
}
