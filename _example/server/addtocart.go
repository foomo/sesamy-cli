package server

type AddToCart struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Items    []*Item `json:"items,omitempty"`
}
