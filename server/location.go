package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// Location TODO: Comment
type Location struct {
	Lat      float64 `json:"latitude"`
	Lon      float64 `json:"longitude"`
	CityName string  `json:"city"`
}

// TODO: Write file with known ip -> Location mappings

func getLocation(ip net.IP) Location {
	var loc Location
	//ip = net.ParseIP("93.220.207.105")

	err := json.Unmarshal(establishConnection("http://freegeoip.net/json/"+ip.String()), &loc)
	if err != nil {
		loc.Lat = 49.8695
		loc.Lon = 8.6507
		loc.CityName = "Darmstadt"
	}

	return loc
}

func establishConnection(url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode != 200 {

	}

	return body
}
