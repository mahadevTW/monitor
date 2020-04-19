package models

import "github.com/hashicorp/go-memdb"

type EndPoint struct {
	URL          string `json:"url"`
	ResponseCode int    `json:"responseCode"`
}

type EndPoints []*EndPoint

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
