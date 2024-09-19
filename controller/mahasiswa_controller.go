package controller

import (
	"encoding/json"
	"go-crud/model"
	"go-crud/service"
	"net/http"

	"github.com/gorilla/mux"
)

// MahasiswaController defines the controller for Mahasiswa operations.
type MahasiswaController struct {
	service service.MahasiswaService
}

// NewMahasiswaController creates a new MahasiswaController.
func NewMahasiswaController(s service.MahasiswaService) *MahasiswaController {
	return &MahasiswaController{s}
}

// CreateMahasiswa handles POST /mahasiswa
// @Summary Create a new mahasiswa
// @Description Create a new mahasiswa entry in the database
// @Accept json
// @Produce json
// @Param mahasiswa body model.Mahasiswa true "Mahasiswa data"
// @Success 201 {object} model.Mahasiswa
// @Router /mahasiswa [post]
func (c *MahasiswaController) CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var m model.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = c.service.Create(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

// GetMahasiswa handles GET /mahasiswa/{nim}
// @Summary Get a mahasiswa by NIM
// @Description Retrieve a mahasiswa from the database by NIM
// @Accept json
// @Produce json
// @Param nim path string true "NIM"
// @Success 200 {object} model.Mahasiswa
// @Failure 404 {string} string "Mahasiswa not found"
// @Router /mahasiswa/{nim} [get]
func (c *MahasiswaController) GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nim := vars["nim"]
	m, err := c.service.GetByNIM(nim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(m)
}

// UpdateMahasiswa handles PUT /mahasiswa/{nim}
// @Summary Update a mahasiswa by NIM
// @Description Update a mahasiswa entry in the database
// @Accept json
// @Produce json
// @Param nim path string true "NIM"
// @Param mahasiswa body model.Mahasiswa true "Mahasiswa data"
// @Success 200 {string} string "Mahasiswa updated"
// @Failure 404 {string} string "Mahasiswa not found"
// @Router /mahasiswa/{nim} [put]
func (c *MahasiswaController) UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nim := vars["nim"]
	var m model.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	m.NIM = nim
	err = c.service.Update(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mahasiswa updated"))
}

// DeleteMahasiswa handles DELETE /mahasiswa/{nim}
// @Summary Delete a mahasiswa by NIM
// @Description Remove a mahasiswa entry from the database by NIM
// @Param nim path string true "NIM"
// @Success 200 {string} string "Mahasiswa deleted"
// @Failure 404 {string} string "Mahasiswa not found"
// @Router /mahasiswa/{nim} [delete]
func (c *MahasiswaController) DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nim := vars["nim"]
	err := c.service.DeleteByNIM(nim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mahasiswa deleted"))
}
