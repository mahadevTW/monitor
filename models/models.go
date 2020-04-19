package models

import (
	"github.com/hashicorp/go-memdb"
)

type EndPoint struct {
	Id           string `json:"id"`
	ServiceName  string `json:"serviceName"`
	URL          string `json:"url"`
	ResponseCode int    `json:"responseCode"`
}

type Health struct {
	EndPointId string `json:"endpointId"`
	Match      bool   `json:"match"`
	Time       string `json:"time"`
}

type EndPoints []*EndPoint
type EndPointHealthResponse struct {
	EndPoint
	Health *Health `json:"health"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type Context struct {
	Db *memdb.MemDB
}

func GetSuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func GetFailureResponse(data interface{}) Response {
	return Response{
		Success: false,
		Data:    data,
	}
}
