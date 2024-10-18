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

func CreateServiceOrders(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	var serviceOrders models.ServiceOrder

	if err := json.Unmarshal(body, &serviceOrders); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(serviceOrders); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewServiceOrders(db)
	serviceOrders.ID, err = repository.CreateServiceOrders(serviceOrders)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, serviceOrders)
}

func SearchServiceOrders(w http.ResponseWriter, r *http.Request) {

}

func ToAlterServiceOrders(w http.ResponseWriter, r *http.Request) {

}

func DeleteServiceOrders(w http.ResponseWriter, r *http.Request) {

}
