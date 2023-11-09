package domain

type ResponseHttpSuccess struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}

type ResponseHttpError struct {
	StatusCode int      `json:"status_code"`
	Error      []string `json:"error"`
}
