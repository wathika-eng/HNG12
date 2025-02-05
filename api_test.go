package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/api/classify-number", classifyNumber)
	return r
}

func TestPingEndpoint(t *testing.T) {
	r := setupRouter()

	// Create a test HTTP request for /ping
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Record the response using httptest
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the status code is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the body contains {"message": "pong"}
	expected := `{"message":"pong"}`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestClassifyNumberValidQuery(t *testing.T) {
	r := setupRouter()

	// Create a test HTTP request with a valid number
	req, _ := http.NewRequest("GET", "/api/classify-number?number=371", nil)

	// Record the response using httptest
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the status code is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert that the response JSON contains the correct number properties
	// You can modify this based on the actual expected output for the number 371
	// For example, if 371 is a prime number
	expected := `{"number":371,"is_prime":true,"is_perfect":false,"properties":["odd"],"digit_sum":11,"fun_fact":"371 is a prime number"}`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestClassifyNumberBlankQuery(t *testing.T) {
	r := setupRouter()

	// Create a test HTTP request with a blank query
	req, _ := http.NewRequest("GET", "/api/classify-number?number=", nil)

	// Record the response using httptest
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the status code is 400
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert the error message is correct
	expected := `{"number":"","error":"query param cannot be blank"}`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestClassifyNumberInvalidQuery(t *testing.T) {
	r := setupRouter()

	// Create a test HTTP request with an invalid number
	req, _ := http.NewRequest("GET", "/api/classify-number?number=abc", nil)

	// Record the response using httptest
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the status code is 400
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert the error message is correct
	expected := `{"number":0,"error":"query param cannot be converted to an interger"}`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestClassifyNumberNegativeQuery(t *testing.T) {
	r := setupRouter()

	// Create a test HTTP request with a negative number
	req, _ := http.NewRequest("GET", "/api/classify-number?number=-5", nil)

	// Record the response using httptest
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the status code is 400
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert the error message is correct
	expected := `{"number":-5,"error":"number cannot be less than zero"}`
	assert.JSONEq(t, expected, w.Body.String())
}
