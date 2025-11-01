package main

import (
	"flag"
	"go-crud/config"
	_ "go-crud/docs"
	"go-crud/internal/environment"
	"go-crud/internal/v1/controllers"
	"go-crud/internal/v1/models"
	"go-crud/internal/v1/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Parse command line flags for environment
	env := flag.String("env", "local", "Environment (local, dev, prod)")
	flag.Parse()

	// Load configuration
	if err := environment.LoadConfig(*env); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup database connection
	db := config.SetupDB()
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize service and controller
	studentService := services.NewStudentService(db)
	studentController := controllers.NewStudentController(studentService)

	// Setup router
	router := mux.NewRouter()

	// API v1 routes
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/student", studentController.CreateStudent).Methods("POST")
	v1.HandleFunc("/student/{studentId}", studentController.GetStudent).Methods("GET")
	v1.HandleFunc("/student/{studentId}", studentController.UpdateStudent).Methods("PUT")
	v1.HandleFunc("/student/{studentId}", studentController.DeleteStudent).Methods("DELETE")

	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	serverPort := ":" + environment.AppConfig.Server.Port
	log.Printf("Server starting on %s", serverPort)
	if err := http.ListenAndServe(serverPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
