package main

import (
	"log"
	"os"

	"api-naco/config"
	"api-naco/db"
	"api-naco/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	config.Load()
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// 🔥 connect DB (ครั้งเดียว)
	db.ConnectDB()
	defer db.DB.Close()
	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	log.Println("DB_NAME:", os.Getenv("POSTGRES_DB"))

	app := fiber.New()
	// storage.InitMinio()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // หรือกำหนดเป็น "http://localhost:5173"
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	// test add api
	routers.SetupRoute(app)
	log.Fatal(app.Listen(":8000"))
}
