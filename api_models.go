package client

import "fmt"

type Error struct {
	TransactionId string         `json:"transid"`
	Reference     string         `json:"reference"`
	ResultCode    string         `json:"resultcode"`
	Result        string         `json:"result"`
	Message       string         `json:"message"`
	Data          map[string]any `json:"data"`
	statuscode    int
}

type Response struct {
	Reference  string           `json:"reference"`
	ResultCode string           `json:"resultcode"`
	Result     string           `json:"result"`
	Message    string           `json:"message"`
	Data       []map[string]any `json:"data"`
}

func (re *Error) Error() string {
	return fmt.Sprintf("request error: statuscode: %d, message: %s, data: %v", re.statuscode, re.Message, re.Data)
}
