package data

type SessionResponseData struct {
	SessionID string `json:"session_id,omitempty"`
}

type SessionResponse struct {
	APIVersion string              `json:"api_version"`
	Status     string              `json:"status"`
	Data       SessionResponseData `json:"data,omitempty"`
}

type Product struct {
	ID   string `json:"product_id"`
	Name string `json:"product_name"`
}

type Session struct {
	ID string `json:"session_id"`
}

type Vote struct {
	SessionID string `json:"session_id"`
	ProductID string `json:"product_id"`
	Vote      bool   `json:"vote"`
}

type SessionVotesResponse struct {
	APIVersion string           `json:"api_version"`
	Status     string           `json:"status"`
	Data       SessionVotesData `json:"data,omitempty"`
}

type SessionVotesData struct {
	Votes []Vote `json:"votes"`
}

type ProductsAPI struct {
	Status string          `json:"status"`
	Data   ProductsAPIData `json:"data"`
}

type ProductsAPIData struct {
	MachineProducts []MachineProduct `json:"machineProducts"`
}

type MachineProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AllProductsResponse struct {
	APIVersion string       `json:"api_version"`
	Status     string       `json:"status"`
	Data       ProductsData `json:"data,omitempty"`
}

type ProductsData struct {
	Products []Product `json:"products"`
}
