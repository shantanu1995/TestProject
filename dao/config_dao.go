package dao

import (
	"log"

	. "github.com/shantanu1995/TestProject/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type ConfigDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "exercise"
)

// Establish a connection to database
func (m *ConfigDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of users
func (m *ConfigDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Find a user by its id
func (m *ConfigDAO) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a user into database
func (m *ConfigDAO) Insert(user User) error {

	

	count, _ := db.C(COLLECTION).Find(bson.M{ "username" : user.UserName}).Count()

	var err error

	if count > 0 {

		err := errors.New("Username already used")

	} else {

		err := db.C(COLLECTION).Insert(&user)

	}

	

	
	return err
}



// Delete an existing user
func (m *ConfigDAO) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing user
func (m *ConfigDAO) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
