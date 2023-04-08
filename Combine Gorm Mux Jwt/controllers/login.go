package controllers

import (
	"Combine-Gorm-Mux-Jwt/users"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var APPLICATION_NAME = "JWT_ APP"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNATURE_KEY = []byte("jsakajjiehjknjksanhih")



type MyClaims struct {
	Username string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	jwt.StandardClaims
}


func CheckUsernameAndPassword(username, password string) (*users.User, error) {
	var user users.User
	err := users.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(w http.ResponseWriter, r *http.Request) {

	var userInput users.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "failed decode to JSON"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	
	user, err := CheckUsernameAndPassword(userInput.Username, userInput.Password)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Invalid username or password"})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(response)
		return
	}

	claims := &MyClaims {
		StandardClaims: jwt.StandardClaims{
			Issuer: APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},

		Username: user.Username,
		NamaLengkap: user.NamaLengkap,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	response := map[string]string{"token": signedToken}
	json.NewEncoder(w).Encode(response)
}




func Register(w http.ResponseWriter, r *http.Request) {
	// Mengambi input dari JSON
	var userInput users.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "failed decode to JSON"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	defer r.Body.Close()
	// Hash password menggunakan bycrypt
	hassPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hassPassword)

	// Insert ke database
	if err := users.DB.Create(&userInput).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Failed to save data"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(map[string]string{"Message": "Success"})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}