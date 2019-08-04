package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"
	"strconv"

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

func remove(s []Exercise, i int) []Exercise {
    return append(s[:i], s[i+1:]...)
}


// GET a user by its ID
func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query()

	id := ids.Get("id")
	from := ids.Get("from")
	to := ids.Get("to")
	limit := ids.Get("limit")

	var limit1 int
	limit1 = 0

	if limit != "" {

	limit1, _= strconv.Atoi(limit)


	}


	fmt.Printf(limit)
	fmt.Printf("from:-" + from)
	fmt.Printf("to :-" + to)
	var exerciselog ExerciseLog
	var tempexerciselog []TempExercise
	user, exercise ,count , err := dao.FindById(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	if from != "" {

		for i := 0; i < len(exercise); i++ { 

		dateStamp, _ := time.Parse("2006-01-02", exercise[i].Date); 
		dateStamp2, _ := time.Parse("2006-01-02", from); 

		fmt.Printf("Value is:-" + exercise[i].Date)
		fmt.Printf("Duration is:-" + exercise[i].Duration)
		println(i)
		fmt.Printf("dateStamp is:-" + dateStamp.String())
		fmt.Printf("dateStamp2 is:-" + dateStamp2.String())
		fmt.Printf("In range:-" + strconv.FormatBool(dateStamp.After(dateStamp2.AddDate(0, 0, -1))))
		
		if (!dateStamp.After(dateStamp2.AddDate(0, 0, -1))) {

		

			fmt.Printf("Value is:-" + exercise[i].Date)

			exercise = remove(exercise,i)
			i=i-1

		}

	}



	}

	if to != "" {

		for i := 0; i < len(exercise); i++ { 

		dateStamp, _ := time.Parse("2006-01-02", exercise[i].Date); 
		dateStamp2, _ := time.Parse("2006-01-02", to); 
		if (!dateStamp.Before(dateStamp2.AddDate(0, 0, +1))) {

		

			exercise = remove(exercise,i)
			i=i-1

		}

	}



	}

	if len(exercise) != count {
		count = len(exercise)
	}
	

	b , _ := json.Marshal(exercise)
	json.Unmarshal([]byte(string(b)), &tempexerciselog)
	exerciselog.ID = user.ID
	exerciselog.UserName = user.UserName
	if limit1 > 0 {
	exerciselog.Count = limit1
	exerciselog.Log = tempexerciselog[:limit1]
	} else{

		exerciselog.Count = count
	exerciselog.Log = tempexerciselog



	}
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
		exerlog.Date = time.Now().Format("2006-02-01")
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
	header := w.Header()
    header.Add("Access-Control-Allow-Origin", "http://localhost:3001")
    header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
    header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
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
	r.HandleFunc("/api/exercise/log", FindUserEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
