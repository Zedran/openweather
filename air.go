package ow

type AirPollution struct {
	Coord coord      `json:"coord"`
	List  []airTStep `json:"list"`
}

type airMain struct {
	// Air Quality Index
	AQI int `json:"aqi"`
}

type airTStep struct {
	// Date and time
	DT         int        `json:"dt"`

	Main       airMain    `json:"main"`
	Components components `json:"components"`
}

type components struct {
	CO    float64 `json:"co"`
	NO    float64 `json:"no"`
	NO2   float64 `json:"no2"`
	O3    float64 `json:"o3"`
	SO2   float64 `json:"so2"`
	PM2_5 float64 `json:"pm2_5"`
	PM10  float64 `json:"pm10"`
	NH3   float64 `json:"nh3"`
}
