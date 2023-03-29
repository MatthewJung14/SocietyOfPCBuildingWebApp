package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
	"gorm.io/gorm"
)

var SECRET_KEY = []byte("teehee")

type User struct {
	gorm.Model
	FirstName        string `json:"firstname" gorm:"firstname"`
	LastName         string `json:"lastname" gorm:"lastname"`
	Email            string `json:"email" gorm:"primaryKey" gorm:"uniqueIndex"`
	Password         string `json:"password" gorm:"password"`
	VerificationCode int    `json:"verification_code"`
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
func (env *Env) UserRegister(response http.ResponseWriter, request *http.Request) {
	fmt.Println("REGISTER")
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
			fmt.Println("REGISTRATION SUCCESS")
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
	fmt.Println("LOGGING IN")
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
	fmt.Println("LOGIN SUCCESS")
}

// An api endpoint to delete a user from the database
func (env *Env) DeactivateUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("DEACTIVATING USER")
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
	fmt.Println("USER DELETED")
}

// Password reset confirmation handler
func (env *Env) PasswordResetConfirm(response http.ResponseWriter, request *http.Request) {
	fmt.Println("CONFIRMING PASSWORD RESET")
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
	result := env.db.Where("email = ?", data.Email).First(&dbUser)
	if result.RowsAffected == 0 {
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
	fmt.Println("PASSWORD RESET SUCCESS")
}

// Password reset request handler
func (env *Env) PasswordResetRequest(response http.ResponseWriter, request *http.Request) {
	fmt.Println("RECEIVED PASSWORD RESET REQUEST")
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
	result := env.db.Where("email = ?", user.Email).First(&dbUser)
	if result.RowsAffected == 0 {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"error":"User not found"}`))
		return
	}

	// Generate verification code and save to database
	code := rand.Intn(999999) + 100000
	dbUser.VerificationCode = code
	env.db.Save(&dbUser)

	// Send email with verification code to user
	SendEmail(user.Email, "SPCB: Verification Code", "Verification Code: "+strconv.Itoa(code)+"\n\n\nPlease do not respond to this email")

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message":"Verification code sent"}`))
	fmt.Println("VERIFICATION CODE SENT")
}

// A function to update a user's credentials - does not update email address
func (env *Env) UpdateUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("BEGINNING ACCOUNT UPDATE")
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
	fmt.Println("ACCOUNT UPDATED")
}
func (env *Env) UpdateUserName(response http.ResponseWriter, request *http.Request) {
	fmt.Println("UPDATING USERNAME")
	response.Header().Set("Content-Type", "application/json")

	// Decode request body into user object
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"error":"Invalid request body"}`))
		return
	}

	// Get user from database
	var dbUser User
	result := env.db.Where("Email = ?", user.Email).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"error":"User not found"}`))
		return
	}

	// Update user's first name and last name
	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	env.db.Save(&dbUser)

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message":"User updated successfully"}`))
	fmt.Println("USERNAME UPDATED")
}
func (env *Env) UpdateUserEmail(response http.ResponseWriter, request *http.Request) {
	fmt.Println("UPDATING EMAIL")
	response.Header().Set("Content-Type", "application/json")
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

	// Check if new email already exists in database
	result = env.db.Where("Email = ?", user.Email).First(&dbUser)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.WriteHeader(http.StatusConflict)
		response.Write([]byte(`{"error":"Email already exists"}`))
		return
	}

	// Update user email
	dbUser.Email = user.Email
	env.db.Save(&dbUser)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message":"Email updated successfully"}`))
	fmt.Println("EMAIL UPDATED")
}

func sendVerificationEmail(env *Env, to string, code int) error {
	//body := fmt.Sprintf("Your verification code is: %d", code)
	//subject := "Verify your email"
	//msg := gomail.NewMessage()
	//msg.SetHeader("From", env.email.From)
	//msg.SetHeader("To", to)
	//msg.SetHeader("Subject", subject)
	//msg.SetBody("text/plain", body)
	//if err := env.email.DialAndSend(msg); err != nil {
	//	return err
	//}
	return nil
}

func SendEmail(to string, subject string, body string) {
	fmt.Println("SENDING EMAIL")
	mail := gomail.NewMessage()
	//This is not secure at all - teehee
	var from string = "SocietyOfPCBuilders@outlook.com"
	var pass string = "iLCH44wYf5KdMqg"
	host := "smtp.office365.com"
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)
	a := gomail.NewDialer(host, 587, from, pass)
	if err := a.DialAndSend(mail); err != nil {
		panic(err)
	}
	fmt.Println("EMAIL SENT")
}
