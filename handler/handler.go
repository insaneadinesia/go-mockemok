package handler

import (
	"encoding/json"

	"github.com/insaneadinesia/go-mockemok/config"
	"github.com/insaneadinesia/go-mockemok/utils"
	"github.com/labstack/echo/v4"
)

const (
	PayloadBody       = "body"
	PayloadQueryParam = "query_param"
	PayloadPathParam  = "path_param"
)

func RequestHandler(c echo.Context) error {
	method := c.Request().Method
	request := c.Get("Request")
	req := request.(config.MockGroupRequest)

	statusCode, resp := GetResponse(c, method, req)
	return c.JSON(statusCode, resp)
}

func GetResponse(c echo.Context, method string, request config.MockGroupRequest) (statusCode int, resp interface{}) {
	statusCode = request.Status
	resp = utils.GetResponseBody(request.Body)

	if request.OverrideBody == nil {
		return
	}

	for _, overrideReq := range request.OverrideBody {
		if IsBodyPayloadMatch(c, overrideReq.Condition) {
			statusCode = overrideReq.Status
			resp = utils.GetResponseBody(overrideReq.Body)

			return
		}
	}

	return
}

func IsBodyPayloadMatch(c echo.Context, condition config.OverrideCondition) bool {
	key := condition.PayloadKey
	value := condition.PayloadValue

	switch condition.PayloadFrom {
	case PayloadPathParam:
		if c.Param(key) != "" && c.Param(key) == value {
			return true
		}

		break

	case PayloadQueryParam:
		if c.QueryParam(key) != "" && c.QueryParam(key) == value {
			return true
		}

		break

	case PayloadBody:
		req := make(map[string]interface{})
		bodyRequest, _ := utils.GetRequestBody(c)

		json.Unmarshal(bodyRequest, &req)

		// First : Check if request is raw json
		if req[key] != "" && req[key] == value {
			return true
		}

		// Second : Check if request is form data
		if c.FormValue(key) != "" && c.FormValue(key) == value {
			return true
		}

		break
	}

	return false
}
