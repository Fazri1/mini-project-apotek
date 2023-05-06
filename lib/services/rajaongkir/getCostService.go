package rajaongkir

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mini-project-apotek/constants"
	"net/http"
	"strings"
)

type CostResponse struct {
	RajaOngkir struct {
		// Query   []interface{} `json:"query"`
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
			// ProvinceID string `json:"province_id"`
		} `json:"results"`
		// Status struct {
		// 	Code        uint   `json:"code"`
		// 	Description string `json:"description"`
		// } `json:"status"`
	} `json:"rajaongkir"`
}

func GetDeliveryCostService(idCity string) (int, interface{}) {
	url := "https://api.rajaongkir.com/starter/cost"
	ps := fmt.Sprintf("origin=78&destination=%s&weight=5&courier=jne", idCity)

	payload := strings.NewReader(ps)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", constants.RO_API_KEY)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var data CostResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	lis := []int{}
	// results := make(map[string][]interface{})
	for _, result := range data.RajaOngkir.Results {
		for _, cost := range result.Costs {
			for _, option := range cost.Cost {
				lis = append(lis, int(option.Value))
				// results[cost.Service] = []interface{}{option.Value, option.Etd}
			}
		}
	}

	if len(lis) > 0 {
		return lis[0], nil
	}
	return 0, "City not found"
	// return results
}
