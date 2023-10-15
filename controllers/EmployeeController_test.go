package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test Create Employee Success (Positif Case), should return 201 code
func TestCreateEmployeeSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/employees", Create)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"first_name": "Jhon",
		"last_name":  "Doe",
		"email":      "jhon_doe@mail.com",
		"hire_date":  "01/02/2003",
	})

	req, err := http.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
	}
}

// Test Create Employee Failed (Negative Case), should return 422 code
// Invalid format date in hire_date payload request
func TestCreateEmployeeFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/employees", Create)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"first_name": "Jhon",
		"last_name":  "Doe",
		"email":      "jhon_doe@mail.com",
		"hire_date":  "01-02-2003",
	})

	req, err := http.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusUnprocessableEntity, w.Code)
	}
}

// Test Get All Employee Data Success (Positif Case), should return 200 code
func TestGetEmployeesSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/employees", GetEmployees)

	req, err := http.NewRequest(http.MethodGet, "/employees", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

// Test Find Employee Data Success (Positif Case), should return 200 code
func TestFindEmployeesSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/employees/1", GetEmployees)

	req, err := http.NewRequest(http.MethodGet, "/employees/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

// Test Find Employee Failed (Negative Case), should return 404 code
// Invalid employee ID request
func TestFindEmployeesFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/employees/:id", FindEmployee)

	req, err := http.NewRequest(http.MethodGet, "/employees/1000", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNotFound, w.Code)
	}
}

// Test Update Employee Success (Positif Case), should return 200 code
func TestUpdateEmployeeSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.PUT("/employees/:id", UpdateEmployee)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"first_name": "Jack",
		"last_name":  "Milner",
		"email":      "jack_milner@mail.com",
		"hire_date":  "03/02/2001",
	})

	req, err := http.NewRequest(http.MethodPut, "/employees/3", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

// Test Update Employee Failed (Negative Case), should return 404 code
// Invalid id employee in request
func TestUpdateEmployeeFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.PUT("/employees/:id", UpdateEmployee)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"first_name": "Jack",
		"last_name":  "Milner",
		"email":      "jack_milner@mail.com",
		"hire_date":  "03/02/2001",
	})

	req, err := http.NewRequest(http.MethodPut, "/employees/1000", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNotFound, w.Code)
	}
}

// Test Delete Employee Data Success (Positif Case), should return 200 code
func TestDeleteEmployeesSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/employees/:id", DeleteEmployee)

	req, err := http.NewRequest(http.MethodDelete, "/employees/3", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

// Test Delete Employee Failed (Negative Case), should return 404 code
// Invalid employee ID request
func TestDeleteEmployeesFailed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/employees/:id", DeleteEmployee)

	req, err := http.NewRequest(http.MethodDelete, "/employees/1000", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNotFound, w.Code)
	}
}
