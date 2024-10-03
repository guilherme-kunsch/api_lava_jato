package controller

import (
	"encoding/json"
	"io"
	"lavajato/src/banco"
	"lavajato/src/models"
	"lavajato/src/repositories"
	"lavajato/src/response"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New() // Inicializa o validador
}

func CreateAreaOfActivity(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var area models.AreaOfActivity
	if err := json.Unmarshal(body, &area); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := validate.Struct(area); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := area.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return

	}

	defer db.Close()

	repository := repositories.NewAreaOfActivity(db)
	area.ID, err = repository.Create(area)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, area)
}

func SearchAreaOfActivity(w http.ResponseWriter, r *http.Request) {

}

func ToAlterAreaOfActivity(w http.ResponseWriter, r *http.Request) {

}

func DeleteAreaOfActivity(w http.ResponseWriter, r *http.Request) {

}
