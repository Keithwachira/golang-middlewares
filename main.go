package main

import (

	"github.com/go-zoo/bone"
	"net/http"
	"middware-go/handlers"
	"middware-go/middlewares"
)

func main() {

	router := bone.New()


	router.Get("/api/jobs", middleware.Adapt(http.HandlerFunc(handlers.GetJobs), middleware.SecondMiddleware(),middleware.LogTime()))
	//run it on port 8080
	http.ListenAndServe("0.0.0.0:8080", router)


}
