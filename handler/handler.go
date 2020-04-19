package handler

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"log"
	"monitor/health"
	"monitor/models"
	"net/http"
)

func PingGet() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func HealthGet(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		endPointId := c.Param("id")
		txn := ctx.Db.Txn(false)
		defer txn.Abort()
		it, err := txn.First("health", "id", endPointId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.GetFailureResponse(err))
			return
		}
		c.JSON(http.StatusOK, models.GetSuccessResponse(it.(*models.Health)))
	}
}

func HealthGetAllEndPoints(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		endPoints, err := health.GetAllEndPoints(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.GetFailureResponse(err))
		}
		txn := ctx.Db.Txn(false)
		defer txn.Abort()
		var response []models.EndPointHealthResponse
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.GetFailureResponse(err))
		}
		for _, e := range endPoints {
			obj, err := txn.First("health", "id", e.Id)
			if err != nil {
				log.Println(err)
			}
			p := obj.(*models.Health)
			response = append(response, models.EndPointHealthResponse{
				EndPoint: *e,
				Health:   p,
			})
		}
		c.JSON(http.StatusOK, models.GetSuccessResponse(response))
	}
}

func getMatchingHealthRecord(endPoint *models.EndPoint, healths []*models.Health) *models.Health {
	for _, h := range healths {
		if h.EndPointId == endPoint.Id {
			return h
		}
	}
	return nil
}

func GetAllEndPoints(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		txn := ctx.Db.Txn(false)
		defer txn.Abort()
		it, err := txn.Get("endPoint", "id")
		if err != nil {
			log.Println(err)
		}
		var endpoints = models.EndPoints{}
		for obj := it.Next(); obj != nil; obj = it.Next() {
			p := obj.(*models.EndPoint)
			endpoints = append(endpoints, p)
		}
		c.JSON(200, models.GetSuccessResponse(endpoints))
	}
}

func PostNewEndPoint(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var endPoint models.EndPoint
		err := c.BindJSON(&endPoint)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.GetFailureResponse(err))
			return
		}
		uuid, _ := uuid.GenerateUUID()
		endPoint.Id = uuid
		txn := ctx.Db.Txn(true)
		if err := txn.Insert("endPoint", &endPoint); err != nil {
			c.JSON(http.StatusInternalServerError, models.GetFailureResponse(err))
			panic(err)
		}
		txn.Commit()
		health.RecordHealth(ctx)
		c.JSON(http.StatusOK, models.GetSuccessResponse(endPoint))
	}
}

func GetAssets() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile("./web/build", true))
}
