package main

import (
	"net/http"
	_ "time"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type DataPoint struct {
	gorm.Model
	Timestamp          float32 `form:"locationTimestamp_since1970"`
	Latitude           float32 `form:"locationLatitude"`
	Longitude          float32 `form:"locationLongitude"`
	Altitude           float32 `form:"locationAltitude"`
	Speed              float32 `form:"locationSpeed"`
	Course             float32 `form:"locationCourse"`
	Heading            float32 `form:"locationTrueHeading"`
	VerticalAccuracy   float32 `form:"locationVerticalAccuracy"`
	HorizontalAccuracy float32 `form:"locationHorizontalAccuracy"`
	AccelerationX      float32 `form:"accelerometerAccelerationX"`
	AccelerationY      float32 `form:"accelerometerAccelerationY"`
	AccelerationZ      float32 `form:"accelerometerAccelerationZ"`
	GyroRotationX      float32 `form:"gyroRotationX"`
	GyroRotationY      float32 `form:"gyroRotationY"`
	GyroRotationZ      float32 `form:"gyroRotationZ"`
	RelativeAltitude   float32 `form:"altimeterRelativeAltitude"`
	BatteryLevel       float32 `form:"batteryLevel"`
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

	r.LoadHTMLGlob("templates/*")
	r.GET("/data_points_html", GetHtml)

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
			//datapoints[i].Timestamp = uint(datapoints[i].CreatedAt.Unix())
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
	var data_point DataPoint
	if err := c.Bind(&data_point); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
	} else {
		data_point.DeletedAt = nil
		db.Create(&data_point)
		c.JSON(200, data_point)
	}
}

func GetHtml(c *gin.Context) {
	var datapoints []DataPoint
	if err := db.Order("ID desc").Limit(100).Find(&datapoints).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	} else {
		dps, _ := json.Marshal(datapoints)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"allDataPoints": string(dps),
		})
	}
}
