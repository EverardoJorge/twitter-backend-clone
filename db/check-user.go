package db

import (
	"context"
	"time"

	"github.com/EverardoJorge/twitter-backend-clone/models"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisteredUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnect.Database("twitter-backend-clone")
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}