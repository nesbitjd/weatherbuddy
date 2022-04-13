package weather

type Results struct {
	Coord  Coordinates `json:"coord"`
	Base   string      `json:"base"`
	Main   Main        `json:"main"`
	Wind   Wind        `json:"wind"`
	Clouds Clouds      `json:"clouds"`
	Sys    ForecastSys `json:"sys"`
	Name   string      `json:"name"`
}

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type ForecastSys struct {
	Population int `json:"population"`
}

type City struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int         `json:"population"`
	Sys        ForecastSys `json:"sys"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
}

type Temperature struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type Clouds struct {
	All int `json:"all"`
}
