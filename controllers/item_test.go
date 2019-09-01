package controllers 

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

type ItemSuite struct {
  suite.Suite
  payload string
}

func (suite *ItemSuite) SetupTest() {
  suite.payload = `{"url": "https://jobs.com/1", "type": "Job", "attributes": [{"name": "Salary", "selector": "#salary", "value": "90000"}]}`
}
func (suite *ItemSuite) TearDownTest() {
  db := config.InitDB()
  db.LogMode(true)
  db.Unscoped().Delete(&models.Item{})
  db.Exec("ALTER SEQUENCE items_id_seq RESTART WITH 1")
  db.Close()
}
func (suite *ItemSuite) TestCreateItem() {
  t := suite.T()
  res, err := http.Post("http://localhost:1323/items", "application/json", strings.NewReader(suite.payload))

  if assert.NoError(t, err) {
    assert.Equal(t, http.StatusCreated, res.StatusCode)
  }
}

func (suite *ItemSuite) TestGetItem() {
  t := suite.T()
  resCreated, _ := http.Post("http://localhost:1323/items", "application/json", strings.NewReader(suite.payload))
  resCreated.Body.Close()
  res, err := http.Get("http://localhost:1323/items/1")
  model := models.Item{}
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
    assert.Equal(t, "Job", model.Type)
    assert.Equal(t, "https://jobs.com/1", model.Url)
  }
}

func TestItemSuite(t *testing.T) {
  suite.Run(t, new(ItemSuite))
}
