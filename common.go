package ow

type clouds struct {
	// Cloudiness, %
	All int `json:"all"`
}

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type precipitation struct {
	// Precipitation value for the last 1 hour
	OneH   float64 `json:"1h"`

	// Precipitation value for the last 3 hours
	ThreeH float64 `json:"3h"`
}

type weather struct {
	// Weather condition ID
	ID          int    `json:"id"`

	// Group of weather parameters (Rain, Snow, Clouds etc.)
	Main        string `json:"main"`

	// Weather condition within the group
	Description string `json:"description"`

	// Weather icon ID
	Icon        string `json:"icon"`
}

type weatherMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`

	// Atmospheric pressure on the sea level
	SeaLevel  float64 `json:"sea_level"`

	// Atmospheric pressure on the ground level
	GrndLevel float64 `json:"grnd_level"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}
