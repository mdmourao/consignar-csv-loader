package google

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mdmourao/consignar-csv-loader/utils"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(utils.Dir(".dev/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestXxx(t *testing.T) {

	result, err := GetCoordinatesFromAddress("RUA ANTÓNIO MONTEIRO 3400-083 OLIVEIRA DO HOSPITAL")

	assert.Nil(t, err, err)

	assert.Equal(t, 40.3586291, result.Latitude)
	assert.Equal(t, -7.8560189, result.Longitude)
}
