package ethereal

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	router.
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", CreateTodo)
		v1.GET("/", FetchAllTodo)
		v1.GET("/:id", FetchSingleTodo)
		v1.PUT("/:id", UpdateTodo)
		v1.DELETE("/:id", DeleteTodo)
	}
	router.Run()
}

func envLoading() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//s3Bucket := os.Getenv("S3_BUCKET")
	//secretKey := os.Getenv("SECRET_KEY")
}
