package rest

import "encoding/xml"

type OpenWeatherResp struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

func (c OpenWeatherResp) GetCloudsDescription() string {
	for _, w := range c.Weather {
		if w.Main == "Clouds" {
			return w.Description
		}
	}
	return ""
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}
type Clouds struct {
	All int `json:"all"`
}
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}




type OpenWeatherRespXML struct {
	XMLName xml.Name `xml:"current"`
	Text    string   `xml:",chardata"`
	City    struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Name  string `xml:"name,attr"`
		Coord struct {
			Text string `xml:",chardata"`
			Lon  string `xml:"lon,attr"`
			Lat  string `xml:"lat,attr"`
		} `xml:"coord"`
		Country  string `xml:"country"`
		Timezone string `xml:"timezone"`
		Sun      struct {
			Text string `xml:",chardata"`
			Rise string `xml:"rise,attr"`
			Set  string `xml:"set,attr"`
		} `xml:"sun"`
	} `xml:"city"`
	Temperature struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Min   string `xml:"min,attr"`
		Max   string `xml:"max,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"temperature"`
	FeelsLike struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"feels_like"`
	Humidity struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"humidity"`
	Pressure struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"pressure"`
	Wind struct {
		Text  string `xml:",chardata"`
		Speed struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
			Unit  string `xml:"unit,attr"`
			Name  string `xml:"name,attr"`
		} `xml:"speed"`
		Gusts     string `xml:"gusts"`
		Direction struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
			Code  string `xml:"code,attr"`
			Name  string `xml:"name,attr"`
		} `xml:"direction"`
	} `xml:"wind"`
	Clouds struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
		Name  string `xml:"name,attr"`
	} `xml:"clouds"`
	Visibility struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"visibility"`
	Precipitation struct {
		Text string `xml:",chardata"`
		Mode string `xml:"mode,attr"`
	} `xml:"precipitation"`
	Weather struct {
		Text   string `xml:",chardata"`
		Number string `xml:"number,attr"`
		Value  string `xml:"value,attr"`
		Icon   string `xml:"icon,attr"`
	} `xml:"weather"`
	Lastupdate struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"lastupdate"`
}

