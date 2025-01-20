package google

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mdmourao/consignar-csv-loader/models"
	"googlemaps.github.io/maps"
)

const (
	LANGUAGE = "pt"
)

func GetCoordinatesFromAddress(address string) (models.Coordinates, error) {
	log.Println(os.Getenv("GOOGLE_MAPS_API_KEY"))
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.GeocodingRequest{
		Address:  address,
		Language: LANGUAGE,
	}
	results, err := c.Geocode(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	if len(results) == 0 {
		return models.Coordinates{}, fmt.Errorf("no results found for address: %s", address)
	}

	coordinates := models.Coordinates{
		Latitude:  results[0].Geometry.Location.Lat,
		Longitude: results[0].Geometry.Location.Lng,
	}

	return coordinates, nil
}
