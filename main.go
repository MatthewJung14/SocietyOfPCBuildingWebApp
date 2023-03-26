package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"math/rand"

	"github.com/rs/cors"
)

var SECRET_KEY = []byte("teehee")

type User struct {
	gorm.Model
	FirstName string `json:"firstname" gorm:"firstname"`
	LastName  string `json:"lastname" gorm:"lastname"`
	Email     string `json:"email" gorm:"primaryKey" gorm:"uniqueIndex"`
	Password  string `json:"password" gorm:"password"`
}

// A silly little struct used to reuse db connections
type Env struct {
	db *gorm.DB
}

// Takes in password, returns a hash
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// Generates a JWT to be used for authorization purposes
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Sets the token expiration time to one day
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

// A middleware function to check that a JWT is legit
func ValidateJWT(next func(response http.ResponseWriter, request *http.Request)) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					response.WriteHeader(http.StatusUnauthorized)
					response.Write([]byte("Unauthorized"))
				}
				return SECRET_KEY, nil
			})

			if err != nil {
				response.WriteHeader(http.StatusUnauthorized)
				response.Write([]byte("Unauthorized" + err.Error()))
			}

			if token.Valid {
				next(response, request)
			}
		} else {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized"))
		}
	})
}

// This function registers a new user
func (env *Env) userRegister(response http.ResponseWriter, request *http.Request) {
	fmt.Println("TEST")
	response.Header().Set("Content-Type", "application/json")
	var user User
	var hold User //Just need an empty instance of a user struct
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	db := env.db
	//Check to see if there is already a user associated with the given email address
	if err := db.Where("Email = ?", user.Email).First(&hold).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Create(&user)
			response.Write([]byte(`User registered`))
			return
		} else {
			panic("something terrible has happened")
		}
	} else {
		response.Write([]byte(`Email is already in use`))
		return
	}
}

// This function logs a user in
func (env *Env) UserLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User = User{}
	var dbUser User = User{}
	db := env.db
	json.NewDecoder(request.Body).Decode(&user)

	dbUser.Email = user.Email

	//Check to see if the user exists
	if err := db.First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Write([]byte("No user with that email exists: " + err.Error()))
			return
		} else {
			panic("something terrible has happened")
		}
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
	response.Write([]byte(`{Successful}`))
}

// An api endpoint to delete a user from the database
func (env *Env) DeactivateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User = User{}
	var dbUser User = User{}
	db := env.db
	json.NewDecoder(request.Body).Decode(&user)

	dbUser.Email = user.Email

	//Check that the user actually exists
	if err := db.First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Write([]byte("No user with that email exists: " + err.Error()))
			return
		} else {
			panic("something terrible has happened")
		}
	}

	env.db.Exec("DELETE FROM Users WHERE email = ?", user.Email)

	//Delete the user whose email matches the one given in the DELETE request
	response.Write([]byte("User " + user.Email + " successfully deleted"))
}

// Password reset confirmation handler
func (env *Env) PasswordResetConfirm(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Get user email and verification code from request body
	var data struct {
		Email            string `json:"email"`
		VerificationCode int    `json:"code"`
		NewPassword      string `json:"new_password"`
	}
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"error":"Invalid request body"}`))
		return
	}

	// Check if user exists in database
	var dbUser User
	result := env.db.Where("Email = ?", data.Email).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"error":"User not found"}`))
		return
	}

	// Verify verification code
	if dbUser.VerificationCode != data.VerificationCode {
		response.WriteHeader(http.StatusUnauthorized)
		response.Write([]byte(`{"error":"Invalid verification code"}`))
		return
	}

	// Update user password
	dbUser.Password = getHash([]byte(data.NewPassword))
	dbUser.VerificationCode = 0
	env.db.Save(&dbUser)
}

// Password reset request handler
func (env *Env) PasswordResetRequest(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Get user email from request body
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"error":"Invalid request body"}`))
		return
	}

	// Check if user exists in database
	var dbUser User
	result := env.db.Where("Email = ?", user.Email).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"error":"User not found"}`))
		return
	}

	// Generate verification code and save to database
	code := rand.Intn(999999) + 100000
	dbUser.VerificationCode = code
	env.db.Save(&dbUser)

	// Send email with verification code to user
	err = sendVerificationEmail(user.Email, code)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Failed to send email"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message":"Verification code sent"}`))
}

// A function to update a user's credentials - does not update email address
func (env *Env) UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User = User{}
	var dbUser User = User{}
	db := env.db
	json.NewDecoder(request.Body).Decode(&user)

	dbUser.Email = user.Email

	//Check that the user actually exists
	if err := db.First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Write([]byte("No user with that email exists: " + err.Error()))
			return
		} else {
			panic("something terrible has happened")
		}
	}

	//Raw SQL >>> GORM
	db.Exec("UPDATE Users SET first_name = ?, last_name = ?, password = ? WHERE email = ?", user.FirstName, user.LastName, getHash([]byte(user.Password)), user.Email)
	response.Write([]byte(`{Successful}`))
}

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
