package main

import (
	"fmt"
	"log"
	"net/http"

	"utronshop.io/controllers"
	"utronshop.io/models"

	"github.com/gernest/utron"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.User{}, &models.Product{}, &models.Order{})

	// Create Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(controllers.NewAdminCtl)
	app.AddController(controllers.NewUser)
	app.AddController(controllers.NewProduct)
	app.AddController(controllers.NewOrder)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
