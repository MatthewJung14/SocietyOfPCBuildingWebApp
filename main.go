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

// A simple little api endpoint that just exists for testing purposes
func test(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Test success\n")
	response.Write([]byte(`Test success`))
}

func main() {
	fmt.Println("BACKEND STARTING")
	//Open le database
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//SendEmail("c.tressler@ufl.edu", "Test", "This is a test email.")

	db.AutoMigrate(&User{})
	db.AutoMigrate(&ComputerEvent{})

	env := &Env{db}

	router := mux.NewRouter()

	router.HandleFunc("/api/signup", env.UserRegister).Methods("POST")
	router.HandleFunc("/api/login", env.UserLogin).Methods("POST")
	router.Handle("/api/test", ValidateJWT(test)).Methods("GET")
	router.Handle("/api/deactivate-account", ValidateJWT(env.DeactivateUser)).Methods("DELETE")
	router.Handle("/api/update-account", ValidateJWT(env.UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/reset-pass", env.PasswordResetRequest).Methods("PUT")
	router.HandleFunc("/api/reset-confirmation", env.PasswordResetConfirm).Methods("PUT")
	router.Handle("/api/admin-test", ValidateJWT(CheckAdminState(env.AdminTest))).Methods("GET")
	router.Handle("/api/change-admin-status", ValidateJWT(CheckAdminState(env.ChangeAdminState))).Methods("PUT")
	router.Handle("/api/create-event", ValidateJWT(env.CreateEvent)).Methods("POST")
	router.Handle("/api/update-event", ValidateJWT(env.UpdateEvent)).Methods("PUT")
	router.Handle("/api/get-event", ValidateJWT(env.GetEventAvailability)).Methods("GET")

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
