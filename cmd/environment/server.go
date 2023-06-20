package environment

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	App        *fiber.App
	Config     *Config
	Collection *mongo.Collection
}

type Config struct {
	Host string
	Port string
}

func NewServer() *Server {
	app := fiber.New()
	app.Use(logger.New())

	config := &Config{}

	dbURI := getEnvVar("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("album_league").Collection("prospects")

	return &Server{
		App:        app,
		Config:     config,
		Collection: collection,
	}
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
