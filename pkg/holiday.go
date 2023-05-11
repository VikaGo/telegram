package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var countries = []string{"USA", "UK", "France", "Germany", "Italy", "India"}

type HolidayResponse struct {
	Holidays []Holiday `json:"holidays"`
}

type Holiday struct {
	Name string `json:"name"`
}

func getHoliday(country string) (string, error) {
	apiKey := os.Getenv("apiKey")
	baseURL := "https://holidays.abstractapi.com/v1/"
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return "", nil
	}

	parameters := url.Values{}
	parameters.Add("api_key", apiKey)
	parameters.Add("country", country)
	parameters.Add("year", "now")
	parameters.Add("month", "now")
	parameters.Add("day", "now")

	u.RawQuery = parameters.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var data HolidayResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	if len(data.Holidays) == 0 {
		return "no holiday found", nil
	}

	holiday := data.Holidays[0].Name
	return holiday, nil
}
