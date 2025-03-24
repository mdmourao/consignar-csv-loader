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
	"github.com/mdmourao/consignar-csv-loader/sqlite3"
)

func main() {
	startup()
	defer db.Disconnect()

	sqlite3.Connect()
	defer sqlite3.Disconnect()

	sqlite3.Migrate()

	f, err := os.OpenFile("./tmp/logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	entities := csv.LoadCsv("2024.csv")
	log.Printf("Loaded %d entities\n", len(entities))

	for i, entity := range entities {
		log.Printf("Processing entity %d/%d\n", i, len(entities))
		err, entityDB := db.GetEntity(entity)
		sqlite3.CreateEntity(models.EntityDb{
			IdentificationNumber: entityDB.IdentificationNumber,
			OriginalName:         entityDB.OriginalName,
			OriginalLocality:     entityDB.OriginalLocality,
			Duns:                 entityDB.Duns,
			Denomination:         entityDB.Denomination,
			Address:              entityDB.Address,
			PostalCode:           entityDB.PostalCode,
			Locality:             entityDB.Locality,
			Cae:                  entityDB.Cae,
			YearsOpen:            entityDB.YearsOpen,
			Website:              entityDB.Website,
			Latitude:             entityDB.Coordinates.Latitude,
			Longitude:            entityDB.Coordinates.Longitude,
		})
		if err != nil {
			panic(err)
		}

	}
}

func processEntity(entity models.Entity) {
	log.Println("processing entity:", entity.IdentificationNumber)

	// scrape NIF data
	nifRecord, err := nif.GetNifDataFromEInforma(entity.IdentificationNumber)
	if err != nil {
		log.Printf("ERROR: %d getting NIF data: %v\n", entity.IdentificationNumber, err)
		return
	}
	nifRecord.PopulateEntity(&entity)

	// get coordinates
	coordinates, err := google.GetCoordinatesFromAddress(fmt.Sprintf("%s %s %s", entity.Address, entity.PostalCode, entity.Locality))
	if err != nil {
		log.Printf("ERROR:  %d getting Coordinates data: %v\n", entity.IdentificationNumber, err)
		return
	}

	entity.Coordinates = coordinates

	// save entity
	// utils.PrettyJson(entity)
	err = db.CreateEntity(entity)
	if err != nil {
		log.Printf("ERROR:  %d creating db entity: %v\n", entity.IdentificationNumber, err)
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
