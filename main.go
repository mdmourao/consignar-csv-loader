package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mdmourao/consignar-csv-loader/csv"
	"github.com/mdmourao/consignar-csv-loader/db"
	"github.com/mdmourao/consignar-csv-loader/google"
	"github.com/mdmourao/consignar-csv-loader/models"
	"github.com/mdmourao/consignar-csv-loader/nif"
	"github.com/mdmourao/consignar-csv-loader/utils"
)

func main() {
	startup()
	defer db.Disconnect()

	f, err := os.OpenFile("./tmp/logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	entities := csv.LoadCsv("2023.csv")
	log.Printf("Loaded %d entities\n", len(entities))

	for i, entity := range entities {
		log.Printf("Processing entity %d/%d\n", i+1, len(entities))
		processEntity(entity)
	}

}

func processEntity(entity models.Entity) {
	log.Println("processing entity:", entity.IdentificationNumber)

	// scrape NIF data
	nifRecord, err := nif.GetNifDataFromEInforma(entity.IdentificationNumber)
	if err != nil {
		log.Printf("ERROR: getting NIF data: %v\n", err)
		return
	}
	nifRecord.PopulateEntity(&entity)

	// get coordinates
	coordinates, err := google.GetCoordinatesFromAddress(fmt.Sprintf("%s %s %s", entity.Address, entity.PostalCode, entity.Locality))
	if err != nil {
		log.Printf("ERROR: getting Coordinates data: %v\n", err)
		return
	}

	entity.Coordinates = coordinates

	// save entity
	utils.PrettyJson(entity)
	err = db.CreateEntity(entity)
	if err != nil {
		log.Printf("ERROR: creating db entity: %v\n", err)
		return
	}

	log.Println("Entity OK")
}

func startup() {
	err := godotenv.Load(".dev/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Connect()
}
