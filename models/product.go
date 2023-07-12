package models

var Products []Prod
var Providers []Prov

type Prod struct {
	Id          int    `csv:"id" json:"id"`
	Title       string `csv:"title" json:"title"`
	Description string `csv:"description" json:"description"`
	Price       int    `csv:"price" json:"price"`
	Brand       string `csv:"brand" json:"brand"`
	Category    string `csv:"category" json:"category"`
	ProviderId  int    `csv:"provider_id" json:"providerId"`
}
type Prov struct {
	Id        int    `csv:"id" json:"id"`
	Title     string `csv:"title" json:"title"`
	CreatedAt string `csv:"created_at" json:"createdAt"`
	Status    string `csv:"status" json:"status"`
}
