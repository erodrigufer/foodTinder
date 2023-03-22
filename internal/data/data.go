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
	Name string `json:"name"`
}

type Session struct {
	ID string `json:"session_id"`
}

type Vote struct {
	SessionID string `json:"session_id"`
	ProductID string `json:"product_id"`
	Vote      bool   `json:"vote"`
}
