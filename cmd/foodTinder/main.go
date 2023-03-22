package main

import (
	"github.com/erodrigufer/foodTinder/internal/api"
)

func main() {

	port := 8000

	app := api.NewApplication(port)

	app.InfoLog.Printf("Starting server at port %d", port)

	err := app.Srv.ListenAndServe()
	app.ErrorLog.Fatal(err)

}
