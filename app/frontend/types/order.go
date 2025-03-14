package types

type Order struct {
	CreatedDate string `json:"CreatedDate"`
	OrderId     string `json:"OrderId"`
	Items       []Item `json:"Items"`
}
type Item struct {
	Picture     string  `json:"Picture"`
	ProductName string  `json:"ProductName"`
	Qty         uint32  `json:"Qty"`
	Cost        float32 `json:"Cost"`
}
