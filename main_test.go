package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/gorm"
)

func TestLoginHandler(t *testing.T) {
	// create a test HTTP server
	env := &Env{} // create a mock environment
	router := http.NewServeMux()
	router.HandleFunc("/user-login", env.userLogin)
	server := httptest.NewServer(router)
	defer server.Close()
	user := User{Email: "testpass2@mail.com", Password: "testpass2"}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", server.URL+"/user-login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", res.StatusCode)
	}

	// Check the response body
	var response struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}
	if response.Token == "" {
		t.Errorf("Expected non-empty token but got empty string")
	}
}

func TestUserRegister(t *testing.T) {
	// Create a new instance of Env struct with a mock database connection
	env := &Env{db: &gorm.DB{}}

	// Create a new HTTP request with a POST method and a JSON payload
	payload := strings.NewReader(`{"Email": "test2@mail.com", "Password": "test2pass"}`)
	req := httptest.NewRequest("POST", "/register", payload)
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the userRegister function with the test HTTP request and response
	env.userRegister(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	expectedBody := "User registered"
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected response body '%s' but got '%s'", expectedBody, recorder.Body.String())
	}
}
