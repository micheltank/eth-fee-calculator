package presenter

type ApiError struct {
	Message string `json:"message"`
	Key     string `json:"key"`
	Detail  string `json:"detail,omitempty"`
}
