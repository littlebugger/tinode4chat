package repository

import (
	"context"
	"errors"
	"github.com/littlebugger/tinode4chat/internal/service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// TODO: maybe add user to his created channel?
// CreateChatRoom inserts a new chat room document into the MongoDB database.
func (repo *MongoRepository) CreateChatRoom(ctx context.Context, room entity.ChatRoom) (*string, error) {
	collection := repo.GetCollection("chatrooms")

	primID, err := collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	id := primID.InsertedID.(primitive.ObjectID).Hex()
	return &id, nil
}

// ListChatRooms retrieves a list of all chat rooms stored in the MongoDB database.
func (repo *MongoRepository) ListChatRooms(ctx context.Context) ([]entity.ChatRoom, error) {
	collection := repo.GetCollection("chatrooms")

	// TODO: this method need paging
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var chatRooms []entity.ChatRoom
	if err = cursor.All(ctx, &chatRooms); err != nil {
		return nil, err
	}

	return chatRooms, nil
}

// AddUserToChatRoom adds a user to a chat room's participant list
func (repo *MongoRepository) AddUserToChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error {
	// Start session for the transaction
	session, err := repo.Client.StartSession()
	if err != nil {
		log.Fatalf("failed to start mongodb transaction sesstion %w", err)
	}
	defer session.EndSession(ctx)

	chatRooms := repo.GetCollection("chatrooms")
	users := repo.GetCollection("users")

	rid, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return err
	}

	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Run transaction with a callback
	err = mongo.WithSession(context.TODO(), session, func(sc mongo.SessionContext) error {
		// Start a transaction
		if err := session.StartTransaction(); err != nil {
			return err
		}

		_, err = chatRooms.UpdateOne(
			ctx,
			bson.M{"_id": rid},
			bson.M{
				"$addToSet": bson.M{"participants": uid}, // Add the user to the participants array
			},
		)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		_, err = users.UpdateOne(
			ctx,
			bson.M{"_id": uid},
			bson.M{
				"$addToSet": bson.M{"joinedRooms": rid}, // Add room to user ChatRooms
			},
		)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Commit the transaction if both operations succeed
		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Transaction failed: %v", err)

		return err
	}

	return nil
}

// RemoveUserFromChatRoom removes a user from a chat room's participant list
func (repo *MongoRepository) RemoveUserFromChatRoom(ctx context.Context, roomID entity.ChatRoomID, userID entity.UserID) error {
	// Start session for the transaction
	session, err := repo.Client.StartSession()
	if err != nil {
		log.Fatalf("failed to start mongodb transaction sesstion %w", err)
	}
	defer session.EndSession(ctx)

	chatRooms := repo.GetCollection("chatrooms")
	users := repo.GetCollection("users")

	rid, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return err
	}

	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Run transaction with a callback
	err = mongo.WithSession(context.TODO(), session, func(sc mongo.SessionContext) error {
		// Start a transaction
		if err := session.StartTransaction(); err != nil {
			return err
		}

		_, err = chatRooms.UpdateOne(ctx, bson.M{"_id": rid}, bson.M{"$pull": bson.M{"participants": uid}})
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		_, err = users.UpdateOne(ctx, bson.M{"_id": uid}, bson.M{"$pull": bson.M{"joinedRooms": rid}})
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		// Commit the transaction if both operations succeed
		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Transaction failed: %v", err)

		return err
	}

	return nil
}

func (repo *MongoRepository) IsRoomExist(ctx context.Context, name string) (bool, error) {
	collection := repo.GetCollection("chatrooms")

	var chatRoom entity.ChatRoom
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&chatRoom)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, err
}
