package models

type Order struct {
	Id         int64  `json:"id"`
	WorkerId   int64  `json:"worker_id"`
	ProductId  int64  `json:"product_id"`
	Tax        int64  `json:"tax"`
	Discount   int64  `json:"discount"`
	TotalPrice int64  `json:"total_price"`
	CreatedAt  string `json:"created_at"`
}


type OrderList struct {
	Orders []Order `json:"orders"`
}

type CreateOrder struct {
	WorkerId   int64  `json:"worker_id"`
	ProductId  int64  `json:"product_id"`
	Tax        int64  `json:"tax"`
	Discount   int64  `json:"discount"`
	TotalPrice int64  `json:"total_price"`
}

type UpdateOrder struct {
	Id         int64  `json:"id"`
	Tax        int64  `json:"tax"`
	Discount   int64  `json:"discount"`
	TotalPrice int64  `json:"total_price"`
}

// type CheckResponse struct {
// 	Check bool `json:"chack"`
// }
