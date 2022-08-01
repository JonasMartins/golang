package data

import (
	"src/cmd/utils"
	"src/infra/orm/gorm/models"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// gets the raw
func GetUsersChatsFromRaw(chats []*utils.ResultGetUsersChats, chatMembers []*utils.ResultChatMembersByMemberId) ([]*models.Chat, error) {

	chatsLength := len(chats)
	chatsObj := []*models.Chat{}
	re := strings.NewReplacer("{", "", "}", "")
	var chat *models.Chat
	var currChatId *uuid.UUID
	var currChatIdString string = ""
	var res string
	var authorId *uuid.UUID
	var messageId *uuid.UUID
	var err error

	for i, c := range chats {

		var author = models.User{}
		var message = models.Message{}
		currChatId, err = ConvertUUidStringToUUidType(c.ChatId)
		if err != nil {
			return nil, err
		}

		if currChatId.String() != currChatIdString {
			// create a new chat
			newChat := models.Chat{}
			newChat.ID = *currChatId
			chat = &newChat
			currChatIdString = currChatId.String()
		}

		if i+1 < chatsLength {
			if chats[i+1].ChatId != currChatIdString {
				chatsObj = append(chatsObj, chat)
			}
		}

		authorId, err = ConvertUUidStringToUUidType(c.AuthorId)
		if err != nil {
			return nil, err
		}
		author.ID = *authorId
		author.Name = c.AuthorName
		messageId, err = ConvertUUidStringToUUidType(c.MessageId)
		if err != nil {
			return nil, err
		}
		message.ID = *messageId
		message.CreatedAt = c.MessageCreatedAt
		message.Body = c.MessageBody
		message.Author = &author
		res = re.Replace(c.Seen)
		message.Seen = strings.Split(res, ",")
		chat.Messages = append(chat.Messages, &message)

		if i+1 == chatsLength {
			chatsObj = append(chatsObj, chat)
		}

	}

	for _, m := range chatMembers {
		var member models.User
		id, _ := ConvertUUidStringToUUidType(m.MemberId)
		member.ID = *id
		member.Name = m.MemberName
		for _, c := range chatsObj {
			if m.ChatId == c.ID.String() {
				c.Members = append(c.Members, &member)
				break
			}
		}
	}

	return chatsObj, nil
}

func ConvertUUidStringToUUidType(s string) (*uuid.UUID, error) {

	u, err := uuid.FromString(s)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
