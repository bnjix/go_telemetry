package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

func main() {
  //db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
  db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=go_telemetry dbname=go_telemetry password=password sslmode=disable")

  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "L1212", Price: 1000})

  // Read
  var product Product
  fmt.Println(db.First(&product, 1)) // find product with id 1
  fmt.Println(db.First(&product, "code = ?", "L1212")) // find product with code l1212

  // Update - update product's price to 2000
  db.Model(&product).Update("Price", 2000)

  // Delete - delete product
  //db.Delete(&product)
}
