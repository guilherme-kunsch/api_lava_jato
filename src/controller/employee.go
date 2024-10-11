package controller

import (
	"encoding/json"
	"io"
	"lavajato/src/banco"
	"lavajato/src/models"
	"lavajato/src/repositories"
	"lavajato/src/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var employee models.Employee

	if err := json.Unmarshal(body, &employee); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(employee); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := employee.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewEmployee(db)
	employee.ID, err = repository.CreateEmployee(employee)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, employee)
}

func SearchEmployee(w http.ResponseWriter, r *http.Request) {
	params := strings.ToLower(r.URL.Query().Get("funcionario"))

}

func SearchEmployeeId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["employeeId"], 10, 32)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewEmployee(db)
	employee, err := repository.SearchEmployeeId(ID)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, employee)
}

func ToAlterEmployee(w http.ResponseWriter, r *http.Request) {

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

}
