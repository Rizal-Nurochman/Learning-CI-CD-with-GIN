package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rizal-Nurochman/repository"
	"github.com/gin-gonic/gin"
)

func setUpTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode) // matikan log berisik Gin
	_ = repository.ConnectDB()

	r := gin.New()
	r.GET("/", HomeHandler)
	r.GET("/books", GetAll)
	return r
}

func TestHomeHandler(t *testing.T)  {
	r := setUpTestRouter()
	w := httptest.NewRecorder() // fake server

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d body: %s", w.Code, w.Body.String())
	}
}

func TestGetAll(t *testing.T) {
	r := setUpTestRouter()
	w := httptest.NewRecorder() //fake server

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d body: %s", w.Code, w.Body.String())
	}
}