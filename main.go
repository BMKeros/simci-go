package main

import (
	"log"
	"simci-go/cmd/api/bootstrap"
)

func main() {

	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
	//// swagger
	//docs.SwaggerInfo.BasePath = "/api"
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//
	//r.Run(":5000")
}
