package main

import (
	"github.com/mihailo-misic/company-resource-api/database"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	router := setupRouter()

	database.Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Body)
}
