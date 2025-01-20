package models

type Entity struct {
	IdentificationNumber uint   `bson:"identification_number" json:"identification_number"`
	OriginalName         string `bson:"original_name" json:"original_name"`
	OriginalLocality     string `bson:"original_locality" json:"original_locality"`

	Duns         uint   `bson:"duns" json:"duns"`
	Denomination string `bson:"denomination" json:"denomination"`
	Address      string `bson:"address" json:"address"`
	PostalCode   string `bson:"postal_code" json:"postal_code"`
	Locality     string `bson:"locality" json:"locality"`
	Cae          string `bson:"cae" json:"cae"`
	YearsOpen    uint   `bson:"years_open" json:"years_open"`
	Website      string `bson:"website" json:"website"`

	Coordinates Coordinates `bson:"coordinates" json:"coordinates"`
}

type Coordinates struct {
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}
