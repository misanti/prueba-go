package models

type Buyer struct {
	ID       int64
	Code     int64  `json:"code"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
