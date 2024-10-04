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

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
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

	if err := repositories.GetValidator().Struct(area); err != nil {
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

func SearchAreaOfActivityId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["activityId"], 10, 32)
	if err != nil {
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
	area, err := repository.SearchId(ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, area)
}

func SearchAreasOfActivity(w http.ResponseWriter, r *http.Request) {
	area := strings.ToLower(r.URL.Query().Get("cargo"))
	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewAreaOfActivity(db)
	areas, err := repository.Search(area)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, areas)
}

func ToAlterAreaOfActivity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	areaID, err := strconv.ParseUint(params["activityId"], 10, 32)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var area models.AreaOfActivity
	if err := json.Unmarshal(requestBody, &area); err != nil {
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
	if err = repository.Update(areaID, area); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteAreaOfActivity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	areaID, err := strconv.ParseUint(params["activityId"], 10, 32)
	if err != nil {
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
	if err = repository.Delete(areaID); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
