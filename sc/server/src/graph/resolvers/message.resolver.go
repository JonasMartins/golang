package resolvers

import (
	"context"
	"fmt"
	"src/graph/generated"
	"time"

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
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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

	chat := models.Chat{}

	message := models.Message{
		Body:     input.Body,
		AuthorId: input.AuthorID,
		ChatId:   input.ChatID,
		Seen:     []string{input.AuthorID},
	}

	if insert := r.DB.Create(&message); insert.Error != nil {
		_err.Message = insert.Error.Error()
		result.Errors = append(result.Errors, &_err)
		tx.Rollback()
	} else {
		if updateChat := r.DB.Model(&chat).Where("id = ?", input.ChatID).Update("updated_at", time.Now()); updateChat.Error != nil {
			_err.Message = updateChat.Error.Error()
			result.Errors = append(result.Errors, &_err)
			tx.Rollback()
		} else {
			tx.Commit()
			result.Created = true
		}
	}

	return &result, nil
}

func (q *queryResolver) GetMessagesByChat(ctx context.Context, chatId string, limit *int, offset *int) (*model.MessagesResponse, error) {
	errArr := []*model.Error{}
	messages := []*models.Message{}
	result := model.MessagesResponse{
		Messages: nil,
		Errors:   errArr,
	}
	if userId := auth.ForUserIdContext(ctx); len(userId) == 0 {
		return &result, fmt.Errorf("access denied")
	}
	_err := model.Error{
		Method:  "GetMessagesByChat",
		Message: "",
		Field:   "id",
		Code:    500,
	}

	if foundMessages := q.DB.Where("chat_id = ?", chatId).Preload("Author").Offset(int(*offset)).Limit(int(*limit)).Find(&messages); foundMessages.Error != nil {
		_err.Message = foundMessages.Error.Error()
		_err.Field = "chatId"
		result.Errors = append(result.Errors, &_err)
	} else {
		result.Messages = messages
	}

	return &result, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, input model.CreateChatInput) (*model.CreateAction, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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

	members := []*models.User{}
	messages := []*models.Message{}

	if foundMembers := r.DB.Find(&members, input.Members); foundMembers.Error != nil {
		_err.Message = foundMembers.Error.Error()
		_err.Field = "members"
		result.Errors = append(result.Errors, &_err)
	} else {

		if len(members) == 0 {
			_err.Message = "Could not find the members for the chat"
			_err.Code = 404
			_err.Field = "members"
			result.Errors = append(result.Errors, &_err)
		} else {

			chat := models.Chat{
				Members:  members,
				Messages: messages,
			}

			// avoid upsert with:
			// 1 -> Omit("Members.*")
			// 2 -> DB.Model(&chat).Association("Members").Append(&members)

			if createChat := r.DB.Omit("Members.*").Create(&chat); createChat.Error != nil {
				_err.Message = createChat.Error.Error()
				result.Errors = append(result.Errors, &_err)
				tx.Rollback()
			} else {
				tx.Commit()
				result.Created = true
			}
		}
	}

	return &result, nil
}
