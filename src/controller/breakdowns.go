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

	"github.com/gorilla/mux"
)

func CreateBreakdowns(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var breakdown models.Breakdowns
	if err := json.Unmarshal(body, &breakdown); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(breakdown); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := breakdown.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewBreakdowns(db)
	breakdown.ID, err = repository.Create(breakdown)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, breakdown)

}

func SearchBreakdownId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["breakdownsId"], 10, 32)
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

	repository := repositories.NewBreakdowns(db)
	breakdown, err := repository.SearchBreakdown(ID)
	if err != nil {
		response.Erro(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusOK, breakdown)
}

func SearchBreakdowns(w http.ResponseWriter, r *http.Request) {

}

func ToAlterBreakdowns(w http.ResponseWriter, r *http.Request) {

}

func DeleteBreakdowns(w http.ResponseWriter, r *http.Request) {

}
