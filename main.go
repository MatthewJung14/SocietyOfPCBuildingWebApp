package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

// Most of the code for this package is from here https://medium.com/@pkbhowmick007/user-registration-and-login-template-using-golang-mongodb-and-jwt-d85f09f1295e

// A simple little api endpoint that just exists for testing purposes
func test(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Test success\n")
	response.Write([]byte(`Test success`))
}

func main() {
	//Open le database
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	env := &Env{db}

	router := mux.NewRouter()

	router.HandleFunc("/api/signup", env.userRegister).Methods("POST")
	router.HandleFunc("/api/login", env.UserLogin).Methods("POST")
	router.Handle("/api/test", ValidateJWT(test)).Methods("GET")
	router.Handle("/api/deactivate-account", ValidateJWT(env.DeactivateUser)).Methods("DELETE")
	router.Handle("/api/update-account", ValidateJWT(env.UpdateUser)).Methods("PUT")

	//This does something important I think
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("localhost:4200", handler))

	err = http.ListenAndServe("localhost:4200", handler)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}
