package data

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type GoogleResponse struct {
	Queries Queries        `json:"queries"`
	Items   []SearchResult `json:"items"`
}
