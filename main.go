package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	db "monitor/db"
	"monitor/handler"
	"monitor/models"
)

func main() {
	schema := db.GetDbSchema()
	memDB, err := memdb.NewMemDB(schema)
	context := models.Context{
		Db: memDB,
	}
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/endPoints", handler.GetAllEndPoints(context))
	r.POST("/endPoint", handler.PostNewEndPoint(context))
	r.Run()
}
