package routes

import (
	"lavajato/src/controller"
	"net/http"
)

var lavajatoRouter = []Route{
	//avarias
	{
		URI:      "/v1/breakdowns",
		Method:   http.MethodPost,
		Function: controller.CreateBreakdowns,
	},
	{
		URI:      "/v1/breakdowns",
		Method:   http.MethodGet,
		Function: controller.SearchBreakdowns,
	},
	{
		URI:      "/v1/breakdowns/{breakdownsId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterBreakdowns,
	},
	{
		URI:      "/v1/breakdowns/{breakdownsId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteBreakdowns,
	},
	//cargo
	{
		URI:      "/v1/areaofactivity",
		Method:   http.MethodPost,
		Function: controller.CreateAreaOfActivity,
	},
	{
		URI:      "/v1/areaofactivity",
		Method:   http.MethodGet,
		Function: controller.SearchAreaOfActivity,
	},
	{
		URI:      "/v1/areaofactivity/{activityId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterAreaOfActivity,
	},
	{
		URI:      "/v1/areaofactivity/{activityId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteAreaOfActivity,
	},
	//cliente
	{
		URI:      "/v1/client",
		Method:   http.MethodPost,
		Function: controller.CreateClient,
	},
	{
		URI:      "/v1/client",
		Method:   http.MethodGet,
		Function: controller.SearchClient,
	},
	{
		URI:      "/v1/client/{clientId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterClient,
	},
	{
		URI:      "/v1/client/{clientId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteClient,
	},
	//funcionario
	{
		URI:      "/v1/employee",
		Method:   http.MethodPost,
		Function: controller.CreateEmployee,
	},
	{
		URI:      "/v1/employee",
		Method:   http.MethodGet,
		Function: controller.SearchEmployee,
	},
	{
		URI:      "/v1/employee/{employeeId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterEmployee,
	},
	{
		URI:      "/v1/employee/{employeeId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteEmployee,
	},
	//Ordens_de_serviço
	{
		URI:      "/v1/servicesorders",
		Method:   http.MethodPost,
		Function: controller.CreateServiceOrders,
	},
	{
		URI:      "/v1/servicesorders",
		Method:   http.MethodGet,
		Function: controller.SearchServiceOrders,
	},
	{
		URI:      "/v1/servicesorders/{serviceordersId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterServiceOrders,
	},
	{
		URI:      "/v1/servicesorders/{serviceordersId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteServiceOrders,
	},
	{
		URI:      "/v1/payment",
		Method:   http.MethodPost,
		Function: controller.CreatePayment,
	},
	{
		URI:      "/v1/payment",
		Method:   http.MethodGet,
		Function: controller.SearchClient,
	},
	{
		URI:      "/v1/payment/{paymentId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterPayment,
	},
	{
		URI:      "/v1/payment/{paymentId}",
		Method:   http.MethodDelete,
		Function: controller.DeletePayment,
	},
}
