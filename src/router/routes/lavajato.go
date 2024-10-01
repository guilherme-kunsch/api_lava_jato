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
		URI:      "/v1/breakdowns/{breakdowsId}",
		Method:   http.MethodPut,
		Function: controller.SearchBreakdowns,
	},
	{
		URI:      "/v1/breakdowns/{breakdowsId}",
		Method:   http.MethodDelete,
		Function: controller.SearchBreakdowns,
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
		Method:   http.MethodPost,
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

}
