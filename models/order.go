package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type Order struct {
	Id            string      `json:"id"`
	ClientID      string      `json:"client_id"`
	Client        interface{} `json:"client"`
	BranchId      string      `json:"branch_id"`
	Branch        interface{} `json:"branch"`
	Address       string      `json:"address"`
	DeliveryPrice float64     `json:"delivery_price"`
	TotalCount    int64       `json:"total_count"`
	TotalPrice    float64     `json:"total_price"`
	Status        string      `json:"status"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

type CreateOrder struct {
	ClientID      string  `json:"client_id"`
	BranchId      string  `json:"branch_id"`
	Address       string  `json:"address"`
	DeliveryPrice float64 `json:"delivery_price"`
	TotalCount    int64     `json:"total_count"`
	TotalPrice    float64 `json:"total_price"`
	Status        string  `json:"status"`
}

type UpdateOrder struct {
	Id            string  `json:"id"`
	ClientID      string  `json:"client_id"`
	BranchId      string  `json:"branch_id"`
	Address       string  `json:"address"`
	DeliveryPrice float64 `json:"delivery_price"`
	TotalCount    int64     `json:"total_count"`
	TotalPrice    float64 `json:"total_price"`
	Status        string  `json:"status"`
}

type GetListOrderRequest struct {
	Offset int64    `json:"offset"`
	Limit  int64    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int     `json:"count"`
	Orders []*Order `json:"orders"`
}
