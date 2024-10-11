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

func CreateClient(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var client models.Client

	if err := json.Unmarshal(body, &client); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(client); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := client.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewClient(db)
	client.ID, err = repository.CreateClient(client)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, client)
}

func SearchClient(w http.ResponseWriter, r *http.Request) {
	client := strings.ToLower(r.URL.Query().Get("nome"))

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewClient(db)
	clients, err := repository.SearchClient(client)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, clients)
}

func SearchClientID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["clientId"], 10, 32)
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

	repository := repositories.NewClient(db)
	client, err := repository.SearchClientID(ID)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, client)
}

func ToAlterClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["clientId"], 10, 32)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	var client models.Client

	if err := json.Unmarshal(requestBody, &client); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(client); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewClient(db)
	if err = repository.UpdateClient(ID, client); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["clientId"], 10, 32)
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

	repository := repositories.NewClient(db)
	if err = repository.DeleteClient(ID); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
