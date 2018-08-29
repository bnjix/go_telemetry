package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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
	//db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=go_telemetry dbname=go_telemetry password=password sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.AutoMigrate(&DataPoint{})

	// Create
	//db.Create(&DataPoint{Timestamp: 42, Speed: 1000.4})
  var datapoints []DataPoint
  fmt.Println(db.Find(&datapoints))
  fmt.Println(datapoints)
  //db.Find(&users)
  //   c.AbortWithStatus(404)
  //   fmt.Println(err)
  // } else {
  //   c.JSON(200, datapoints)
  // }
  // fmt.Println(db.First(&product, 1))
	// Read
	//var product Product
	//fmt.Println(db.First(&product, 1)) // find product with id 1
	//fmt.Println(db.First(&product, "code = ?", "L1212")) // find product with code l1212

	// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&product)
}
