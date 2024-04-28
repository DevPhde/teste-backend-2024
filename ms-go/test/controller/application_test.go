package controller

import (
	"encoding/json"
	"ms-go/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestIndexHome(t *testing.T) {
	expectedBody := map[string]interface{}{
		"message": "[ms-go] | Success",
		"status":  float64(200),
	}

	router := router.SetupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, err)

	assert.Equal(t, expectedBody, responseBody)
}
