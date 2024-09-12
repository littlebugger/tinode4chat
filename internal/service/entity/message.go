package entity

import (
	"time"
)

type MessageID = string

type Message struct {
	ID         MessageID  `bson:"_id,omitempty"`
	Content    string     `bson:"content"`
	Author     UserID     `bson:"author"`
	ChatRoomID ChatRoomID `bson:"chatRoomID"`
	Timestamp  time.Time  `bson:"timestamp"`
}
