package models

import (
  "github.com/jinzhu/gorm"
)

type Listing struct {
  gorm.Model
  Url string
  Selector string
}
