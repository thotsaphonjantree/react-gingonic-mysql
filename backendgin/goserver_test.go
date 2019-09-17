package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

var id int

func TestGetAllMajor(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/major", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestGetAllStudent(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/student", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestInsertStudent(t *testing.T) {
	r := setupRouter()

	body := []byte(`{
		"student_code": "T000001",
		"first_name": "Testing",
		"last_name": "Fortest",
		"major_id": 1
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/student", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)


	// assert field of response body
	assert.Nil(t, err)
	id = int(resp["student_id"].(float64))
	assert.Equal(t, "T000001", resp["student_code"])
	assert.Equal(t, "Testing", resp["first_name"])
	assert.Equal(t, "Fortest", resp["last_name"])
	assert.Equal(t, float64(1), resp["major_id"])
}

func TestGetStudent(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/student/%d", id), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	// assert field of response body
	assert.Nil(t, err)
	assert.Equal(t, float64(id), resp["student_id"])
	assert.Equal(t, "T000001", resp["student_code"])
	assert.Equal(t, "Testing", resp["first_name"])
	assert.Equal(t, "Fortest", resp["last_name"])
	assert.Equal(t, float64(1), resp["major_id"])
}

func TestUpdateStudent(t *testing.T) {
	r := setupRouter()

	body := []byte(`{
		"student_code": "T000002",
		"first_name": "Testingtwo",
		"last_name": "Fortesttwo",
		"major_id": 2
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/student/%d", id), bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	// assert field of response body
	assert.Nil(t, err)
	assert.Equal(t, float64(id), resp["student_id"])
	assert.Equal(t, "T000002", resp["student_code"])
	assert.Equal(t, "Testingtwo", resp["first_name"])
	assert.Equal(t, "Fortesttwo", resp["last_name"])
	assert.Equal(t, float64(2), resp["major_id"])
}

func TestDeleteStudent(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/student/%d", id), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}




