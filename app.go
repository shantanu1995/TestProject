package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/shantanu1995/TestProject/config"
	. "github.com/shantanu1995/TestProject/dao"
	. "github.com/shantanu1995/TestProject/models"
)

var config = Config{}
var dao = ConfigDAO{}

// GET list of users
func AllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

// GET a user by its ID
func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var exerciselog ExerciseLog
	user, exercise ,count , err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	exerciselog.ID = user.ID
	exerciselog.UserName = user.UserName
	exerciselog.Count = count
	exerciselog.Log = exercise
	respondWithJson(w, http.StatusOK, exerciselog)
}

// POST a new user
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, map[string]string{"id" : user.ID.Hex() , "username" : user.UserName})
}

// PUT update an existing user
func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var exerlog Exercise
	if err := json.NewDecoder(r.Body).Decode(&exerlog); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if exerlog.Date == "" {
		exerlog.Date = time.Now().Format("01-02-2006")
	}
	exerlog1,err := dao.Update(exerlog);
	if  err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exerlog1.ID = exerlog.ID
	respondWithJson(w, http.StatusOK, exerlog1)
}

// DELETE an existing user
func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Database1 = config.Database1
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/exercise/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/api/exercise/new-user", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/api/exercise/add", UpdateUserEndPoint).Methods("POST")
	r.HandleFunc("/api/exercise/delete-user", DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/api/exercise/user/{id}", FindUserEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
