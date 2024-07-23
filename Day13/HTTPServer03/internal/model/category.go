package model

type Category struct {
	//use capital letter case
	CategoryId   int    `json:"id"`
	CategoryName string `json:"name"`
}

type Categories []Category
