package models

type Seller struct {
	ID        int64
	Code      int64  `json:"code"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Followers []Buyer
}
