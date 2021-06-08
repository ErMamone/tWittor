package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Usuario struct {
	ID				primitive.ObjectID		`bson:"_id, omitempy" json:"id"`
	Nombre 			string					`bson:"nombre" json:"nombre,omitempy"`
	Apellido 		string					`bson:"apellido" json:"apellido,omitempy"`
	FechaNacimiento time.Time				`bson:"fechaNacimiento" json:"fechaNacimiento,omitempy"`
	Email 			string 					`bson:"email" json:"email"`
	Password		string 					`bson:"password" json:"password,omitempty"`
	Avatar 			string					`bson:"avatar" json:"avatar,omitempty"`
	Banner 			string 					`bson:"banner" json:"banner,omitempty"`
	Biografia 		string 					`bson:"biografia" json:"biografia,omitempty"`
	Ubicaion 		string					`bson:"ubicaion" json:"ubicaion,omitempty"`
	SitioWeb 		string 					`bson:"sitioWeb" json:"sitioWeb,omitempty"`
}