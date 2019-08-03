package models

import "gopkg.in/mgo.v2/bson"


type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserName        string        `bson:"username" json:"username"`
	
}

type Exercise struct {
    ID          string `bson:"_id" json:"id"`
    UserName        string        `bson:"username" json:"username"`
	Description     string      `bson:"description" json:"description"`
	Duration        string      `bson:"duration" json:"duration"`
	Date            string      `bson:"date" json:"date"`

}

type ExerciseLog struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
    UserName        string        `bson:"username" json:"username"`
    Count          int          `bson:"count" json:"count"`
    Log            []TempExercise  `bson:"log" json:"log"`


}

type TempExercise struct {

	Description     string      `bson:"description" json:"description"`
	Duration        string      `bson:"duration" json:"duration"`
	Date            string      `bson:"date" json:"date"`


}
