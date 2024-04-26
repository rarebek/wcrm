package models

type Product struct {
	Id          int64  `json:"id" bulding:"1234"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"categoryId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

type CreateProduct struct {
	Title       string `json:"title" description:"Lavash"`
	Description string `json:"description" description:"Juda mazzali"`
	Price       int64  `json:"price" description:"20000"`
	Discount    int64  `json:"discount" description:"12"`
	Picture     string `json:"picture" description:"http://static/images/myimage.jpg"`
	CategoryId  int64  `json:"categoryId" description:"7"`
}


type UpdateProduct struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"categoryId"`
}

type CheckResponse struct {
	Check bool `json:"chack"`
}
