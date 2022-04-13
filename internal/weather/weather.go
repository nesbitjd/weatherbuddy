package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config string

// GetWeather retrieves the weather data
func GetWeather(city string, api_key Config) (Results, error) {

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, api_key)

	resp, err := http.Get(url)
	if err != nil {
		// log.Fatalln(err)
		return Results{}, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		return Results{}, err
	}

	results := &Results{}

	err = json.Unmarshal(body, &results)
	if err != nil {
		//fmt.Println(err)
		return Results{}, err
	}

	return *results, nil
}
