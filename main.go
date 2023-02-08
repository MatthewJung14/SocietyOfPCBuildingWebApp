package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//Most of the code for this package is from here https://medium.com/@pkbhowmick007/user-registration-and-login-template-using-golang-mongodb-and-jwt-d85f09f1295e

var client *mongo.Client

var SECRET_KEY = []byte("gosecretkey")

type User struct {
	Name     string `json:"name" gorm:"name"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

// Takes in password, returns a hash
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// This does something???
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

// This function registers a new user
func userRegister(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Starting registration\n")
	response.Header().Set("Content-Type", "application/json")
	var user User
	fmt.Print("Decoding JSON\n")
	json.NewDecoder(request.Body).Decode(&user)
	fmt.Println("Hashing password")
	user.Password = getHash([]byte(user.Password))
	fmt.Println("Adding to database")
	fmt.Print("Huzzah")
}

// This function logs a user in
func userLogin(response http.ResponseWriter, request *http.Request) {
	http.Error(response, "Test message", 0)
	response.Header().Set("Content-Type", "application/json")
	var user User
	var dbUser User
	json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("GODB").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))

}

func test(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Test success\n")
}

func main() {
	fmt.Print("Starting\n")

	router := mux.NewRouter()
	router.HandleFunc("/login/register", userRegister).Methods("POST")
	router.HandleFunc("/login", userLogin).Methods("POST")
	router.HandleFunc("/test", test).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:4200", router))

	// err := http.ListenAndServe("localhost:4200", router)
	// if err != nil {
	//     log.Fatalln("There's an error with the server," err)
	// }
}
