package tests

import (
	"github.com/exybore/goweather"
	"testing"
)

// Gets the forecast and displays all the weather conditions
func TestForecast(t *testing.T) {
		api, err := goweather.NewAPI("", "fr", "metric")
		if err != nil {
			t.Error("error:", err)
			return
		}
    result, err := api.Forecast("Lyon,FR")
    if err != nil {
    	t.Error("error:", err)
    	return
		}

    t.Log(result.City)
    for _, condition := range result.Conditions {
    	t.Log(condition)
	  }
}