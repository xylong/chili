package chili

import "github.com/gin-gonic/gin"

type Response struct {
	data interface{}
}

func R() *Response {
	return &Response{}
}

func (r *Response) Json(data gin.H) *Response {
	r.data = data
	return r
}

func (r *Response) String(data string) *Response {
	r.data = data
	return r
}

func (r *Response) Byte(data []byte) *Response {
	r.data = data
	return r
}

func (r *Response) Data() interface{} {
	return r.data
}
