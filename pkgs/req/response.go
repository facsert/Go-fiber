package req

import "fmt"

type Response struct {
	Data   any    `json:"data"`
	Result bool   `json:"result"`
	Msg    string `json:"msg"`
}

func NewResp(data any) *Response {
	return &Response{
		Data:   data,
		Result: true,
		Msg:    "success",
	}
}

func (r *Response) SetResult(result bool) *Response {
	r.Result = result
	return r
}

func (r *Response) SetMsg(msg string, a ...any) *Response {
	r.Msg = fmt.Sprintf(msg, a...)
	return r
}
