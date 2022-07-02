package main

import (
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todos"
	port           string = ":9000"
)
