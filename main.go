package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
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
	r.Use(middleware.Logger)
	r.Get("/",homeHandler)
	r.Mount("/todo",todoHandlers())
	srv:=&http.Server{
		Addr: port,
		Handler: r,
		ReadTimeout: 60*time.Second,
		WriteTimeout: 60*time.Second,
		IdleTimeout: 60*time.Second,

	}
	go func ()  {
		log.Println("Listning to port",port)
		if err :=srv.ListenAndServe();err!=nil{
			log.Println("Listen:%s\n",err)

		}
	}()
}


func todoHandlers() http.Handler {
	rg:=chi.NewRouter()
rg.Group(func(r chi.Router) {
	r.Get("/",fetchTodos)
	r.Post("/",createTodo)
	r.Put("/{id}",updateTodo)
	r.Delete("/{id}",createTodo)
})
return rg
}


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
