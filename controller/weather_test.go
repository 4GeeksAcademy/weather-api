package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/robertoduessmann/weather-api/model"
)

func TestCurrentWeather(t *testing.T) {
	var weather model.Weather
	var blankWeather model.Weather

	router := new(mux.Router)
	router.HandleFunc("/weather/{city}", CurrentWeather)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/weather/Curitiba", nil)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected request to be OK.")
		t.FailNow()
	}

	if err := json.Unmarshal(w.Body.Bytes(), &weather); err != nil {
		t.Errorf("Unexpected error while unmarshal response.")
		t.FailNow()
	}

	if weather == blankWeather {
		t.Errorf("Expected weather information NOT to be empty.")
		t.FailNow()
	}
}
