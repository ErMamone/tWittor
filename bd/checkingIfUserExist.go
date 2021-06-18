package bd

import (
	"context"
	"time"

	"github.com/ErMamone/tWittor/models"
	"github.com/ErMamone/tWtittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckingIfUserExist recive an email by param and check if exist in DB*/
func CheckingIfUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCn.Database("tWittor")

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
