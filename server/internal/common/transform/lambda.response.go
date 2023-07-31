package transform

import "encoding/json"

type LambdaResponseModel struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	MessageId string      `json:"message_id"`
	Code      int         `json:"code"`
	Error     error       `json:"error"`
}

func (l LambdaResponseModel) ToJson() string {
	res, _ := json.Marshal(l)
	return string(res)
}
