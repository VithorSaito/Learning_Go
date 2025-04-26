package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup() *echo.Echo {
	e := echo.New()
	storage = []Storage{} 
	incrementStorage()

	e.GET("/", getStorage)
	e.POST("/create", createItens)
	e.DELETE("/delete/:Name", deleteItens)

	return e
}

func TestIncrementStorage(t *testing.T) {
	storage = []Storage{}
	incrementStorage()

	assert.Equal(t, 3, len(storage))
	assert.Equal(t, "beans", storage[0].Name)
	assert.Equal(t, 2, storage[0].Amount)
	assert.Equal(t, "rice", storage[1].Name)
	assert.Equal(t, 10, storage[1].Amount)
	assert.Equal(t, "juice", storage[2].Name)
	assert.Equal(t, 4, storage[2].Amount)
}

func TestCreateItem(t *testing.T) {
	e := setup()

	newItem := Storage{Name: "pasta", Amount: 5}
	body, _ := json.Marshal(newItem)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, createItens(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp []Storage
		json.Unmarshal(rec.Body.Bytes(), &resp)

		found := false
		for _, item := range resp {
			if item.Name == "pasta" && item.Amount == 5 {
				found = true
			}
		}
		assert.True(t, found)
	}
}

func TestDeleteExistingItem(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodDelete, "/delete/rice", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("Name")
	c.SetParamValues("rice")

	if assert.NoError(t, deleteItens(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		for _, item := range storage {
			assert.NotEqual(t, "rice", item.Name)
		}
	}
}

func TestDeleteNonExistingItem(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodDelete, "/delete/chocolate", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("Name")
	c.SetParamValues("chocolate")

	if assert.NoError(t, deleteItens(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp map[string]string
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Equal(t, 0, len(resp))
	}
}

func TestGetStorage(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, getStorage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp []Storage
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.Equal(t, 3, len(resp))
	}
}