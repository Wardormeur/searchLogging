package main

import (
  "net/http"
  "testing"
  "strings"
  "encoding/json"
  "searchCurator/models"
  "searchCurator/config"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type ListingSuite struct {
  suite.Suite
  payload string
}

func (suite *ListingSuite) SetupTest() {
  suite.payload = `{"url": "https://indeed.com", "selector": "body>div" }`
}
func (suite *ListingSuite) TearDownTest() {
  db := config.InitDB()
  // db.LogMode(true)
  db.Unscoped().Delete(&models.Listing{})
  db.Exec("ALTER SEQUENCE listings_id_seq RESTART WITH 1")
  db.Close()
}
func (suite *ListingSuite) TestCreateListing() {
  t := suite.T()
  res, err := http.Post("http://localhost:1323/listings", "application/json", strings.NewReader(suite.payload))

  if assert.NoError(t, err) {
    assert.Equal(t, http.StatusCreated, res.StatusCode)
  }
}

func (suite *ListingSuite) TestGetListing() {
  t := suite.T()
  resCreated, _ := http.Post("http://localhost:1323/listings", "application/json", strings.NewReader(suite.payload))
  resCreated.Body.Close()
  res, err := http.Get("http://localhost:1323/listings/1")
  model := models.Listing{}
  defer res.Body.Close()
  if assert.NoError(t, err) {
    d := json.NewDecoder(res.Body)
    d.DisallowUnknownFields()
    d.Decode(&model)
    // Ensure creation went fine
    assert.Equal(t, http.StatusCreated, resCreated.StatusCode)
    // Ensure API data is returned
    assert.Equal(t, http.StatusOK, res.StatusCode)
    assert.Equal(t, uint(1), model.ID)
    assert.Equal(t, "https://indeed.com", model.Url)
    assert.Equal(t, "body>div", model.Selector)
  }
}

func TestListingSuite(t *testing.T) {
  suite.Run(t, new(ListingSuite))
}
