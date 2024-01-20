package domain

type ErrorReport struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
