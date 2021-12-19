package api_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"test-api-bs/internal/api"
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

}
func TestSaveEntry(t *testing.T) {

	testRouter := api.New(api.SetupTestDB())

	type testData struct {
		name     string
		endpoint string
		payload  string
		expect   int
	}

	var tests []testData
	tests = append(tests, testData{"Correct Entry", "/portfolio/1/entry", `{"coinName":"bitcoin", "amount": 10, "price": 123.23, "transactionFee": 10}`, 200})
	tests = append(tests, testData{"Incorrect Entry", "/portfolio/1/entry", `{"coinName":"bitcoin", "amount": "10", "price": "abc", "transactionFee": 10}`, 400})
	tests = append(tests, testData{"Incorrect Entry Wrong Portfolio", "/portfolio/abc/entry", `{"coinName":"bitcoin", "amount": "10", "price": 10, "transactionFee": 10}`, 400})
	tests = append(tests, testData{"Incorrect Entry Portfolio Not found", "/portfolio/8/entry", `{"coinName":"bitcoin", "amount": 10, "price": 10, "transactionFee": 10}`, 404})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", test.endpoint, bytes.NewBuffer([]byte(test.payload)))
			req.Header.Add("Content-Type", "application/json")
			if err != nil {
				fmt.Println(err)
			}
			resp := httptest.NewRecorder()
			testRouter.ServeHTTP(resp, req)
			log.Println(resp.Body.String())
			if resp.Code != test.expect {
				t.Fail()
			}
		})
	}

}
