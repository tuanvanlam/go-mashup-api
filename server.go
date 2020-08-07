package main

import (
	"go-mashup-api/controller"
	router "go-mashup-api/http"
	"go-mashup-api/service"
)

var (
	carDetailsService    = service.NewCarDetailsService()
	carDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           = router.NewChiRouter()
)

func main() {
	const port string = ":8080"
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)
	httpRouter.SERVE(port)
}
