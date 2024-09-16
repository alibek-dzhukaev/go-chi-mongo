package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/alibek-dzhukaev/go-abs-beg/repository/mongodb"
	"github.com/alibek-dzhukaev/go-abs-beg/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error while reading .env")
	}

	slog.Info("env loaded successfully")
}

func mongoConnection() *mongo.Client {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverApi)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database(os.Getenv("MONGODB_DBNAME")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("pinged your deployment. u successfully connected to mongidb")

	return client
}

func main() {
	mongoClient := mongoConnection()
	defer mongoClient.Disconnect(context.Background())

	// mongo connection to collection
	collection := mongoClient.Database(os.Getenv("MONGO_DBNAME")).Collection(os.Getenv("MONGO_COLLECTION_NAME"))

	// userservice instance
	userService := usecase.UserService{
		DBClient: mongodb.MongoClient{
			Client: *collection,
		},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("server is healthy"))
		})

		r.Post("/users", userService.CreateUser)
		r.Get("/users/{id}", userService.GetUserById)
		r.Get("/users", userService.GetAllUsers)
		r.Put("/users/{id}", userService.UpdateUserAgeById)
		r.Delete("/users/{id}", userService.DeleteUserById)
		r.Delete("/users", userService.DeleteAllUsers)
	})

	http.ListenAndServe(":4444", r)
}
