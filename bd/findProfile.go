package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ErMamone/tWittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* findProfile busca un perfil en la base*/
func findProfile(ID string) (models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCn.Database("twittor")
	col := db.Collection("users")

	var perfil models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado "+ err.Error())
		return perfil, err
	}

	return perfil, nil

}