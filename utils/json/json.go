package json

import "encoding/json"

func ToJSON(v interface{}) (string, error) {
	if result, err := json.Marshal(v); err == nil {
		return string(result), err
	} else {
		return "", err
	}
}
