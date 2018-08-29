package main

import (
	//"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type DataPoint struct {
	gorm.Model
	Timestamp          uint
	Latitude           float32
	Longitude          float32
	Altitude           float32
	Speed              float32
	Course             float32
	Heading            float32
	VerticalAccuracy   float32
	HorizontalAccuracy float32
	TrueHeading        float32
	AccelerationX      float32
	AccelerationY      float32
	AccelerationZ      float32
	GyroRotationX      float32
	GyroRotationY      float32
	GyroRotationZ      float32
	RelativeAltitude   float32
	BatteryLevel       float32
}

func main() {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=go_telemetry dbname=go_telemetry password=password sslmode=disable")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&DataPoint{})

	r := gin.Default()

	r.GET("/data_points", GetDataPoints)
	r.GET("/ping", GetPing)
	r.POST("/data_points", CreateDataPoint)
	r.Run()
}

func GetDataPoints(c *gin.Context) {
	var datapoints []DataPoint
	if err := db.Find(&datapoints).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	} else {
		for i := 0; i < len(datapoints); i++{
			datapoints[i].Timestamp = uint(datapoints[i].CreatedAt.Unix())
		}
		c.JSON(200, datapoints)
	}
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateDataPoint(c *gin.Context) {
	data_point := DataPointFromJson(c.Params)
	//db.Create(&data_point)
	//c.JSON(200, data_point)
	c.JSON(200, gin.H{})
}

func DataPointFromJson(params gin.Params) DataPoint {
	return DataPoint{}
}
