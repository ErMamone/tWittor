package bd

import (
	"context"
	"log"
	"time"

	"github.com/ErMamone/tWittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckingIfUserExist recive an email by param and check if exist in DB*/
func CheckingIfUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCn.Database("twittor")

	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		log.Printf("Usuario Creado! -" + ID)
		return result, false, ID
	}

	log.Printf("Usuario encontrado -" + ID)
	return result, true, ID
}
