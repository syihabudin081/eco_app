package utils

type PaginationMeta struct {
	CurrentPage int `json:"currentPage"`
	TotalPages  int `json:"totalPages"`
	TotalItems  int `json:"totalItems"`
}

type Response struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    interface{}     `json:"data"`           // Use an interface to accommodate different data types
	Meta    *PaginationMeta `json:"meta,omitempty"` // Pointer to avoid including it when nil
	Error   string          `json:"error,omitempty"`
}
