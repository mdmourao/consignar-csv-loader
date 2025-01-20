package nif

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/mdmourao/consignar-csv-loader/models"
)

func GetNifData(nif uint) (models.NifRecord, error) {
	url := fmt.Sprintf("http://www.nif.pt/?json=1&q=%d&key=%s", nif, os.Getenv("NIFPT_API_KEY"))

	resp, err := http.Get(url)
	if err != nil {
		return models.NifRecord{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.NifRecord{}, err
	}

	var nifPt models.Response

	err = json.Unmarshal(body, &nifPt)
	if err != nil {
		return models.NifRecord{}, err
	}

	return nifPt.Records[strconv.FormatUint(uint64(nif), 10)], nil
}
