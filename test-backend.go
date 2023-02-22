package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	// create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(LoginHandler))
	defer server.Close()

	// create a valid login request body
	form := url.Values{}
	form.Add("username", "testuser")
	form.Add("password", "testpassword")
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
	expected := "Welcome, testuser!"
	if !strings.Contains(expected, expected) {
		t.Errorf("Expected response body to contain '%v', but got '%v'", expected, res.Body)
	}
}
