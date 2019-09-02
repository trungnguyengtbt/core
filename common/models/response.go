package models

import (
	. "github.com/trungnguyengtbt/core/common/models/results"
)

type Response struct {
	Code     int         `json:"code"`
	Messages interface{} `json:"messages"`
	Data     interface{} `json:"data"`
}

func BadRequestResponse(messages interface{}) *Response {
	return &Response{
		Results.InternalError.Code,
		messages,
		nil,
	}
}
func SuccessResponse(data interface{}) *Response {
	return &Response{
		Results.Success.Code,
		Results.Success.Messages,
		data,
	}
}

func InternalErrorResponse() *Response {
	return &Response{
		Results.InternalError.Code,
		Results.InternalError.Messages,
		nil,
	}
}

func SaveErrorResponse() *Response {
	return &Response{
		Results.InternalError.Code,
		Results.InternalError.Messages,
		nil,
	}
}

func ExistDataResponse(existData interface{}) *Response {
	return &Response{
		Results.ExistData.Code,
		Results.ExistData.Messages,
		existData,
	}
}

func NotFoundDataResponse(existData interface{}) *Response {
	return &Response{
		Results.ExistData.Code,
		Results.ExistData.Messages,
		existData,
	}
}

func InvalidAuthenticationResponse() *Response {
	return &Response{
		Results.InvalidAuthentication.Code,
		Results.InvalidAuthentication.Messages,
		nil,
	}
}
