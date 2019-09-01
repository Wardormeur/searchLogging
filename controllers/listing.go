package controllers

import (
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
  "searchCurator/models"
  "fmt"
  "strconv"
  "net/http"
)
func GetListing(db *gorm.DB) (func(c echo.Context) error) {
  return func (c echo.Context) error {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    listing := models.Listing{}
    if dbc := db.First(&listing, uint(id)); dbc.Error != nil {
      fmt.Println(dbc.Error, &listing, int(id))
      return echo.NewHTTPError(http.StatusNotFound)
    }
    return c.JSON(http.StatusOK, listing)
  }
}

func CreateListing(db *gorm.DB) (func(c echo.Context) error) {
  return func (c echo.Context) error {
    listing := new(models.Listing)
    if err := c.Bind(listing); err != nil {
      fmt.Printf("%v\n",err)
      return err
    }
    if dbc := db.Create(&listing); dbc.Error != nil {
      fmt.Printf("%v\n",dbc.Error)
      return dbc.Error
    }
    return c.JSON(http.StatusCreated, listing)
  }
}
