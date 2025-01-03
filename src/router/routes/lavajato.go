package routes

import (
	"lavajato/src/controller"
	"net/http"
)

var lavajatoRouter = []Route{
	//Breakdowns
	{
		URI:      "/v1/breakdowns",
		Method:   http.MethodPost,
		Function: controller.CreateBreakdowns,
	},
	{
		URI:      "/v1/breakdowns/{breakdownsId}",
		Method:   http.MethodGet,
		Function: controller.SearchBreakdownId,
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
	//Area Of Activity
	{
		URI:      "/v1/areaofactivity",
		Method:   http.MethodPost,
		Function: controller.CreateAreaOfActivity,
	},
	{
		URI:      "/v1/areaofactivity",
		Method:   http.MethodGet,
		Function: controller.SearchAreasOfActivity,
	},
	{
		URI:      "/v1/areaofactivity/{activityId}",
		Method:   http.MethodGet,
		Function: controller.SearchAreaOfActivityId,
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
	//Client
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
		Method:   http.MethodGet,
		Function: controller.SearchClientID,
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
	//Employee
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
		Method:   http.MethodGet,
		Function: controller.SearchEmployeeId,
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
	//Service Orders
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
		URI:      "/v1/serviceorders/{serviceordersId}",
		Method:   http.MethodGet,
		Function: controller.SearchServiceOrdersID,
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
	//payment
	{
		URI:      "/v1/payment",
		Method:   http.MethodPost,
		Function: controller.CreatePayment,
	},
	{
		URI:      "/v1/payment",
		Method:   http.MethodGet,
		Function: controller.SearchPayment,
	},
	{
		URI:      "/v1/payment/{paymentId}",
		Method:   http.MethodGet,
		Function: controller.SearchPaymentID,
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
	//services
	{
		URI:      "/v1/service",
		Method:   http.MethodPost,
		Function: controller.CreateService,
	},
	{
		URI:      "/v1/service/{serviceId}",
		Method:   http.MethodGet,
		Function: controller.SearchServiceID,
	},
	{
		URI:      "/v1/service/{serviceId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterService,
	},
	{
		URI:      "/v1/service/{serviceId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteService,
	},
	//Vehicle
	{
		URI:      "/v1/vehicle",
		Method:   http.MethodPost,
		Function: controller.CreateVehicle,
	},
	{
		URI:      "/v1/vehicle",
		Method:   http.MethodGet,
		Function: controller.SearchVehicle,
	},
	{
		URI:      "/v1/vehicle/{vehicleId}",
		Method:   http.MethodGet,
		Function: controller.SearchVehicleID,
	},
	{
		URI:      "/v1/vehicle/{vehicleId}",
		Method:   http.MethodPut,
		Function: controller.ToAlterVehicle,
	},
	{
		URI:      "/v1/vehicle/{vehicleId}",
		Method:   http.MethodDelete,
		Function: controller.DeleteVehicle,
	},
}
