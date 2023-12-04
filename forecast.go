package ow

type Forecast struct {
	// Internal parameter
	COD     string       `json:"cod"`

	// Internal parameter
	Message int          `json:"message"`
	
	// A number of forecast time steps returned in the API response
	Cnt     int          `json:"cnt"`

	List    []fcastTStep `json:"list"`

	City    city         `json:"city"`
}

type city struct {
	// City ID
	ID         int    `json:"id"`
	
	// City name
	Name       string `json:"name"`

	Coord      coord  `json:"coord"`

	// Country code
	Country    string `json:"country"`

	Population int    `json:"population"`

	// Shift [sec] from UTC
	Timezone   int    `json:"timezone"`

	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

type fcastMain struct {
	weatherMain

	// Internal forecast parameter
	TempKF      float64 `json:"temp_kf"`
}

type fcastTStep struct {
	// Time of data forecasted
	DT         int           `json:"dt"`

	Main       fcastMain     `json:"main"`
	Weather    []weather     `json:"weather"`
	Clouds     clouds        `json:"clouds"`
	Wind       wind          `json:"wind"`
	Visibility int           `json:"visibility"`
	
	// Probability of precipitation
	PoP        float64       `json:"pop"`

	Rain       precipitation `json:"rain"`
	Snow       precipitation `json:"snow"`

	Sys        fcastSys      `json:"sys"`
	DT_TXT     string        `json:"dt_txt"`
}

type fcastSys struct {
	// Part of the day
	PoD string `json:"pod"`
}
