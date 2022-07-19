package resolvers

import (
	"context"
	"fmt"
	"src/graph/generated"

	"src/graph/model"
	"src/infra/orm/gorm/models"
	"src/main/auth"
)

type messageResolver struct{ *Resolver }

var _ generated.MessageResolver = (*messageResolver)(nil)

func (m *messageResolver) Seen(ctx context.Context, obj *models.Message) ([]string, error) {
	return obj.Seen, nil
}

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.CreateMessageInput) (*model.CreateAction, error) {

	errArr := []*model.Error{}
	result := model.CreateAction{
		Created: false,
		Errors:  errArr,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}

	_err := model.Error{
		Method:  "CreateMessage",
		Message: "",
		Field:   "",
		Code:    500,
	}

	message := models.Message{
		Body:     input.Body,
		AuthorId: input.AuthorID,
		ChatId:   input.ChatID,
	}

	if insert := r.DB.Create(&message); insert.Error != nil {
		_err.Message = insert.Error.Error()
		result.Errors = append(result.Errors, &_err)
	} else {
		result.Created = true
	}

	return &result, nil
}

func (q *queryResolver) GetMessagesByChat(ctx context.Context, chatId string) (*model.MessagesResponse, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateChat(ctx context.Context, input model.CreateChatInput) (*model.CreateAction, error) {

	errArr := []*model.Error{}
	result := model.CreateAction{
		Created: false,
		Errors:  errArr,
	}

	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}

	_err := model.Error{
		Method:  "CreateMessage",
		Message: "",
		Field:   "",
		Code:    500,
	}

	memebers := []*models.User{}
	messages := []*models.Message{}

	if foundMembers := r.DB.Find(&memebers, input.Members); foundMembers.Error != nil {
		_err.Message = foundMembers.Error.Error()
		_err.Field = "members"
		result.Errors = append(result.Errors, &_err)
	}

	if len(memebers) == 0 {
		_err.Message = "Could not find the members for the chat"
		_err.Code = 404
		_err.Field = "members"
		result.Errors = append(result.Errors, &_err)
	} else {

		chat := models.Chat{
			Members:  memebers,
			Messages: messages,
		}

		if createChat := r.DB.Create(&chat); createChat.Error != nil {
			_err.Message = createChat.Error.Error()
			result.Errors = append(result.Errors, &_err)
		} else {
			result.Created = true
		}
	}

	return &result, nil
}
