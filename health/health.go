package health

import (
	"log"
	"monitor/models"
	"net/http"
	"time"
)

func GetEndPointHealth(endPoint models.EndPoint) *models.Health {
	resp, err := http.Get(endPoint.URL)
	health := &models.Health{
		EndPointId: endPoint.Id,
		Match:      false,
		Time:       time.Now().Format(time.RFC850),
	}
	if err != nil {
		return health
	}
	if resp.StatusCode == endPoint.ResponseCode {
		health.Match = true
	}
	return health
}

func GetEndPointsHealth(endPoints models.EndPoints) []*models.Health {
	var results []*models.Health
	for _, e := range endPoints {
		health := GetEndPointHealth(*e)
		results = append(results, health)
	}
	return results
}

func RecordHealth(ctx models.Context) {
	endpoints, _ := GetAllEndPoints(ctx)
	txn := ctx.Db.Txn(true)
	healthRecords := GetEndPointsHealth(endpoints)
	for _, record := range healthRecords {
		if err := txn.Insert("health", record); err != nil {
			log.Println(err)
		}
	}
	txn.Commit()
}

func GetAllEndPoints(ctx models.Context) (models.EndPoints, error) {
	txn := ctx.Db.Txn(false)
	var endpoints = models.EndPoints{}
	it, err := txn.Get("endPoint", "id")
	if err != nil {
		log.Println(err)
	}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*models.EndPoint)
		endpoints = append(endpoints, p)
	}
	txn.Abort()
	return endpoints, err
}
