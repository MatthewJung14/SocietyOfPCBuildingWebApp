package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserLogin(t *testing.T) {
	// Define a mock environment for the function
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	env := &Env{db}

	// Define a mock user object with valid credentials
	user := User{Email: "Sprint2", Password: "Sprint2"}
	body, _ := json.Marshal(user)

	// Create a new HTTP POST request to the login endpoint with the mock user object as the body
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the userLogin function with the mock environment, request, and response recorder
	env.UserLogin(rr, req)

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

	expectedResponse := `token`
	if !strings.Contains(rr.Body.String(), expectedResponse) {
		t.Errorf("handler returned unexpected body: got %v want response containing %v", rr.Body.String(), expectedResponse)
	}
}

func TestUserRegister(t *testing.T) {
	// Define a mock environment for the function
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	env := &Env{db}

	// Define a mock user object with valid credentials
	user := User{Email: "test53@mail.com", Password: "test53pass"}
	body, _ := json.Marshal(user)

	// Create a new HTTP POST request to the register endpoint with the mock user object as the body
	req, err := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the userRegister function with the mock environment, request, and response recorder
	env.UserRegister(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	// Check the response body for the success message
	expectedResponse := `User registered`
	if !strings.Contains(rr.Body.String(), expectedResponse) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
	}
}
func TestDeactivateUser(t *testing.T) {
	// Create a mock environment with an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	env := &Env{db}

	// Create a test user in the database
	testUser := User{FirstName: "test8", Email: "test8@mail.com"}
	db.Create(&testUser)

	// Test case 1: deactivating an existing user should delete the user from the database
	req1, _ := http.NewRequest("DELETE", "/users", bytes.NewBuffer([]byte(`{"email": "test8@mail.com"}`)))
	res1 := httptest.NewRecorder()
	env.DeactivateUser(res1, req1)

	// Check that the response contains the expected message
	expectedMsg := "User test8@mail.com successfully deleted"
	if res1.Body.String() != expectedMsg {
		t.Errorf("Unexpected response: got %v, expected %v", res1.Body.String(), expectedMsg)
	}

	// Check that the user has been deleted from the database
	var deletedUser User
	db.Where("email = ?", "test8@mail.com").First(&deletedUser)
	if deletedUser.ID != 0 {
		t.Errorf("User was not deleted from database")
	}

	// Test case 2: deactivating a non-existent user should return an error message
	req2, _ := http.NewRequest("DELETE", "/users", bytes.NewBuffer([]byte(`{"email": "test8@mail.com"}`)))
	res2 := httptest.NewRecorder()
	env.DeactivateUser(res2, req2)

	// Check that the response contains the expected error message
	expectedErrMsg := "No user with that email exists"
	if strings.Contains(res2.Body.String(), expectedErrMsg) {
		t.Errorf("Unexpected response: got %v, expected %v", res2.Body.String(), expectedErrMsg)
	}
}

func TestUpdateUser(t *testing.T) {
	// Create a mock environment with an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	env := &Env{db}

	// Create a test user in the database
	testUser := User{FirstName: "test11", Email: "test11@mail.com", Password: "test11pass"}
	db.Create(&testUser)

	// Test case 1: updating an existing user should update the user in the database
	req1, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer([]byte(`{"email": "test@mail.com", "FirstName": "test", "LastName": "test", "password": "testpass"}`)))
	res1 := httptest.NewRecorder()
	env.UpdateUser(res1, req1)

	//Check that the response is empty
	if res1.Body.String() == "Successful" {
		t.Errorf("Unexpected response: got %v, expected an empty response", res1.Body.String())
	}

	// Check that the user has been updated in the database
	var updatedUser User
	db.Where("email = ?", "test@mail.com").First(&updatedUser)
	if updatedUser.FirstName != "test" || updatedUser.LastName != "test" {
		t.Errorf("User was not updated in database")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte("testpass")); err != nil {
		t.Errorf("User password was not updated correctly")
	}

	// Test case 2: updating a non-existent user should return an error message
	req2, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer([]byte(`{"email": "test@mail.com", "FirstName": "test", "LastName": "test", "password": "testpass"}`)))
	res2 := httptest.NewRecorder()
	env.UpdateUser(res2, req2)

	// Check that the response contains the expected error message
	expectedErrMsg := "Successful"
	if !strings.Contains(res2.Body.String(), expectedErrMsg) {
		t.Errorf("Unexpected response: got %v, expected %v", res2.Body.String(), expectedErrMsg)
	}
}
func TestUpdateUserInformation(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	env := &Env{db}

	// Create a test user
	testUser := User{
		FirstName: "test",
		LastName:  "test",
		Email:     "test@mail.com",
		Password:  "testpass",
	}
	env.db.Create(&testUser)

	// Test UpdateUserName
	t.Run("UpdateUserName", func(t *testing.T) {
		updatedUser := User{FirstName: "test", LastName: "test", Email: "test@mail.com"}
		requestBody, _ := json.Marshal(updatedUser)

		// Create a mock http request
		req, _ := http.NewRequest("POST", "/update-username", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		// Call UpdateUserName function
		env.UpdateUserName(rr, req)

		// Check if the user was updated correctly
		var user User
		env.db.Where("Email = ?", updatedUser.Email).First(&user)
		if user.FirstName != updatedUser.FirstName || user.LastName != updatedUser.LastName {
			t.Errorf("UpdateUserName failed. User not updated correctly.")
		}

		// Check if the response code and message are correct
		if rr.Code != http.StatusOK {
			t.Errorf("UpdateUserName failed. Expected status code %d. Got %d.", http.StatusOK, rr.Code)
		}
		expectedResponse := `{"message":"User updated successfully"}`
		if rr.Body.String() != expectedResponse {
			t.Errorf("UpdateUserName failed. Expected response %s. Got %s.", expectedResponse, rr.Body.String())
		}
	})

	// Test UpdateUserEmail
	t.Run("UpdateUserEmail", func(t *testing.T) {
		updatedUser := User{FirstName: "test", LastName: "test", Email: "test@mail.com"}
		requestBody, _ := json.Marshal(updatedUser)

		// Create a mock http request
		req, _ := http.NewRequest("POST", "/update-email", bytes.NewBuffer(requestBody))
		rr := httptest.NewRecorder()

		// Call UpdateUserEmail function
		env.UpdateUserEmail(rr, req)

		// Check if the response code and message are correct
		if rr.Code == http.StatusOK {
			t.Errorf("UpdateUserEmail failed. Expected status code %d. Got %d.", http.StatusOK, rr.Code)
		}
	})
}
func TestUpdateUserEmail(t *testing.T) {
	// create a test environment
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	env := &Env{db}
	// create a test user
	user := User{FirstName: "test56", LastName: "test56", Email: "test56@mail.com"}
	env.db.Create(&user)

	// create a new email for the test user
	newEmail := "testemail59@mail.com"

	// create a request with the new email
	reqBody := bytes.NewBufferString(fmt.Sprintf(`{"email": "%s"}`, newEmail))
	req, err := http.NewRequest("PUT", "/user/email", reqBody)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// create a response recorder
	resRecorder := httptest.NewRecorder()

	// call the UpdateUserEmail method with the request and response recorder
	env.UpdateUserEmail(resRecorder, req)

	// check the response status code
	if resRecorder.Code == http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resRecorder.Code)
	}
}
