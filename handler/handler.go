package handler

import (
	"github.com/gin-gonic/gin"
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

func GetAllEndPoints(ctx models.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		txn := ctx.Db.Txn(false)
		defer txn.Abort()
		it, err := txn.Get("endPoint", "id")
		if err != nil {
			panic(err)
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
		txn := ctx.Db.Txn(true)
		if err := txn.Insert("endPoint", &endPoint); err != nil {
			c.JSON(http.StatusInternalServerError, models.GetFailureResponse(err))
			panic(err)
		}
		txn.Commit()
		c.JSON(http.StatusOK, models.GetSuccessResponse(""))
	}
}
