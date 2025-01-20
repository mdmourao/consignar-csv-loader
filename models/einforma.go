package models

type EInformaResponse struct {
	Nif          uint   `bson:"nif" json:"nif"`
	Duns         uint   `bson:"duns" json:"duns"`
	Denomination string `bson:"denomination" json:"denomination"`
	Address      string `bson:"address" json:"address"`
	PostalCode   string `bson:"postal_code" json:"postal_code"`
	Locality     string `bson:"locality" json:"locality"`
	Cae          string `bson:"cae" json:"cae"`
	YearsOpen    uint   `bson:"years_open" json:"years_open"`
	Website      string `bson:"website" json:"website"`
}

func (e EInformaResponse) PopulateEntity(entity *Entity) {

	entity.Duns = e.Duns
	entity.Denomination = e.Denomination
	entity.Address = e.Address
	entity.PostalCode = e.PostalCode
	entity.Locality = e.Locality
	entity.Cae = e.Cae
	entity.YearsOpen = e.YearsOpen
	entity.Website = e.Website

}
