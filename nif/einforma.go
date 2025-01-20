package nif

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mdmourao/consignar-csv-loader/models"
)

const (
	REGEX_EXTRACT_NUMBER = "[0-9]+"
)

func GetNifDataFromEInforma(nif uint) (models.EInformaResponse, error) {

	url := fmt.Sprintf("https://www.einforma.pt/servlet/app/portal/ENTP/prod/ETIQUETA_EMPRESA_CONTRIBUINTE/nif/%d/contribuinte/%d", nif, nif)

	res, err := http.Get(url)
	if err != nil {
		return models.EInformaResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return models.EInformaResponse{}, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return models.EInformaResponse{}, err
	}

	var einforma = models.EInformaResponse{}

	doc.Find("tr").Each(func(i int, s1 *goquery.Selection) {

		s1.Find("td").Each(func(i2 int, s2 *goquery.Selection) {
			if strings.Contains(s1.Text(), "NIF:") && i2 == 1 {
				nif, _ := strconv.Atoi(s2.Text())
				einforma.Nif = uint(nif)
				return
			}

			if strings.Contains(s1.Text(), "DUNS:") && i2 == 1 {
				duns, _ := strconv.Atoi(s2.Text())
				einforma.Duns = uint(duns)
				return
			}

			if strings.Contains(s1.Text(), "Denominação:") && i2 == 1 {
				einforma.Denomination = s2.Text()
				return
			}

			if strings.Contains(s1.Text(), "Morada:") && i2 == 1 {
				einforma.Address = s2.Text()
				return
			}

			if strings.Contains(s1.Text(), "Código Postal:") && i2 == 1 {
				einforma.PostalCode = s2.Text()[0:8]
				einforma.Locality = s2.Text()[9:]
				return
			}

			if strings.Contains(s1.Text(), "Atividade (CAE):") && i2 == 1 {
				einforma.Cae = s2.Text()
				return
			}

			if strings.Contains(s1.Text(), "Antiguidade:") && i2 == 1 {
				re := regexp.MustCompile(REGEX_EXTRACT_NUMBER)
				yearStr := re.FindString(s2.Text())

				yearsOpen, _ := strconv.Atoi(yearStr)
				einforma.YearsOpen = uint(yearsOpen)
				return
			}

			if strings.Contains(s1.Text(), "Website:") && i2 == 1 {
				einforma.Website = s2.Text()
				return
			}

		})

	})

	if einforma.PostalCode == "" || einforma.Locality == "" {
		return models.EInformaResponse{}, fmt.Errorf("Postal Code and Locality not found")
	}

	return einforma, nil
}
