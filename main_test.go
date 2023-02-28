package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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

/*
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
*/
