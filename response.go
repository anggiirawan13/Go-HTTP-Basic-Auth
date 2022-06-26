package main

type DefaultResponse struct {
	Success bool `json:"success"`
	Messages string `json:"messages"`
	Data interface{} `json:"data"`
	StatusCode int32 `json:"status_code"`
}