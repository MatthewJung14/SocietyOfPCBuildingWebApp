package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	// create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(userLogin))
	defer server.Close()

	// create a valid login request body with username and password fields
	form := url.Values{}
	form.Add("test2", "Name")
	form.Add("test2@mail.com", "Email")
	form.Add("testpass2", "Password")
	body := strings.NewReader(form.Encode())

	// create a POST request with the login request body
	req, err := http.NewRequest("POST", server.URL+"/login", body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// perform the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// verify that the response status code is 200 OK
	if res.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %v", res.StatusCode)
	}

	// verify that the response body contains the expected message
	expected := "Successful"
	if !strings.Contains(expected, expected) {
		t.Errorf("Expected response body to contain '%v', but got '%v'", expected, res.Body)
	}
}

func TestUserRegister(t *testing.T) {
	// Create a request body with user data
	userData := User{
		FirstName: "test2",
		//might be an issue since im not adding a last name.
		Email:    "test2@mail.com",
		Password: "testpass2",
	}
	requestBody, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the request body
	request, err := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a test server with the router and handle the request
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userRegister)
	handler.ServeHTTP(rr, request)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "User registered"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
