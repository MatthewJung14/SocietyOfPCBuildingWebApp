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

func TestUserLogin(t *testing.T) {
	// Define a mock environment for the function
	env := Env{db: &gorm.DB{}}

	// Define a mock user object with valid credentials
	user := User{Email: "test@mail.com", Password: "testpass"}
	body, _ := json.Marshal(user)

	// Create a new HTTP POST request to the login endpoint with the mock user object as the body
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the userLogin function with the mock environment, request, and response recorder
	env.userLogin(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	// Check the response body for a JWT token and success message

	expectedResponse := `{"response":"Successful"}`
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
	}
}

func TestUserRegister(t *testing.T) {
	// Create a new instance of Env struct with a mock database connection
	env := &Env{db: &gorm.DB{}}

	// Create a new HTTP request with a POST method and a JSON payload
	payload := strings.NewReader(`{"Email": "test@mail.com", "Password": "testpass"}`)
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
	if !strings.Contains(recorder.Body.String(), "Successful") {
		t.Errorf("Expected response body to contain 'Successful' but got '%s'", recorder.Body.String())
	}
}
