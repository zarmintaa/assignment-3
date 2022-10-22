package helpers

import (
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

func GetDataJson() models.StatusResponse {
	jsonFile, err := os.Open("./data/data.json")
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	var dataWeather models.StatusResponse

	errMarshal := json.Unmarshal(jsonByte, &dataWeather)

	if errMarshal != nil {
		fmt.Println(errMarshal.Error())
	}

	return dataWeather
}

func UpdateDataJson() {
	var weather models.StatusResponse

	weather.Status.Water = rand.Intn(100-1+1) + 1
	weather.Status.Wind = rand.Intn(100-1+1) + 1

	result, err := json.Marshal(weather)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("./data/data.json", []byte(result), 0644)

}

func GetStatusWeather(category string, value int) (int, string) {
	var status string = ""
	if category == "water" {
		if value < 5 {
			status = "Aman"
		} else if value >= 6 && value <= 8 {
			status = "Siaga"
		} else if value > 8 {
			status = "Bahaya"
		}
	}

	if category == "wind" {
		if value < 6 {
			status = "Aman"
		} else if value >= 7 && value <= 15 {
			status = "Siaga"
		} else if value > 15 {
			status = "Bahaya"
		}
	}

	return value, status
}
