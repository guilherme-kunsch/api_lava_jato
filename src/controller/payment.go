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

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	var payment models.Payment

	if err := json.Unmarshal(body, &payment); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(payment); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewPayment(db)
	payment.ID, err = repository.CreatePayment(payment)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, payment)

}

func SearchPaymentID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["paymentId"], 10, 32)
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

	repository := repositories.NewPayment(db)
	payment, err := repository.SearchPaymentID(ID)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, payment)
}

func SearchPayment(w http.ResponseWriter, r *http.Request) {
	params := strings.ToLower(r.URL.Query().Get("funcionario"))

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewPayment(db)
	payment, err := repository.SearchPayment(params)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusOK, payment)
}

func ToAlterPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["paymentId"], 10, 32)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	var payment models.Payment

	if err := json.Unmarshal(requestBody, &payment); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repositories.GetValidator().Struct(payment); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conection()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewPayment(db)
	if err := repository.UpdatePayment(ID, payment); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletePayment(w http.ResponseWriter, r *http.Request) {

}
