package v1

import (
	"time"

	"github.com/google/uuid"
	v1 "github.com/yunusemre12500/poll-api/internal/application/poll/v1"
)

type ChoiceItem struct {
	Position uint   `bson:"position"`
	Text     string `bson:"text"`
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

func (choiceItem *ChoiceItem) IntoGetPollByIdResponseBodyQuestionChoiceItem() *v1.GetPollByIdResponseBodyQuestionChoiceItem {
	return &v1.GetPollByIdResponseBodyQuestionChoiceItem{
		Position: choiceItem.Position,
		Text:     choiceItem.Text,
	}
}

func (choiceItem *ChoiceItem) IntoListPollsResponseBodyQuestionChoiceItem() *v1.ListPollsResponseBodyQuestionChoiceItem {
	return &v1.ListPollsResponseBodyQuestionChoiceItem{
		Position: choiceItem.Position,
		Text:     choiceItem.Text,
	}
}

type QuestionItem struct {
	AllowMultipleChoices bool          `bson:"allowMultipleChoices"`
	Choices              []*ChoiceItem `bson:"choices"`
	Position             uint          `bson:"position"`
	Text                 string        `bson:"text"`
}

func NewQuestionItemFromCreatePollRequestBodyQuestionItem(item *v1.CreatePollRequestBodyQuestionItem) *QuestionItem {
	var choiceItems []*ChoiceItem

	for _, createPollRequestQuestionChoiceItem := range item.Choices {
		choiceItems = append(choiceItems, NewChoiceItemFromCreatePollRequestBodyQuestionChoiceItem(createPollRequestQuestionChoiceItem))
	}

	return &QuestionItem{
		AllowMultipleChoices: item.AllowMultipleChoices,
		Choices:              choiceItems,
		Position:             item.Position,
		Text:                 item.Text,
	}
}

func (questionItem *QuestionItem) IntoCretePollResponseBodyQuestionItem() *v1.CreatePollResponseBodyQuestionItem {
	var createPollResponseBodyQuestionChoiceItems []*v1.CreatePollResponseBodyQuestionChoiceItem

	for _, choiceItem := range questionItem.Choices {
		createPollResponseBodyQuestionChoiceItems = append(createPollResponseBodyQuestionChoiceItems, choiceItem.IntoCreatePollResponseBodyQuestionChoiceItem())
	}

	return &v1.CreatePollResponseBodyQuestionItem{
		AllowMultipleChoices: questionItem.AllowMultipleChoices,
		Choices:              createPollResponseBodyQuestionChoiceItems,
		Position:             questionItem.Position,
		Text:                 questionItem.Text,
	}
}

func (questionItem *QuestionItem) IntoGetPollByIdResponseBodyQuestionItem() *v1.GetPollByIdResponseBodyQuestionItem {
	var getPollByIdResponseBodyChoiceItems []*v1.GetPollByIdResponseBodyQuestionChoiceItem

	for _, choiceItem := range questionItem.Choices {
		getPollByIdResponseBodyChoiceItems = append(getPollByIdResponseBodyChoiceItems, choiceItem.IntoGetPollByIdResponseBodyQuestionChoiceItem())
	}

	return &v1.GetPollByIdResponseBodyQuestionItem{
		AllowMultipleChoices: questionItem.AllowMultipleChoices,
		Choices:              getPollByIdResponseBodyChoiceItems,
		Position:             questionItem.Position,
		Text:                 questionItem.Text,
	}
}

func (questionItem *QuestionItem) IntoListPollsResponseBodyQuestionItem() *v1.ListPollsResponseBodyQuestionItem {
	var listPollsResponseBodyQuestionChoiceItems []*v1.ListPollsResponseBodyQuestionChoiceItem

	for _, choiceItem := range questionItem.Choices {
		listPollsResponseBodyQuestionChoiceItems = append(listPollsResponseBodyQuestionChoiceItems, choiceItem.IntoListPollsResponseBodyQuestionChoiceItem())
	}

	return &v1.ListPollsResponseBodyQuestionItem{
		AllowMultipleChoices: questionItem.AllowMultipleChoices,
		Choices:              listPollsResponseBodyQuestionChoiceItems,
		Position:             questionItem.Position,
		Text:                 questionItem.Text,
	}
}

type Poll struct {
	CreatedAt time.Time       `bson:"createdAt"`
	EndsAt    time.Time       `bson:"endsAt"`
	ID        uuid.UUID       `bson:"_id"`
	Questions []*QuestionItem `bson:"questions"`
	Title     string          `bson:"title"`
}

func NewPollFromCreatePollRequestBody(body *v1.CreatePollRequestBody) *Poll {
	var questionItems []*QuestionItem

	for _, createPollRequestBodyQuestion := range body.Questions {
		questionItems = append(questionItems, NewQuestionItemFromCreatePollRequestBodyQuestionItem(createPollRequestBodyQuestion))
	}

	return &Poll{
		CreatedAt: time.Now().UTC(),
		EndsAt:    body.EndsAt.UTC(),
		ID:        uuid.New(),
		Questions: questionItems,
		Title:     body.Title,
	}
}

func (poll *Poll) IntoCreatePollResponseBody() *v1.CreatePollResponseBody {
	var createPollResponseBodyQuestionItems []*v1.CreatePollResponseBodyQuestionItem

	for _, question := range poll.Questions {
		createPollResponseBodyQuestionItems = append(createPollResponseBodyQuestionItems, question.IntoCretePollResponseBodyQuestionItem())
	}

	return &v1.CreatePollResponseBody{
		CreatedAt: poll.CreatedAt,
		EndsAt:    poll.EndsAt,
		ID:        poll.ID,
		Questions: createPollResponseBodyQuestionItems,
		Title:     poll.Title,
	}
}

func (poll *Poll) IntoGetPollByIdResponseBody() *v1.GetPollByIdResponseBody {
	var getPollByIdResponseBodyQuestionItems []*v1.GetPollByIdResponseBodyQuestionItem

	for _, questionItem := range poll.Questions {
		getPollByIdResponseBodyQuestionItems = append(getPollByIdResponseBodyQuestionItems, questionItem.IntoGetPollByIdResponseBodyQuestionItem())
	}

	return &v1.GetPollByIdResponseBody{
		CreatedAt: poll.CreatedAt,
		EndsAt:    poll.EndsAt,
		ID:        poll.ID,
		Questions: getPollByIdResponseBodyQuestionItems,
		Title:     poll.Title,
	}
}

func (poll *Poll) IntoListPollsResponseBody() *v1.ListPollsResponseBody {
	var listPollsResponseBodyQuestionItems []*v1.ListPollsResponseBodyQuestionItem

	for _, questionItem := range poll.Questions {
		listPollsResponseBodyQuestionItems = append(listPollsResponseBodyQuestionItems, questionItem.IntoListPollsResponseBodyQuestionItem())
	}

	return &v1.ListPollsResponseBody{
		CreatedAt: poll.CreatedAt,
		EndsAt:    poll.EndsAt,
		ID:        poll.ID,
		Questions: listPollsResponseBodyQuestionItems,
		Title:     poll.Title,
	}
}
