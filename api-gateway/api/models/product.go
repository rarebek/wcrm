package models

type Product struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

type CreateProduct struct {
	Title       string `json:"title" exmaple:"Lavash"`
	Description string `json:"description" example:"Juda mazzali"`
	Price       int64  `json:"price" example:"20000"`
	Discount    int64  `json:"discount" example:"12"`
	Picture     string `json:"picture" example:"http://static/images/myimage.jpg"`
	CategoryId  int64  `json:"category_id" example:"7"`
}

type UpdateProduct struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"category_id"`
}

type CheckResponse struct {
	Check bool `json:"chack"`
}
