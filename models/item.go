package models

import (
  "encoding/json"
  "github.com/jinzhu/gorm"
  "github.com/jinzhu/gorm/dialects/postgres"
)

type Item struct {
  gorm.Model
  Type string
  Url string
  Attributes []Attribute `gorm:"-"`
  AttributesB postgres.Jsonb `json:"-", gorm:"type:jsonb;"`
}

type Attribute struct {
  Name string `json:"name"`
  Selector string `json:"selector"`
  Value string `json:"value"`
}

func (i *Item) BeforeSave() error {
	var err error
	i.AttributesB.RawMessage, err = json.Marshal(i.Attributes)
	return err
}

func (i *Item) AfterFind() error {
	var err error
	err = json.Unmarshal(i.AttributesB.RawMessage, &i.Attributes)
	return err
}
