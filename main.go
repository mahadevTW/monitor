package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb"
	"github.com/jasonlvhit/gocron"
	db "monitor/db"
	"monitor/handler"
	"monitor/health"
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
	//some dummy endpoints Endpoints
	insertDummyEndPoints(context)
	go health.RecordHealth(context)
	gocron.Every(5).Second().Do(health.RecordHealth, context)
	gocron.Start()
	r := gin.Default()
	r.Use(handler.GetAssets())
	apiRoutes := r.Group("/api")
	apiRoutes.GET("/ping", handler.PingGet())
	apiRoutes.GET("/endPoints", handler.GetAllEndPoints(context))
	apiRoutes.POST("/endPoint", handler.PostNewEndPoint(context))
	apiRoutes.GET("/endPoint/health", handler.HealthGetAllEndPoints(context))
	apiRoutes.GET("/endPoint/health/:id", handler.HealthGet(context))
	r.Run()
}

func insertDummyEndPoints(context models.Context) {
	endPoints := []*models.EndPoint{
		&models.EndPoint{
			Id:           "some-random-other-id",
			URL:          "http://localhost:8080/api/ping",
			ResponseCode: 200,
			ServiceName:  "Document Service",
		},
		&models.EndPoint{
			Id:           "some-random-id",
			URL:          "http://localhost:8080/api/ping",
			ResponseCode: 200,
			ServiceName:  "Auth Service",
		},
	}
	txn := context.Db.Txn(true)
	for _, endPoint := range endPoints {
		err := txn.Insert("endPoint", endPoint)
		if err != nil {
			panic(err)
		}
	}
	txn.Commit()
}
