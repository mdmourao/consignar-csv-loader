package db

import (
	"context"

	"github.com/mdmourao/consignar-csv-loader/models"
)

func CreateEntity(entity models.Entity) error {
	_, err := client.Database("consignar").Collection("entities").InsertOne(context.TODO(), entity)
	return err
}
