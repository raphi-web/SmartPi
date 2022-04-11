package sensor

import (
	"encoding/json"
	"net/http"
)

type Sensor struct {
	Type  string
	Unit  string
	Value float32
}

func SensorFromRequest(r *http.Request) (Sensor, error) {
	var s Sensor

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil {
		return s, err
	}
	return s, nil
}
