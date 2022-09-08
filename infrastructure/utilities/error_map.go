package utilities

import (
	"encoding/json"
	"errors"
)

func ConvertMapToError(err map[string]interface{}) error {
	js, jsErr := json.Marshal(err)
	if jsErr != nil {
		return errors.New("Unable to Parse Error JSON.")
	}
	if len(err) != 0 {
		return errors.New(string(js))
	}
	return nil
}
