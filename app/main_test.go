package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func HelloWorld(name string) string {
	return "Hello " + name
}

// func TestGetCustomers(t *testing.T) {
// router := setupRouter()
// router.GET("/api/v1/customer")
// req, _ := http.NewRequest("GET", "/api/v1/customer", nil)
// w := httptest.NewRecorder()
// router.ServeHTTP(w, req)

// var customers []entity.Customer
// json.Unmarshal(w.Body.Bytes(), &customers)

// assert.Equal(t, 200, w.Code)
// assert.NotEmpty(t, customers)

// }

// var router = SetupRouter()

// type LoginInput struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Response struct {
// 	Data LoginInput `json:"data"`
// }

// func TestSuccessLogin(t *testing.T) {
// 	payload := LoginInput{}
// 	payload.Username = "admin"
// 	payload.Password = "admin"

// 	encoded, err := json.Marshal(payload)

// 	if err != nil {
// 		t.FailNow()
// 	}

// 	rr := httptest.NewRecorder()
// 	req, err := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(encoded))
// 	req.Header.Set("Content-Type", "application/json")

// 	if err != nil {
// 		t.FailNow()
// 	}

// 	router.ServeHTTP(rr, req)

// 	res := make(map[string]interface{})
// 	json.NewDecoder(rr.Body).Decode(&res)
// 	parse := res["data"].(map[string]interface{})

// 	assert.Equal(t, rr.Code, 200)
// 	assert.Equal(t, parse["username"], "admin")
// 	assert.Equal(t, parse["password"], "admin")

// }rr
var router = SetupRouter()
var accessToken string

func TestSuccessLogin(t *testing.T) {
	payload := gin.H{
		"user_name": "adit",
		"password":  "adit",
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(encoded))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.FailNow()
	}

	router.ServeHTTP(w, req)
	res := make(map[string]interface{})
	json.NewDecoder(w.Body).Decode(&res)
	parse := res["data"].(map[string]interface{})

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, http.MethodPost, req.Method)
	assert.Equal(t, parse["user_name"], "adit")

	accessToken = parse["access_token"].(string)
}

func TestSuccessGetCustomerById(t *testing.T) {

	var ID int64
	ID = 7
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/customer/"+strconv.Itoa(int(ID)), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	if err != nil {
		t.FailNow()
	}

	router.ServeHTTP(recorder, req)
	res := make(map[string]interface{})
	json.NewDecoder(recorder.Body).Decode(&res)
	logrus.Error(res)
	parse := res["data"].(map[string]interface{})

	assert.Equal(t, float64(ID), parse["id"])
	assert.Equal(t, "Success", res["message"])

}

func TestTableTestHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Adi",
			request:  "Adi",
			expected: "Hello Adi",
		},
		{
			name:     "Aditya",
			request:  "Aditya",
			expected: "Hello Aditya",
		},
		{
			name:     "Adit",
			request:  "Adit",
			expected: "Hello Adit",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
