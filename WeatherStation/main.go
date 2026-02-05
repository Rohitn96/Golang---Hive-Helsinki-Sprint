package main

import "fmt"
import "os"

type WeatherStation struct {
	id int
	key string
	currentvalue *float64
}

func main() {
	var userInput string

	airTemp := WeatherStation{1, "airTemp", nil}
	airPressure := WeatherStation{2, "airPressure", nil}
	precipitation := WeatherStation{7, "precipitation", nil}
	windSpeed := WeatherStation{11, "windSpeed", nil}
	windDirection := WeatherStation{12, "windDirection", nil}
	humidity := WeatherStation{13, "humidity", nil}
	dewPoint := WeatherStation{14, "dewPoint", nil }
	soilMoisture := WeatherStation{15, "soilMoisture", nil}
	cloudCover := WeatherStation{22, "cloudCover", nil}

	station := []*WeatherStation{&airTemp, &airPressure, &precipitation, &windSpeed,
		&windDirection, &humidity, &dewPoint, &soilMoisture, &cloudCover}
 		//  * pointers is for indicating the position of the value and indicating tha the specific type is getting changed
   		// & is  for indicating the memory position of value
		fmt.Println("--- Weather Station ---")
	for {
		fmt.Scanln(&userInput)
		if userInput == "" {
			fmt.Println("Type 'get', 'clear' or 'exit'")
		} else if userInput == "get" {
			Display(station)
		} else if userInput == "clear" {
			Reset(station)
		} else if userInput == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}else {
			var id int
			var raw string

			_, err := fmt.Sscanf(userInput, "%d,%s", &id, &raw)
			if err != nil {
				continue
			}
			if raw == "NULL" {
				for _, v := range station {
					if v.id == id {
						v.currentvalue = nil
						break
					}
				}
			} else {
			var value float64
			if _, err := fmt.Sscanf(raw, "%f", &value); err != nil {
				continue
			}
			SetValue(station, id, value)
}
			
		}
	}
}


func Reset(station []*WeatherStation) {
		for _, v := range station{
			v.currentvalue = nil
		}
}

func SetValue(station []*WeatherStation, id int, value float64) {
	for _,v := range station {
		if id == v.id {
			//	newvalue := value
			v.currentvalue = &value
			//fmt.Println(value)
			return
		}
	}
}

func Display(station []*WeatherStation) {
	var result string
	for _, v := range station{
		if v.currentvalue == nil {
			result = "NULL"
		}else{
			result = fmt.Sprintf("%g", *v.currentvalue)
		//	fmt.Println(*v.currentvalue)
		}
		// result = (v.currentvalue)
		fmt.Printf("%s:%s\n", v.key, result)
	}
}
