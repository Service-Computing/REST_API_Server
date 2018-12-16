package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	_ "unsafe"

	"github.com/HeChX/REST_API_Server/database"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
)

var secret = "test"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	fmt.Println(r)
	err := r.ParseForm()

	var users map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &users)

	var username = users["username"].(string)
	var password = users["password"].(string)

	if err != nil && username != "" && password != "" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("register request format is invalid.\n"))
		return
	}

	user.Username = username
	user.Password = password

	myDB := database.GetDB()

	if myDB.CheckUserIsExist(user.Username) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User already exists.\n"))
		return
	}

	myDB.InsertUser(user.Username, user.Password)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Registered successfully.\n"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := r.ParseForm()

	var users map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &users)
	var username = users["username"].(string)
	var password = users["password"].(string)

	if err != nil && username != "" && password != "" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("register request format is invalid.\n"))
		return
	}

	user.Username = username
	user.Password = password

	myDB := database.GetDB()

	if !myDB.CheckUserIsExist(user.Username) {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("The user does not exists.\n"))
		return
	} else if !myDB.CheckPassword(user.Username, user.Password) {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("The password is invalid.\n"))
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(secret))
	fmt.Println(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w)
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request) bool {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	var isValid bool
	if err == nil {
		if token.Valid {
			isValid = true
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
			isValid = false
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
		isValid = false
	}
	return isValid
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func queryPeople(w http.ResponseWriter, r *http.Request) {
	// if !ValidateToken(w, r) {
	myDB := database.GetDB()
	vars := mux.Vars(r)
	id := vars["id"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, myDB.QueryPeople(id))
	// }

}

func queryPlanet(w http.ResponseWriter, r *http.Request) {
	if ValidateToken(w, r) {
		myDB := database.GetDB()
		vars := mux.Vars(r)
		id := vars["id"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, myDB.QueryPlanet(id))
	}
}

func qPlanet(w http.ResponseWriter, r *http.Request) {
	if ValidateToken(w, r) {
		myDB := database.GetDB()
		vars := mux.Vars(r)
		id := vars["id"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, myDB.QueryPlanet(id))
	}
}

func queryFilm(w http.ResponseWriter, r *http.Request) {
	if ValidateToken(w, r) {
		myDB := database.GetDB()
		vars := mux.Vars(r)
		id := vars["id"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, myDB.QueryFilm(id))
	}
}

func querySpecies(w http.ResponseWriter, r *http.Request) {
	if ValidateToken(w, r) {
		myDB := database.GetDB()
		vars := mux.Vars(r)
		id := vars["id"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, myDB.QuerySpecies(id))
	}
}

func queryStarship(w http.ResponseWriter, r *http.Request) {
	if ValidateToken(w, r) {
		myDB := database.GetDB()
		vars := mux.Vars(r)
		id := vars["id"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, myDB.QueryStarship(id))
	}
}

func queryVehicle(w http.ResponseWriter, r *http.Request) {
	if ValidateToken(w, r) {
		myDB := database.GetDB()
		vars := mux.Vars(r)
		id := vars["id"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, myDB.QueryVehicle(id))
	}
}
