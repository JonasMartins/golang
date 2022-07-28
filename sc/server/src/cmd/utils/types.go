package utils

import "time"

type ResultGetUsersChats struct {
	ChatId           string
	MessageId        string
	MessageBody      string
	MessageCreatedAt time.Time
	AuthorId         string
	AuthorName       string
	Seen             string
}

type ResultChatMembersByMemberId struct {
	ChatId     string
	MemberId   string
	MemberName string
}