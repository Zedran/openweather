package ow

type Current struct {
	Coord      coord         `json:"coord"`
	Weather    []weather     `json:"weather"`
	Base       string        `json:"base"`
	Main       weatherMain   `json:"main"`
	Visibility int           `json:"visibility"`
	Wind       wind          `json:"wind"`
	Clouds     clouds        `json:"clouds"`
	Rain       precipitation `json:"rain"`
	Snow       precipitation `json:"snow"`

	// Time of data calculation
	DT         int           `json:"dt"`
	Sys        currentSys    `json:"sys"`

	// Shifts [sec] from UTC
	Timezone   int           `json:"timezone"`

	// City ID
	ID         int           `json:"id"`

	// City name
	Name       string        `json:"name"`

	// Internal parameter
	COD        int           `json:"cod"`
}

type currentSys struct {
	// Internal parameter
	Type    int    `json:"type"`

	// Internal parameter
	ID      int    `json:"id"`

	// Internal parameter
	Message string `json:"message"`

	// Country code
	Country string `json:"country"`

	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
