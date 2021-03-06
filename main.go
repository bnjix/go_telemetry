package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	_ "time"
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
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	r.Run()
}

func GetDataPoints(c *gin.Context) {
	var datapoints []DataPoint
	if err := db.Find(&datapoints).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	} else {
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
		broadcast <- data_point
		c.JSON(200, data_point)
	}
}

func GetHtml(c *gin.Context) {
	var datapoints []DataPoint
	if err := db.Order("ID desc").Limit(30).Find(&datapoints).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	} else {
		dps, _ := json.Marshal(datapoints)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"allDataPoints": string(dps),
		})
	}
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var broadcast = make(chan DataPoint)

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		msg := <-broadcast
		conn.WriteJSON(msg)
	}
}
