package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	SettingsEndpointFormat = "https://www.pc-limit.com/settings?uuid=%s"
)

var settings Settings

type DayLimit struct {
	Limit int
}
type Settings struct {
	Uuid string
	Name string
	Status string
	DayLimits map[int]DayLimit
}

func ReadSettings() {

	uuid, err := ioutil.ReadFile("uuid.txt")
    if err != nil {
        log.Fatal(err)
    }

	resp, err := http.Get(fmt.Sprintf(SettingsEndpointFormat, uuid))
	if err != nil {
		log.Printf("Error reading config, %v.\n", err)
		return
	}
	
	body, err := ioutil.ReadAll(resp.Body)
  	if err != nil {
		log.Printf("Error reading config, %v.\n", err)
		return
	}

	err = json.Unmarshal(body, &settings)
	if err != nil {
		log.Printf("Error parsing config, %v.\n", err)
		return
	}
}