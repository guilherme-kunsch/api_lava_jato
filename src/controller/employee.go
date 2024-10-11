package controller

import (
	"encoding/json"
	"io"
	"lavajato/src/banco"
	"lavajato/src/models"
	"lavajato/src/repositories"
	"lavajato/src/response"
	"net/http"
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

	if err := repositories.GetValidator().Struct(body); err != nil {
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

}

func ToAlterEmployee(w http.ResponseWriter, r *http.Request) {

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

}
