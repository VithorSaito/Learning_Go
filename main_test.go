package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestIncrementStorage(t *testing.T) {
	storage = []Storage{} // Resetando o storage
	incrementStorage()

	if len(storage) != 3 {
		t.Errorf("Esperado 3 itens no storage, mas tem %d", len(storage))
	}
}

func TestGetStorage(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage = []Storage{{Name: "test", Amount: 5}}
	err := getStorage(c)

	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Esperado status 200, mas recebeu %d", rec.Code)
	}
}

func TestCreateItens(t *testing.T) {
	e := echo.New()
	body := `{"Name":"bread","Amount":3}`
	req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	storage = []Storage{}

	err := createItens(c)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if len(storage) != 1 || storage[0].Name != "bread" {
		t.Errorf("Item não foi criado corretamente. Storage: %+v", storage)
	}
}

func TestDeleteItensSuccess(t *testing.T) {
	e := echo.New()
	storage = []Storage{{Name: "milk", Amount: 1}}

	req := httptest.NewRequest(http.MethodDelete, "/delete/milk", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("Name")
	c.SetParamValues("milk")

	err := deleteItens(c)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if len(storage) != 0 {
		t.Errorf("Item não foi deletado corretamente. Storage: %+v", storage)
	}
}

func TestDeleteItensNotFound(t *testing.T) {
	e := echo.New()
	storage = []Storage{{Name: "water", Amount: 2}}

	req := httptest.NewRequest(http.MethodDelete, "/delete/juice", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("Name")
	c.SetParamValues("juice")

	err := deleteItens(c)
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}

	if len(storage) != 1 {
		t.Errorf("Storage foi alterado incorretamente. Storage: %+v", storage)
	}
}
