package main

import (
	// "context"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"

	// "git"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "log"
	// "time"
	"fmt"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"TEST/controllers"
	"net/http"
)

func main() {
	fmt.Println("starting the app...")

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://abhishek:abhishek@cluster0.24gdy.mongodb.net/Cluster0?retryWrites=true&w=majority"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // defer client.Disconnect(ctx)
	// // err = client.Ping(ctx, readpref.Primary())
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	// ------------------------------------------------------------

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s
}
