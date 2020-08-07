package controller

import (
	"encoding/json"
	"go-mashup-api/service"
	"net/http"
)

type controller struct {
}

var (
	carDetailsService service.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(response http.ResponseWriter, request *http.Request)
}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(result)
}
