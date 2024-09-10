package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Content    string             `bson:"content"`
	Author     primitive.ObjectID `bson:"author"`
	ChatRoomID primitive.ObjectID `bson:"chatRoomID"`
	Timestamp  primitive.DateTime `bson:"timestamp"`
}
