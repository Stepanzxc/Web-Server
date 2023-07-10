package models

var Products []Prod
var Providers []Prov

type Prod struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	ProviderId  int    `json:"providerId"`
}
type Prov struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdTime"`
	Status    string `json:"status"`
}
