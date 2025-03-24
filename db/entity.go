package db

import (
	"context"

	"github.com/mdmourao/consignar-csv-loader/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateEntity(entity models.Entity) error {
	_, err := client.Database("consignar").Collection("entities2024").InsertOne(context.TODO(), entity)
	return err
}

func GetEntity(entity models.Entity) (error, models.Entity) {
	filter := bson.D{{"identification_number", entity.IdentificationNumber}}

	dbEntity := models.Entity{}
	err := client.Database("consignar").Collection("entities2024").FindOne(context.TODO(), filter).Decode(&dbEntity)

	return err, dbEntity
}
