package models

type Order struct {
	Id          int64    `json:"id"`
	TableNumber int64    `json:"table_number"`
	WorkerId    int64    `json:"worker_id"`
	ProductIds  []string `json:"product_ids"`
	Tax         int64    `json:"tax"`
	Discount    int64    `json:"discount"`
	TotalPrice  int64    `json:"total_price"`
	CreatedAt   string   `json:"created_at"`
}

type GetOrderResponse struct {
	Id         int64    `json:"id"`
	WorkerId   int64    `json:"worker_id"`
	WorkerName string   `json:"worker_name"`
	ProductIds []string `json:"product_ids"`
	Tax        int64    `json:"tax"`
	Discount   int64    `json:"discount"`
	TotalPrice int64    `json:"total_price"`
	CreatedAt  string   `json:"created_at"`
}

type OrderList struct {
	Orders     []Order `json:"orders"`
	WorkerName string  `json:"worker_name"`
}

type CreateOrder struct {
	WorkerId    string         `json:"worker_id"`
	Products    []ProductCheck `json:"products"`
	TableNumber int64          `json:"table_number"`
	Tax         int64          `json:"tax"`
	TotalPrice  int64          `json:"total_price"`
}

type UpdateOrder struct {
	Id         string `json:"id"`
	Tax        int64  `json:"tax"`
	Discount   int64  `json:"discount"`
	TotalPrice int64  `json:"total_price"`
}

// type CheckResponse struct {
// 	Check bool `json:"chack"`
// }
