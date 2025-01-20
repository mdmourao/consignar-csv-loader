package csv

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/mdmourao/consignar-csv-loader/models"
)

func LoadCsv(fileName string) []models.Entity {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var entities []models.Entity
	for _, record := range records {
		id, err := strconv.Atoi(strings.ReplaceAll(record[0], "\ufeff", ""))
		if err != nil {
			panic(err)
		}

		entity := models.Entity{
			IdentificationNumber: uint(id),
			OriginalName:         record[1],
			OriginalLocality:     record[2],
		}
		entities = append(entities, entity)
	}
	return entities
}
