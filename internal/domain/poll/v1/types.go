package v1

import (
	"time"

	"github.com/google/uuid"
	v1 "github.com/yunusemre12500/poll-api/internal/application/poll/v1"
)

type ChoiceItem struct {
	Position uint   `json:"position"`
	Text     string `json:"text"`
}

func NewChoiceItemFromCreatePollRequestBodyQuestionChoiceItem(item *v1.CreatePollRequestBodyQuestionChoiceItem) *ChoiceItem {
	return &ChoiceItem{
		Position: item.Position,
		Text:     item.Text,
	}
}

func (choiceItem *ChoiceItem) IntoCreatePollResponseBodyQuestionChoiceItem() *v1.CreatePollResponseBodyQuestionChoiceItem {
	return &v1.CreatePollResponseBodyQuestionChoiceItem{
		Position: choiceItem.Position,
		Text:     choiceItem.Text,
	}
}

type QuestionItem struct {
	AllowMultipleChoices bool          `json:"allowMultipleChoices"`
	Choices              []*ChoiceItem `json:"choices"`
	Position             uint          `json:"position"`
	Text                 string        `json:"text"`
}

func NewQuestionItemFromCreatePollRequestBodyQuestionItem(item *v1.CreatePollRequestBodyQuestionItem) *QuestionItem {
	var choiceItems []*ChoiceItem

	for _, choice := range item.Choices {
		choiceItems = append(choiceItems, NewChoiceItemFromCreatePollRequestBodyQuestionChoiceItem(choice))
	}

	return &QuestionItem{
		AllowMultipleChoices: item.AllowMultipleChoices,
		Choices:              choiceItems,
		Position:             item.Position,
		Text:                 item.Text,
	}
}

func (questionItem *QuestionItem) IntoCretePollResponseQuestionItem() *v1.CreatePollResponseBodyQuestionItem {
	var createPollResponseBodyQuestionChoiceItems []*v1.CreatePollResponseBodyQuestionChoiceItem

	for _, choice := range questionItem.Choices {
		createPollResponseBodyQuestionChoiceItems = append(createPollResponseBodyQuestionChoiceItems, choice.IntoCreatePollResponseBodyQuestionChoiceItem())
	}

	return &v1.CreatePollResponseBodyQuestionItem{
		AllowMultipleChoices: questionItem.AllowMultipleChoices,
		Choices:              createPollResponseBodyQuestionChoiceItems,
		Position:             questionItem.Position,
		Text:                 questionItem.Text,
	}
}

type Poll struct {
	CreatedAt time.Time       `json:"createdAt"`
	EndsAt    time.Time       `json:"endsAt"`
	ID        uuid.UUID       `json:"id"`
	Questions []*QuestionItem `json:"questions"`
	Title     string          `json:"title"`
}

func NewPollFromCreatePollRequestBody(body *v1.CreatePollRequestBody) *Poll {
	var questions []*QuestionItem

	for _, question := range body.Questions {
		questions = append(questions, NewQuestionItemFromCreatePollRequestBodyQuestionItem(question))
	}

	return &Poll{
		CreatedAt: time.Now().UTC(),
		EndsAt:    body.EndsAt.UTC(),
		ID:        uuid.New(),
		Questions: questions,
		Title:     body.Title,
	}
}

func (poll *Poll) IntoCreatePollResponseBody() *v1.CreatePollResponseBody {
	var createPollResponseBodyQuestionItems []*v1.CreatePollResponseBodyQuestionItem

	for _, question := range poll.Questions {
		createPollResponseBodyQuestionItems = append(createPollResponseBodyQuestionItems, question.IntoCretePollResponseQuestionItem())
	}

	return &v1.CreatePollResponseBody{
		CreatedAt: poll.CreatedAt,
		EndsAt:    poll.EndsAt,
		ID:        poll.ID,
		Questions: createPollResponseBodyQuestionItems,
		Title:     poll.Title,
	}
}
