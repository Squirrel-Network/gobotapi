package types

type Invoice struct {
	Currency string `json:"currency"`
	Description string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Title string `json:"title"`
	TotalAmount int `json:"total_amount"`
}