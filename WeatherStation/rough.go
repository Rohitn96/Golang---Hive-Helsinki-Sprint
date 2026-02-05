package main

import "fmt"

type WeatherStation struct {
	id int
	key string
	currentvalue *float64
}

func main() {

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
				var userInput string
				fmt.Scanln(&userInput)

				// Parse id,value ahead of time
				var id int
				var value float64
				_, err := fmt.Sscanf(userInput, "%d,%f", &id, &value)

				switch {
				case userInput == "exit":
					fmt.Println("Exiting...")
					return
				case userInput == "get":
					Display(station)
				case userInput == "clear":
					Reset(station)
				case err == nil: // Successfully parsed as id,value
					SetValue(station, id, value)
				default:
					fmt.Println("Invalid input. Use get, clear, exit, or id,value (e.g., 1,23.5)")
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
