package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func GetResponseBody(body string) interface{} {
	response := make(map[string]interface{})
	json.Unmarshal([]byte(body), &response)
	return response
}

func GetRequestBody(c echo.Context) (body []byte, err error) {
	if c.Request().Body != nil { // Read
		body, err = ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return body, err
		}
	}

	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, err
}
