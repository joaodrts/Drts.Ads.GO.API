package entities

type Ad struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateStart   string `json:"date_start"`
	DateEnd     string `json:"date_end"`
	Status      string `json:"status"`
	Link        string `json:"link"`
}
