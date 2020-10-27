package utils

import (
	"encoding/json"

	"github.com/johnearl92/xendit-ta.git/internal/model/errors"
)

// FromJSON unmarshall to output
func FromJSON(in []byte, out interface{}) errors.JSONErrors {
	err := json.Unmarshal(in, out)
	if err != nil {
		return errors.New().Add(
			"400",
			map[string]string{"pointer": "/data"},
			"Request cannot be read as JSON",
			err.Error())
	}
	return nil
}

// ToJSON marshalls input to JSON
func ToJSON(in interface{}) ([]byte, error) {
	out, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	return out, nil
}
