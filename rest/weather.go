package rest

import "encoding/xml"

// OpenWeatherRespXML represents the current weather response of the openweathermap API
type OpenWeatherRespXML struct {
	XMLName       xml.Name      `xml:"current"`
	Text          string        `xml:",chardata"`
	City          City          `xml:"city"`
	Temperature   Temperature   `xml:"temperature"`
	FeelsLike     FeelsLikeXML  `xml:"feels_like"`
	Humidity      Humidity      `xml:"humidity"`
	Pressure      Pressure      `xml:"pressure"`
	Wind          Wind          `xml:"wind"`
	Clouds        Clouds        `xml:"clouds"`
	Visibility    Visibility    `xml:"visibility"`
	Precipitation Precipitation `xml:"precipitation"`
	Weather       Weather       `xml:"weather"`
	Lastupdate    Lastupdate    `xml:"lastupdate"`
}
type Coord struct {
	Text string `xml:",chardata"`
	Lon  string `xml:"lon,attr"`
	Lat  string `xml:"lat,attr"`
}
type Sun struct {
	Text string `xml:",chardata"`
	Rise string `xml:"rise,attr"`
	Set  string `xml:"set,attr"`
}
type City struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Coord    Coord  `xml:"coord"`
	Country  string `xml:"country"`
	Timezone string `xml:"timezone"`
	Sun      Sun    `xml:"sun"`
}
type FeelsLikeXML struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Unit  string `xml:"unit,attr"`
}
type Temperature struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Min   string `xml:"min,attr"`
	Max   string `xml:"max,attr"`
	Unit  string `xml:"unit,attr"`
}
type Humidity struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Unit  string `xml:"unit,attr"`
}
type Pressure struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Unit  string `xml:"unit,attr"`
}
type Wind struct {
	Text      string    `xml:",chardata"`
	Speed     Speed     `xml:"speed"`
	Gusts     string    `xml:"gusts"`
	Direction Direction `xml:"direction"`
}
type Speed struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Unit  string `xml:"unit,attr"`
	Name  string `xml:"name,attr"`
}
type Direction struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Code  string `xml:"code,attr"`
	Name  string `xml:"name,attr"`
}
type Clouds struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
	Name  string `xml:"name,attr"`
}
type Visibility struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
}
type Precipitation struct {
	Text string `xml:",chardata"`
	Mode string `xml:"mode,attr"`
}
type Weather struct {
	Text   string `xml:",chardata"`
	Number string `xml:"number,attr"`
	Value  string `xml:"value,attr"`
	Icon   string `xml:"icon,attr"`
}
type Lastupdate struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
}
