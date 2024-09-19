package main

import (
	"go-crud/config"
	"go-crud/controller"
	_ "go-crud/docs" // Import your docs
	"go-crud/model"
	"go-crud/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Setup database connection
	db := config.SetupDB()
	if err := db.AutoMigrate(&model.Mahasiswa{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize service and controller
	mahasiswaService := service.NewMahasiswaService(db)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaService)

	// Setup router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/mahasiswa", mahasiswaController.CreateMahasiswa).Methods("POST")
	router.HandleFunc("/mahasiswa/{nim}", mahasiswaController.GetMahasiswa).Methods("GET")
	router.HandleFunc("/mahasiswa/{nim}", mahasiswaController.UpdateMahasiswa).Methods("PUT")
	router.HandleFunc("/mahasiswa/{nim}", mahasiswaController.DeleteMahasiswa).Methods("DELETE")

	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
