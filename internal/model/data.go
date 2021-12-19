package model

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	Locked   bool   `json:"locked,omitempty"`
}

type Portfolio struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}

type Entry struct {
	CoinName       string  `json:"coinName"`
	Amount         int     `json:"amount"`
	Price          float64 `json:"price"`
	TransactionFee float64 `json:"transactionFee"`
}
