package model

type ResponseMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
