package handlers

import "encoding/json"

func ConvertToJson(input interface{}) (string, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
