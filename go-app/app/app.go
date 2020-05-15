package app

import (
	"context"
	"go-app/configs"
	"go-app/repositories/bookrepo"
	"go-app/services/bookservice"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-app/controllers"
)

var (
	r = gin.Default()
)

// Run is the App Entry Point
func Run() {

	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoDB.URI) // use env variables
	// Connect to MongoDB
	mongoDB, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	/*
		====== Setup repositories =======
	*/
	bookRepo := bookrepo.NewBookRepo(mongoDB)
	/*
		====== Setup services ===========
	*/
	bookService := bookservice.NewBookService(bookRepo)
	/*
		====== Setup controllers ========
	*/
	bookCtl := controllers.NewBookController(bookService)

	/*
		======== Routes ============
	*/

	// API Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to your App on Docker",
		})
	})

	/*
		===== Book Routes =====
	*/
	r.POST("/books", bookCtl.PostBook)
	r.Run()
}
