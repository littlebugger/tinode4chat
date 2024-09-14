package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
