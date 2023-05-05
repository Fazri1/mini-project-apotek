package services

import (
	"encoding/json"
	"io/ioutil"
	"mini-project-apotek/constants"
	"net/http"
)

type ResponseData struct {
	RajaOngkir struct {
		Query   []interface{} `json:"query"`
		Results []struct {
			Province   string `json:"province"`
			ProvinceID string `json:"province_id"`
		} `json:"results"`
		Status struct {
			Code        uint   `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
	} `json:"rajaongkir"`
}

func GetProvinceService() map[string]string {
	url := "https://api.rajaongkir.com/starter/province"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("key", constants.API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	results := make(map[string]string)
	for _, result := range data.RajaOngkir.Results {
		results[result.ProvinceID] = result.Province
	}

	defer res.Body.Close()
	return results
}
