package controllers

import (
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
  "searchCurator/models"
  "fmt"
  "strconv"
  "net/http"
)
func GetItem(db *gorm.DB) (func(c echo.Context) error) {
  return func (c echo.Context) error {
    id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
    item := models.Item{}
    if dbc := db.First(&item, uint(id)); dbc.Error != nil {
      fmt.Println(dbc.Error, &item, int(id))
      return echo.NewHTTPError(http.StatusNotFound)
    }
    return c.JSON(http.StatusOK, item)
  }
}

func CreateItem(db *gorm.DB) (func(c echo.Context) error) {
  return func (c echo.Context) error {
    item := new(models.Item)
    if err := c.Bind(item); err != nil {
      fmt.Printf("%v\n",err)
      return err
    }
    if dbc := db.Create(&item); dbc.Error != nil {
      fmt.Printf("%v\n",dbc.Error)
      return dbc.Error
    }
    return c.JSON(http.StatusCreated, item)
  }
}
