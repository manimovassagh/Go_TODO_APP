package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

func homeHandler(w http.ResponseWriter,r *http.Request)  {
	err:=rnd.Template(w, http.StatusOK,[]string{"static/home.tpl"},nil)
	checkErr(err)
}

func fetchTodos(w http.ResponseWriter,r http.Request){
	todos:= []todoModel{}
	if err:=db.C(collectionName).Find(bson.M{}).All(&todos);err!=nil{
		rnd.JSON(w,http.StatusProcessing,renderer.M{
			"message":"Failed to fetch todos",
			"error":err,
		
		})
		return
	}

	todoList:=[]todo{}
	
	for_,t := range todos{
	todtodoList=append(todtodoList,todo{
		ID:t.ID.Hex(),
		Title:t.title,
		completed:t.completed,
		createdAt:t.createdAt
	})
	}


}

func main() { 
	//optional
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())
	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		log.Println("Listning to port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Listen:%s\n", err)

		}
	}()

	<-stopstopChan
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel (
	log.Println("server gracefully stopped")
	)

}

func todoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateTodo)
		r.Delete("/{id}", createTodo)
	})
	return rg
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
