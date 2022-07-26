package utils

import "time"

type ResultGetUsersChats struct {
	chat_id            string
	message_body       string
	message_created_at time.Time
	author_id          string
	author_name        string
	seen               []string
}
