package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gob4ng/go-sdk/utils"
)

type GenericResponse struct {
	ResponseCode      string `json:"response_code"`
	ResponseMessage   string `json:"response_message"`
	ResponseStatus    bool   `json:"response_status"`
	ResponseTimestamp int64  `json:"response_timestamp"`
}

type GenericResponseData struct {
	ResponseCode      string       `json:"response_code"`
	ResponseMessage   string       `json:"response_message"`
	ResponseStatus    bool         `json:"response_status"`
	ResponseTimestamp int64        `json:"response_timestamp"`
	Data              ResponseData `json:"data"`
}

type ResponseData any

type Response struct {
	ginContext      *gin.Context
	genericResponse GenericResponse
	responseData    any
}

type ResponseBuilder struct {
	Response Response
}

type responseBuilderContext interface {
	SetResponseCode(responseCode string) ResponseBuilder
	SetResponseMessage(responseMessage string) ResponseBuilder
	SetResponseStatus(responseStatus bool) ResponseBuilder
	SetResponseTimestamp(responseTimestamp int64) ResponseBuilder
	SetResponseData(responseData any) ResponseBuilder
	Build() any
}

func NewResponseBuilder(ginContext *gin.Context) responseBuilderContext {
	return ResponseBuilder{Response: Response{
		ginContext: ginContext,
	}}
}

func (r ResponseBuilder) SetResponseCode(responseCode string) ResponseBuilder {
	r.Response.genericResponse.ResponseCode = responseCode
	return r
}

func (r ResponseBuilder) SetResponseMessage(responseMessage string) ResponseBuilder {
	r.Response.genericResponse.ResponseMessage = responseMessage
	return r
}

func (r ResponseBuilder) SetResponseStatus(responseStatus bool) ResponseBuilder {
	r.Response.genericResponse.ResponseStatus = responseStatus
	return r
}

func (r ResponseBuilder) SetResponseTimestamp(responseTimestamp int64) ResponseBuilder {
	r.Response.genericResponse.ResponseTimestamp = responseTimestamp
	return r
}

func (r ResponseBuilder) SetResponseData(responseData any) ResponseBuilder {
	r.Response.responseData = responseData
	return r
}

func (r ResponseBuilder) Build() any {

	responseCode := "0"
	if r.Response.genericResponse.ResponseCode != "" {
		responseCode = r.Response.genericResponse.ResponseCode
	}

	responseMessage := "default"
	if r.Response.genericResponse.ResponseMessage != "" {
		responseMessage = r.Response.genericResponse.ResponseMessage
	}

	responseTimestamp := utils.GetUnixTimestamp()
	if r.Response.genericResponse.ResponseTimestamp != 0 {
		responseTimestamp = r.Response.genericResponse.ResponseTimestamp
	}

	if r.Response.responseData != nil {
		response := GenericResponseData{
			ResponseCode:      responseCode,
			ResponseMessage:   responseMessage,
			ResponseStatus:    r.Response.genericResponse.ResponseStatus,
			ResponseTimestamp: responseTimestamp,
			Data:              r.Response.responseData,
		}

		return response
	}

	response := GenericResponse{
		ResponseCode:      responseCode,
		ResponseMessage:   responseMessage,
		ResponseStatus:    r.Response.genericResponse.ResponseStatus,
		ResponseTimestamp: responseTimestamp,
	}

	return response

}
