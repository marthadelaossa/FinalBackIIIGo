package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/marthadelaossa/FinalBackIIIGo/cmd/server/router"
	"github.com/marthadelaossa/FinalBackIIIGo/pkg/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/go-sql-driver/mysql"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	// Recover from panic.
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.\
	goEnv := os.Getenv("GO_ENV")
	if goEnv != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Connect to the database.
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// Create a new Gin engine.
	router := gin.New()
	router.Use(gin.Recovery())
	// Add the logger middleware.
	router.Use(middleware.Logger())

	// Add the swagger handler.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the application.
	runApp(db, router)

	// Close the connection.
	defer db.Close()

}

func runApp(db *sql.DB, engine *gin.Engine) {
	router := routes.NewRouter(engine, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.MapRoutes()
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}

}

// connectDB connects to the database.
func connectDB() (*sql.DB, error) {
	// Get values from environment variables
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUsername == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil, fmt.Errorf("all environment variables must be defined for the database connection")
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error opening the database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error verifying the database connection: %v", err)
	}

	return db, nil
}
