package models

type Product struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductList struct {
	Products []Product `json:"products"`
	Count    int64     `json:"count"`
}

type CreateProduct struct {
	Title       string `json:"title"` 
	OwnerId     string `json:"owner_id"`
	Description string `json:"description"` 
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
	CategoryId  string `json:"category_id"`
}

type UpdateProduct struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Picture     string `json:"picture"`
}

type SearchProductRequest struct {
	Page    int64  `json:"page"`
	Limit   int64  `json:"limit"`
	Title   string `json:"title"`
	OwnerId string `json:"owner_id"`
}

type GetAllProductByCategoryIdRequest struct {
	CategoryId string `json:"category_id"`
	Page       int64  `json:"page"`
	Limit      int64  `json:"limit"`
}

type CheckResponse struct {
	Check bool `json:"chack"`
}
