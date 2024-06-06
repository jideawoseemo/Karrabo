package main

import (
	cmd "github.com/Tharolloo/go-playground/cmd/routes"
	"github.com/Tharolloo/go-playground/config"
	"log"

	"github.com/Tharolloo/go-playground/vin/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var app = config.NewAppTools()

func main() {
	srv := controller.NewVirtual(app)

	r := gin.Default()
	cmd.SetupRoutes(r, srv)
	logrus.Info("Server is running on port 8080")

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
