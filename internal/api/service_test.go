package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test-api-bs/internal/api"
	"test-api-bs/internal/model"
	"testing"
)

func TestGetUserById(t *testing.T) {

	// router_test.go
	testRouter := api.New(api.SetupTestDB())

	tests := make(map[string]int)
	tests["/users/1"] = http.StatusOK
	tests["/users/2"] = http.StatusOK
	tests["/users/6"] = http.StatusNotFound
	tests["/users/abc"] = http.StatusBadRequest
	for key, value := range tests {
		t.Run(key, func(t *testing.T) {
			req, err := http.NewRequest("GET", key, nil)
			if err != nil {
				fmt.Println(err)
			}

			resp := httptest.NewRecorder()
			testRouter.ServeHTTP(resp, req)
			if resp.Code != value {
				t.Fail()
			}
		})
	}

}
func TestGetUsers(t *testing.T) {

	// router_test.go
	testRouter := api.New(api.SetupTestDB())

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	t.Log(resp.Body.String())

	if resp.Code != http.StatusOK {
		t.Fail()
	}

}
func TestGetUserPortfolio(t *testing.T) {

	// router_test.go
	testRouter := api.New(api.SetupTestDB())

	tests := make(map[string]int)
	tests["/user/1/portfolio"] = http.StatusOK
	tests["/user/2/portfolio"] = http.StatusOK
	tests["/user/6/portfolio"] = http.StatusNotFound
	tests["/user/abc/portfolio"] = http.StatusBadRequest
	for key, value := range tests {
		t.Run(key, func(t *testing.T) {
			req, err := http.NewRequest("GET", key, nil)
			if err != nil {
				fmt.Println(err)
			}

			resp := httptest.NewRecorder()
			testRouter.ServeHTTP(resp, req)
			if resp.Code != value {
				t.Fail()
			}
		})
	}
	// req, err := http.NewRequest("GET", "/user/1/portfolio", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// resp := httptest.NewRecorder()
	// testRouter.ServeHTTP(resp, req)

	// t.Log(resp.Body.String())

	// if resp.Code != http.StatusOK {
	// 	t.Fail()
	// }

}
func TestSaveEntry(t *testing.T) {

	testRouter := api.New(api.SetupTestDB())

	var entry model.Entry
	entry.Amount = 100
	entry.Price = 1.23
	entry.TransactionFee = 2
	entry.CoinName = "bitcoin"

	jsonData, _ := json.Marshal(entry)

	req, err := http.NewRequest("POST", "/portfolio/1/entry", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	t.Log(resp.Body.String())

	if resp.Code != http.StatusOK {
		t.Fail()
	}

}
