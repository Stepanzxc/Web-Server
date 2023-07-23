package models

type Category struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
type ProductWithProvider struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Brand       string    `json:"brand"`
	Category    Category  `json:"category"`
	Provider    Providers `json:"provider"`
}
type Providers struct {
	Id        int    `csv:"id" json:"id"`
	Title     string `csv:"title" json:"title"`
	CreatedAt string `csv:"created_at" json:"createdAt"`
	Status    string `csv:"status" json:"status"`
}
