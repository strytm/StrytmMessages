package messages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-resty/resty/v2"
)

var isDebug = true

type ResponseModelStruct struct {
	Result struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Detail     interface{} `json:"detail,omitempty"`
		FileName   string      `json:"file_name,omitempty"`
		LineNumber int         `json:"line_number,omitempty"`
	}
}

type MessageModelStruct struct {
	ResponseWriter http.ResponseWriter

	JsonModel ResponseModelStruct

	JsonModelForStruct struct {
		Result interface{} `json:"result"`
	}

	JsonModelForStructWithPageModel struct {
		Result interface{} `json:"result"`
		Paging interface{} `json:"paging"`
	}
}

func (data *MessageModelStruct) ForwardResponse(resp *resty.Response) {
	var s = new(ResponseModelStruct)
	json.Unmarshal(resp.Body(), &s)
	data.ShowStringMessageAndStatusCode(s.Result.Message, resp.StatusCode(), s.Result.Detail)
}

func (data *MessageModelStruct) ShowResultJson(s interface{}) {

	if isDebug {
		_, fn, line, _ := runtime.Caller(1)

		data.JsonModel.Result.FileName = fn
		data.JsonModel.Result.LineNumber = line
	}

	if data.ResponseWriter == nil {
		panic("ResponseWriter set nashode ast")
	}

	data.ResponseWriter.WriteHeader(http.StatusOK)

	data.JsonModelForStruct.Result = s

	stringBytes, err := json.Marshal(data.JsonModelForStruct)

	if err != nil {
		data.ResponseWriter.WriteHeader(http.StatusNotImplemented)
		_, err = data.ResponseWriter.Write([]byte("moshkeli dar tabdil shodan be json pish omade"))
		if err != nil {

			fmt.Println(err)
		}
	}

	_, err = data.ResponseWriter.Write(stringBytes)
	if err != nil {

		fmt.Println(err)
	}

}

func (data *MessageModelStruct) ShowResultWithPageJson(s interface{}, p interface{}) {

	if isDebug {
		_, fn, line, _ := runtime.Caller(1)

		data.JsonModel.Result.FileName = fn
		data.JsonModel.Result.LineNumber = line
	}

	if data.ResponseWriter == nil {
		panic("ResponseWriter set nashode ast")
	}

	data.ResponseWriter.WriteHeader(http.StatusOK)

	data.JsonModelForStructWithPageModel.Result = s
	data.JsonModelForStructWithPageModel.Paging = p

	stringBytes, err := json.Marshal(data.JsonModelForStructWithPageModel)

	if err != nil {
		data.ResponseWriter.WriteHeader(http.StatusNotImplemented)
		_, err = data.ResponseWriter.Write([]byte("moshkeli dar tabdil shodan be json pish omade"))
		if err != nil {

			fmt.Println(err)
		}
	}

	_, err = data.ResponseWriter.Write(stringBytes)
	if err != nil {

		fmt.Println(err)
	}

}

func (data *MessageModelStruct) ShowStringMessageAndStatusCode(message string, statusCode int, detail interface{}) {

	if isDebug {
		_, fn, line, _ := runtime.Caller(1)
		data.JsonModel.Result.FileName = fn
		data.JsonModel.Result.LineNumber = line
	}

	if data.ResponseWriter == nil {
		panic("ResponseWriter set nashode ast")
	}

	data.ResponseWriter.WriteHeader(statusCode)

	data.JsonModel.Result.StatusCode = statusCode
	data.JsonModel.Result.Message = message
	data.JsonModel.Result.Detail = detail

	stringBytes, err := json.Marshal(data.JsonModel)

	if err != nil {
		data.ResponseWriter.WriteHeader(http.StatusNotImplemented)
		_, err = data.ResponseWriter.Write([]byte("moshkeli dar tabdil shodan be json pish omade"))
		if err != nil {

			fmt.Println(err)
		}
	}

	_, err = data.ResponseWriter.Write(stringBytes)
	if err != nil {

		fmt.Println(err)
	}

}
