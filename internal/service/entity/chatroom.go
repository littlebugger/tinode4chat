package entity

type ChatRoomID = string

type ChatRoom struct {
	ID           ChatRoomID `bson:"_id,omitempty"`
	Name         string     `bson:"name"`
	Participants []UserID   `bson:"participants,omitempty"`
}
