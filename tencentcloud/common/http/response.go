package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	// "log"
	"net/http"

	"github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error" omitempty`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

type DeprecatedAPIErrorResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	CodeDesc  string `json:"codeDesc"`
	RequestId string `json:"RequestId"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		return
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(strconv.Itoa(deprecated.Code), deprecated.Message, deprecated.RequestId)
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		return
	}
	log.Printf("[DEBUG] Response Body=%s", body)

	// 兼容旧版本协议
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return
	}
	var obj []byte
	m := f.(map[string]interface{})
	if _, ok := m["Response"]; !ok {
		if m["code"].(float64) != 0 {
			var b interface{}
			err = json.Unmarshal(body, &b)
			n := b.(map[string]interface{})

			err = errors.NewTencentCloudSDKError(strconv.FormatFloat(n["code"].(float64), 'f', -1, 64), n["message"].(string), n["requestId"].(string))
			if err != nil {
				return
			}
		} else {
			obj = []byte(`{"Response":` + string(body[:]) + "}")
		}

	} else {
		obj = body
	}

	err = response.ParseErrorFromHTTPResponse(obj)
	if err != nil {
		return
	}

	err = json.Unmarshal(obj, &response)

	return
}
