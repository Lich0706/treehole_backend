package model

type ErrorRes struct {
	// Error message
	Error string `json:"error" example:"message"`
}

type GetAuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetSingleRecordResponse struct {
	// Get single record response
	Data interface{} `json:"data"`
}

type GetRecordsResponse struct {
	// Get multiple record response
	Data       []map[string]interface{} `json:"data"`
	Pagination map[string]interface{}   `json:"pagination,omitempty"`
}
