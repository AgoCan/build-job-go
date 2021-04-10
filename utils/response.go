package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Code     int         `json:"code"`
	Error    string      `json:"error"`
	Data     interface{} `json:"data,omitempty"`
	CreateAt time.Time   `json:"create_at"`
}

func NewResponse(code int, errMsg string, data interface{}) *Response {
	return &Response{
		Code:     code,
		Error:    errMsg,
		Data:     data,
		CreateAt: time.Now(),
	}
}

func (r *Response) send(url string, contentType string) (err error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	return err
}
