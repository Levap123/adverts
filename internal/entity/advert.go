package entity

type Advert struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	Price  int    `json:"price,omitempty"`
	Status string `json:"status,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}
