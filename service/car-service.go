package service

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

var (
	carServiceUrl = ""
)

type fetchCarDataService struct {
}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", carServiceUrl)

	// Call the external API
	resp, _ := client.Get(carServiceUrl)

	// Write response to the channel
	carDataChannel <- resp
}
