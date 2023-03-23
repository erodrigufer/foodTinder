package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/erodrigufer/foodTinder/internal/data"
)

func retrieveMachineProducts(url string) (io.ReadCloser, error) {
	dataRequestDuration := time.Duration(10 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), dataRequestDuration)
	// Cancelling a context releases resources associated with it,
	// cancel should be call as soon as the operations running in a context
	// complete.
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		err = fmt.Errorf("unable to create a new GET request with timeout context: %w", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("unable to send HTTP request: %w", err)
		return nil, err
	}

	// Check if we got a 200 Code response, if not return error with
	// status code.
	if res.StatusCode != 200 {
		err = fmt.Errorf("HTTP request failed with status code %d (%s).\n", res.StatusCode, res.Status)
		return nil, err
	}

	return res.Body, nil

}

func decodeMachineProducts(response io.ReadCloser) ([]data.Product, error) {
	var input data.ProductsAPI

	err := rawJSON(response, &input)
	if err != nil {
		return nil, err
	}
	defer response.Close()

	output := make([]data.Product, 0, 20)
	for _, s := range input.Data.MachineProducts {
		var product data.Product
		product.ID = s.ID
		product.Name = s.Name
		output = append(output, product)
	}

	return output, nil
}

func (app *Application) fetchProducts() error {
	payload, err := retrieveMachineProducts("https://amperoid.tenants.foodji.io/machines/4bf115ee-303a-4089-a3ea-f6e7aae0ab94")
	if err != nil {
		return err
	}

	products, err := decodeMachineProducts(payload)
	if err != nil {
		return err
	}

	for _, product := range products {
		err := app.models.Products.Insert(&product)
		if err != nil {
			return err
		}
	}

	return nil

}
