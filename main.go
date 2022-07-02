package main

import (
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
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
	
)
