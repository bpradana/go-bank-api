package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type IPLocation struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func GetIPLocation(ip string) (location string, err error) {
	// Check if ip is local
	if strings.HasPrefix(ip, "192") {
		return "local", nil
	}

	// Get ip location
	resp, err := http.Get(fmt.Sprintf("http://ipinfo.io/%s", ip))
	if err != nil {
		log.Println("[balances] [usecase] error getting location, err: ", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[balances] [usecase] error reading location, err: ", err.Error())
		return "", err
	}

	// Parse response body
	var ipLocation IPLocation
	err = json.Unmarshal(body, &ipLocation)
	if err != nil {
		log.Println("[balances] [usecase] error unmarshalling location, err: ", err.Error())
		return "", err
	}

	// Return location
	location = fmt.Sprintf("%s, %s, %s", ipLocation.City, ipLocation.Region, ipLocation.Country)
	return location, nil
}
