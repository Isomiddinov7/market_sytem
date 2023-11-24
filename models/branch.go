package models

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type BranchActive struct {
	Active bool `json:"active"`
}

type CreateBranch struct {
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	ImageUrl      string  `json:"image_url"`
	WorkStartHour string  `json:"work_start_hour"`
	WorkEndHour   string  `json:"work_end_hour"`
	Address       string  `json:"address"`
	DeliveryPrice float64 `json:"delivery_price"`
	Active        bool    `json:"active"`
	UpdatedAt     string  `json:"updated_at"`
}

type BranchData struct {
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	ImageUrl      string  `json:"image_url"`
	WorkStartHour string  `json:"work_start_hour"`
	WorkEndHour   string  `json:"work_end_hour"`
	Address       string  `json:"address"`
	DeliveryPrice float64 `json:"delivery_price"`
	Active        bool    `json:"active"`
}

type Branch struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	ImageUrl      string  `json:"image_url"`
	WorkStartHour string  `json:"work_start_hour"`
	WorkEndHour   string  `json:"work_end_hour"`
	Address       string  `json:"address"`
	DeliveryPrice float64 `json:"delivery_price"`
	Active        bool    `json:"active"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type UpdateBranch struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	ImageUrl      string  `json:"image_url"`
	WorkStartHour string  `json:"work_start_hour"`
	WorkEndHour   string  `json:"work_end_hour"`
	Address       string  `json:"address"`
	DeliveryPrice float64 `json:"delivery_price"`
	Active        bool    `json:"active"`
}

type GetListBranchRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListBranchResponse struct {
	Count    int       `json:"count"`
	Branches []*Branch `json:"branches"`
}
