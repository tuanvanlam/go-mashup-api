package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

var (
	ownerServiceUrl = ""
)

type fetchOwnerDataService struct {
}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", ownerServiceUrl)

	// Call the external API
	resp, _ := client.Get(ownerServiceUrl)

	// Write response to the channel
	ownerDataChannel <- resp
}
