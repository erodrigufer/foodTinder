package main

import (
	"github.com/erodrigufer/foodTinder/internal/api"
)

func main() {

	port := 8000

	// dsn := "postgres://ft:foodTinder@localhost/ft"
	dsn := "host=localhost port=5432 user=ft password=foodTinder dbname=ft sslmode=disable"

	app := api.NewApplication(port, dsn)
	defer app.DB.Close()

	app.InfoLog.Printf("Starting server at port %d", port)

	err := app.Srv.ListenAndServe()
	app.ErrorLog.Fatal(err)

}
