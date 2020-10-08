package utils

import "encoding/json"

func GetResponseBody(body string) interface{} {
	response := make(map[string]interface{})
	json.Unmarshal([]byte(body), &response)
	return response
}
