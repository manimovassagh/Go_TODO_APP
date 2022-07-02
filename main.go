package main

import (
	"log"
	"time"

	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todos"
	port           string = ":9000"
)

//Just to satisfy the compiler

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		title     string        `bson:"title"`
		completed bool          `bson:"completed"`
		createdAt time.Time     `bson:"created_at"`
	}
	todo struct {
		ID        string    `json:"id"`
		title     string    `json:"title"`
		completed string    `json:"completed"`
		createdAt time.Time `json:"created_at"`
	}
)

func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}


func main()  {
	r:=chi.NewRouter()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
