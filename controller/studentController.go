package controller

import (
	"encoding/json"
	"go-crud/internal/constparam"
	"go-crud/internal/logging"
	"go-crud/internal/utils"
	"go-crud/internal/v1/handler/response"
	"go-crud/model"
	"go-crud/service"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// StudentController defines the controller for Student operations.
type StudentController struct {
	service service.StudentService
}

// NewStudentController creates a new StudentController.
func NewStudentController(s service.StudentService) *StudentController {
	return &StudentController{s}
}

// CreateStudent handles POST /student
// @Summary Create a new student
// @Description Create a new student entry in the database
// @Accept json
// @Produce json
// @Param requestId header string false "requestId"
// @Param student body model.Student true "Student data"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /student [post]
func (c *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s model.Student
	var res response.Response
	var err error
	httpStatus := constparam.StatusCreated
	
	requestId := r.Header.Get("requestId")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	requestKey := "CREATE_STUDENT"
	
	err = json.NewDecoder(r.Body).Decode(&s)
	
	logging.LoggingDb(requestId, requestKey, "", constparam.HandlerRequest,
		constparam.MethodPost, utils.GetFuncName(), s,
		constparam.LogStudentCrud)
	
	if err != nil {
		httpStatus = constparam.StatusBadRequest
		res = response.HandleError(requestId, requestKey, constparam.StatusBadRequest, err)
	} else {
		err = c.service.Create(s)
		if err != nil {
			httpStatus = constparam.StatusInternalErr
			res = response.HandleError(requestId, requestKey, constparam.StatusInternalErr, err)
		} else {
			res = response.HandleSuccess(requestId, requestKey, constparam.StatusCreated, "Student created successfully", s)
		}
	}
	
	logging.LoggingDb(requestId, requestKey, strconv.Itoa(res.Code),
		constparam.HandlerResponse, "", utils.GetFuncName(), res,
		constparam.LogStudentCrud)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(res)
}

// GetStudent handles GET /student/{studentId}
// @Summary Get a student by Student ID
// @Description Retrieve a student from the database by Student ID
// @Accept json
// @Produce json
// @Param requestId header string false "requestId"
// @Param studentId path string true "Student ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /student/{studentId} [get]
func (c *StudentController) GetStudent(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var err error
	httpStatus := constparam.StatusOk
	
	requestId := r.Header.Get("requestId")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	requestKey := "GET_STUDENT|" + studentId
	
	logging.LoggingDb(requestId, requestKey, "", constparam.HandlerRequest,
		constparam.MethodGet, utils.GetFuncName(), studentId,
		constparam.LogStudentCrud)
	
	s, err := c.service.GetByStudentID(studentId)
	if err != nil {
		httpStatus = constparam.StatusNotFound
		res = response.HandleError(requestId, requestKey, constparam.StatusNotFound, err)
	} else {
		res = response.HandleSuccess(requestId, requestKey, constparam.StatusOk, "Student retrieved successfully", s)
	}
	
	logging.LoggingDb(requestId, requestKey, strconv.Itoa(res.Code),
		constparam.HandlerResponse, "", utils.GetFuncName(), res,
		constparam.LogStudentCrud)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(res)
}

// UpdateStudent handles PUT /student/{studentId}
// @Summary Update a student by Student ID
// @Description Update a student entry in the database
// @Accept json
// @Produce json
// @Param requestId header string false "requestId"
// @Param studentId path string true "Student ID"
// @Param student body model.Student true "Student data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /student/{studentId} [put]
func (c *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var s model.Student
	var res response.Response
	var err error
	httpStatus := constparam.StatusOk
	
	requestId := r.Header.Get("requestId")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	requestKey := "UPDATE_STUDENT|" + studentId
	
	err = json.NewDecoder(r.Body).Decode(&s)
	
	logging.LoggingDb(requestId, requestKey, "", constparam.HandlerRequest,
		constparam.MethodPut, utils.GetFuncName(), s,
		constparam.LogStudentCrud)
	
	if err != nil {
		httpStatus = constparam.StatusBadRequest
		res = response.HandleError(requestId, requestKey, constparam.StatusBadRequest, err)
	} else {
		s.StudentID = studentId
		err = c.service.Update(s)
		if err != nil {
			httpStatus = constparam.StatusNotFound
			res = response.HandleError(requestId, requestKey, constparam.StatusNotFound, err)
		} else {
			res = response.HandleSuccess(requestId, requestKey, constparam.StatusOk, "Student updated successfully", s)
		}
	}
	
	logging.LoggingDb(requestId, requestKey, strconv.Itoa(res.Code),
		constparam.HandlerResponse, "", utils.GetFuncName(), res,
		constparam.LogStudentCrud)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(res)
}

// DeleteStudent handles DELETE /student/{studentId}
// @Summary Delete a student by Student ID
// @Description Remove a student entry from the database by Student ID
// @Param requestId header string false "requestId"
// @Param studentId path string true "Student ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /student/{studentId} [delete]
func (c *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var err error
	httpStatus := constparam.StatusOk
	
	requestId := r.Header.Get("requestId")
	if requestId == "" {
		requestId = uuid.New().String()
	}
	
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	requestKey := "DELETE_STUDENT|" + studentId
	
	logging.LoggingDb(requestId, requestKey, "", constparam.HandlerRequest,
		constparam.MethodDelete, utils.GetFuncName(), studentId,
		constparam.LogStudentCrud)
	
	err = c.service.DeleteByStudentID(studentId)
	if err != nil {
		httpStatus = constparam.StatusNotFound
		res = response.HandleError(requestId, requestKey, constparam.StatusNotFound, err)
	} else {
		res = response.HandleSuccess(requestId, requestKey, constparam.StatusOk, "Student deleted successfully", nil)
	}
	
	logging.LoggingDb(requestId, requestKey, strconv.Itoa(res.Code),
		constparam.HandlerResponse, "", utils.GetFuncName(), res,
		constparam.LogStudentCrud)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(res)
}
