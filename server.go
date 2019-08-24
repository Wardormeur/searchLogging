package main

import (
  "searchCurator/controllers"
  "searchCurator/config"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  e := echo.New()
  db := config.InitDB()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.POST("/listings", controllers.CreateListing(db))
  e.GET("/listings/:id", controllers.GetListing(db))

  e.Logger.Fatal(e.Start(":1323"))
  defer config.CloseConnection(db)
}
