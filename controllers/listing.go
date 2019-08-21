package controllers

import (
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
  "../models"
  "strconv"
  "net/http"
)
func GetListing(db *gorm.DB) (func(c echo.Context) error) {
  return func (c echo.Context) error {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    listing := models.Listing{}
    if dbc := db.First(&listing, id); dbc.Error != nil {
      return echo.NewHTTPError(http.StatusNotFound)
    }
    return c.JSON(http.StatusOK, listing)
  }
}

func CreateListing(db *gorm.DB) (func(c echo.Context) error) {
  return func (c echo.Context) error {
    listing := new(models.Listing)
    if err := c.Bind(listing); err != nil {
      return err
    }
    return c.JSON(http.StatusCreated, listing)
  }
}
