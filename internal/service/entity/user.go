package entity

type UserID = string

type User struct {
	ID          UserID      `bson:"_id,omitempty"`
	Email       string      `bson:"email"`
	Username    string      `bson:"username"`
	Password    string      `bson:"password"`
	JoinedRooms []MessageID `bson:"joinedRooms,omitempty"`
}

func (u User) Valid() bool {
	return u.Email != "" && u.Username != "" && u.Password != ""
}
