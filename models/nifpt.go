package models

type Geo struct {
	Region string `bson:"region" json:"region"`
	County string `bson:"county" json:"county"`
	Parish string `bson:"parish" json:"parish"`
}

type Place struct {
	Address string `bson:"address" json:"address"`
	Pc4     string `bson:"pc4" json:"pc4"`
	Pc3     string `bson:"pc3" json:"pc3"`
	City    string `bson:"city" json:"city"`
}

type Structure struct {
	Nature          string `bson:"nature" json:"nature"`
	Capital         string `bson:"capital" json:"capital"`
	CapitalCurrency string `bson:"capital_currency" json:"capital_currency"`
}

type Contacts struct {
	Email   string `bson:"email" json:"email"`
	Phone   string `bson:"phone" json:"phone"`
	Website string `bson:"website" json:"website"`
	Fax     string `bson:"fax" json:"fax"`
}

type NifRecord struct {
	Nif        int       `bson:"nif" json:"nif"`
	SeoURL     string    `bson:"seo_url" json:"seo_url"`
	Title      string    `bson:"title" json:"title"`
	Address    string    `bson:"address" json:"address"`
	Pc4        string    `bson:"pc4" json:"pc4"`
	Pc3        string    `bson:"pc3" json:"pc3"`
	City       string    `bson:"city" json:"city"`
	Activity   string    `bson:"activity" json:"activity"`
	Status     string    `bson:"status" json:"status"`
	Cae        string    `bson:"cae" json:"cae"`
	Contacts   Contacts  `bson:"contacts" json:"contacts"`
	Structure  Structure `bson:"structure" json:"structure"`
	Geo        Geo       `bson:"geo" json:"geo"`
	Place      Place     `bson:"place" json:"place"`
	Racius     string    `bson:"racius" json:"racius"`
	Alias      string    `bson:"alias" json:"alias"`
	Portugalio string    `bson:"portugalio" json:"portugalio"`
}

type Credits struct {
	Used string   `bson:"used" json:"used"`
	Left []string `bson:"left" json:"left"`
}

type Response struct {
	Result        string               `bson:"result" json:"result"`
	Records       map[string]NifRecord `bson:"records" json:"records"`
	NifValidation bool                 `bson:"nif_validation" json:"nif_validation"`
	IsNif         bool                 `bson:"is_nif" json:"is_nif"`
	Credits       Credits              `bson:"credits" json:"credits"`
}
