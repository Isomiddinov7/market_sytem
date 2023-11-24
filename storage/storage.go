package storage

import "github.com/Isomiddinov7/exam/models"

type StorageI interface {
	Category() CategoryRepoI
	Product() ProductRepoI
	Branch() BranchRepoI
	Client() ClientRepoI
	Order() OrderRepoI
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	GetByID(req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (int64, error)
	Delete(req *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	GetByID(req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(req *models.UpdateProduct) (int64, error)
	Delete(req *models.ProductPrimaryKey) error
}

type BranchRepoI interface {
	Create(req *models.CreateBranch) (*models.Branch, error)
	GetByID(req *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(req *models.UpdateBranch) (int64, error)
	Delete(req *models.BranchPrimaryKey) error
	NoClose(req *models.BranchActive) (*models.GetListBranchResponse, error)
}

type ClientRepoI interface {
	Create(req *models.CreateClient) (*models.Client, error)
	GetByID(req *models.ClientPrimaryKey) (*models.Client, error)
	GetList(req *models.GetListClientRequest) (*models.GetListClientResponse, error)
	Update(req *models.UpdateClient) (int64, error)
	Delete(req *models.ClientPrimaryKey) error
}

type OrderRepoI interface {
	Create(req *models.CreateOrder) (*models.Order, error)
	GetByID(req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(req *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(req *models.UpdateOrder) (int64, error)
	Delete(req *models.OrderPrimaryKey) error
}
