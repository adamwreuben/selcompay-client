package client

import (
	"encoding/json"
)

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

// Error is the implementation of the error interface.
func (re Error) Error() string {
	d, err := json.Marshal(re)
	if err != nil {
		return err.Error()
	}

	return string(d)
}
