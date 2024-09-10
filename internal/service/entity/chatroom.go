package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatRoom struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	Participants []primitive.ObjectID `bson:"participants,omitempty"`
}
