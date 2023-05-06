package rajaongkir

import (
	"encoding/json"
	"io/ioutil"
	"mini-project-apotek/constants"
	"net/http"
)

type Response struct {
	RajaOngkir struct {
		Query   []interface{} `json:"query"`
		Results []struct {
			CityID   string `json:"city_id"`
			CityName string `json:"city_name"`
			// Type       string `json:"type"`
			// PostalCode string `json:"postal_code"`
			// ProvinceID string `json:"province_id"`
			// Province   string `json:"province"`
		} `json:"results"`
		// Status struct {
		// 	Code        uint   `json:"code"`
		// 	Description string `json:"description"`
		// } `json:"status"`
	} `json:"rajaongkir"`
}

func GetCityService() map[string]string {
	url := "https://api.rajaongkir.com/starter/city"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", constants.RO_API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	results := make(map[string]string)
	for _, result := range data.RajaOngkir.Results {
		results[result.CityName] = result.CityID
	}

	defer res.Body.Close()
	return results
}
