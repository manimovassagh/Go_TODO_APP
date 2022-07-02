package main

import (
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostname string = "localhost:27017"
)
