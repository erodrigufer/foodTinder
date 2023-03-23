package api

import (
	"fmt"
	"net/http"

	"github.com/erodrigufer/foodTinder/internal/data"
)

func (app *Application) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.models.Products.Products()
	if err != nil {
		// Internal Server Error.
		err = fmt.Errorf("error while retrieving list with all products: %w", err)
		app.serverError(w, err)
		return
	}

	resp := data.AllProductsResponse{
		APIVersion: API_VERSION,
		Status:     "success",
		Data: data.ProductsData{
			Products: products,
		},
	}

	err = writeJSON(w, http.StatusOK, resp)
	if err != nil {
		// Internal Server Error.
		err = fmt.Errorf("error while writing JSON response: %w", err)
		app.serverError(w, err)
		return

	}

}
