package main

import (
	"IssueTrackerAPI/database"
	m "IssueTrackerAPI/models"
	"IssueTrackerAPI/routes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		port,
		dbName,
	)
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Employee{}, &m.Project{}, &m.ProjectEmployee{})
}

func init() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Static("/images", "./fileimg")
	routes.AuthRoutes(app)
	routes.ProjectRoutes(app)
	routes.ExcelRoutes(app)
	app.Listen(":3000")
}