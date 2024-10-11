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

func (re *Error) Error() string {
	return fmt.Sprintf("request error: statuscode: %d, transactionID: %s, reference: %s, resultcode: %s, result: %s, message: %s, data: %v", re.statuscode, re.TransactionId, re.Reference, re.ResultCode, re.Result, re.Message, re.Data)
}
