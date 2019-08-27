// This package contains all the types associated to the library : weather conditions, cities...
package types

import "encoding/json"

// Structure representing current weather data
type Current struct {
	// The city where the weather applies
	City       City

	// The actual weather conditions at the moment
	Conditions WeatherConditions
}

// Unmarshal the JSON data returned from the API into the structure
func (current *Current) UnmarshalJSON(raw []byte) error {
 	var data map[string]interface{}
 	if err := json.Unmarshal(raw, &data); err != nil {
 		return err
 	}

 	coord := data["coord"].(map[string]interface{})
 	sys := data["sys"].(map[string]interface{})

 	current.City = City{
 		Name:    data["name"].(string),
 		Country: sys["country"].(string),
 		Sunrise: sys["sunrise"].(float64),
 		Sunset:  sys["sunset"].(float64),
 		Coordinates: Coordinates{
 			Longitude: coord["lon"].(float64),
 			Latitude:  coord["lat"].(float64),
		},
 	}
  current.Conditions = ParseWeatherConditions(data)

 	return nil
}
