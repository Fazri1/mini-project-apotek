package rajaongkir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mini-project-apotek/constants"
	"net/http"
	"strings"
)

type CityResponse struct {
	RajaOngkir struct {
		Query   []interface{} `json:"query"`
		Results []struct {
			CityID   string `json:"city_id"`
			CityName string `json:"city_name"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type CostResponse struct {
	RajaOngkir struct {
		Results []struct {
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value uint   `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func GetCityService() (map[string]string, error) {
	var data CityResponse
	url := "https://api.rajaongkir.com/starter/city"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return map[string]string{}, err
	}

	req.Header.Add("key", constants.RO_API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return map[string]string{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return map[string]string{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return map[string]string{}, err
	}

	results := make(map[string]string)
	for _, result := range data.RajaOngkir.Results {
		results[result.CityName] = result.CityID
	}

	defer res.Body.Close()
	return results, nil
}

func GetDeliveryCostService(idCity string) (int, interface{}) {
	var data CostResponse
	url := "https://api.rajaongkir.com/starter/cost"
	payloadString := fmt.Sprintf("origin=78&destination=%s&weight=5&courier=jne", idCity)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return 0, err.Error()
	}
	req.Header.Add("key", constants.RO_API_KEY)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err.Error()
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err.Error()
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err.Error()
	}

	responseRO := []int{}
	for _, result := range data.RajaOngkir.Results {
		for _, cost := range result.Costs {
			for _, option := range cost.Cost {
				responseRO = append(responseRO, int(option.Value))
			}
		}
	}

	if len(responseRO) > 0 {
		return responseRO[0], nil
	}
	return 0, "City not found"
}