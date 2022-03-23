package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"main.go/types"
)

// GetWeather retrieves the weather data
func GetWeather(city string) (types.Results, error) {
	api_key := os.Getenv("API_KEY")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key)

	resp, err := http.Get(url)
	if err != nil {
		// log.Fatalln(err)
		return types.Results{}, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		return types.Results{}, err
	}

	fmt.Printf("Here is the body: %+v\n", string(body))

	results := &types.Results{}

	err = json.Unmarshal(body, &results)
	if err != nil {
		//fmt.Println(err)
		return types.Results{}, err
	}

	return *results, nil
}
