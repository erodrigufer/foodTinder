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
	ID   string `json:"id"`
	Name string `json:"name"`
}
