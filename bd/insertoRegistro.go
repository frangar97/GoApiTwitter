package bd

import (
	"context"
	"time"

	"github.com/frangar97/goapitwitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(usuario models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := col.InsertOne(ctx, usuario)

	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjId.String(), true, nil
}
