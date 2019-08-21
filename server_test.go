package handler

import (
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"

  "github.com/labstack/echo"
  "github.com/stretchr/testify/assert"
)

var (
  mockDB = map[int]*Listing{
    1: &Listing{1, "http://indeed.com", "body>div"},
  }
  listingJSON = `{"id": 1,"url":"http://indeed.com","itemSelector":"body>div"}`
)

func TestCreateListing(t *testing.T) {
  e := echo.New()
  req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(listingJSON))
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
  rec := httptest.NewRecorder()
  c := e.NewContext(req, rec)
  h := &handler{mockDB}
  
  if assert.NoError(t, h.createListing(c)) {
    assert.Equal(t, http.StatusCreated, rec.Code)
    assert.Equal(t, listingJSON, rec.Body.String())
  }
}
