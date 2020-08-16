package items

// Item main struct
type Item struct {
	ID          string  `json:"id"`
	Seller      int64   `json:"seller"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Status      string  `json:"status"`
}
