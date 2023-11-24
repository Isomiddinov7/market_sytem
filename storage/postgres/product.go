package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Isomiddinov7/exam/models"
	"github.com/Isomiddinov7/exam/pkg/helpers"
	"github.com/spf13/cast"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Id() int {
	resp, _ := r.GetList(&models.GetListProductRequest{Offset: 0})
	return resp.Count
}

func (r *productRepo) Create(req *models.CreateProduct) (*models.Product, error) {
	resp := r.Id()
	nol := "0"
	productId := cast.ToString(resp + 1)
	var son string
	lenght := 7 - len(productId)

	for lenght > 0 {
		son += nol
		lenght--
	}
	productId = "P-" + son + productId
	var (
		query = `
			INSERT INTO "product"(
				"id",
				"title",
				"description",
				"price",
				"image_url",
				"category_id",
				"updated_at"
			) VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	)

	_, err := r.db.Exec(
		query,
		productId,
		req.Title,
		req.Description,
		req.Price,
		req.ImageUrl,
		helpers.NewNullString(req.CategoryId),
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.ProductPrimaryKey{Id: productId})
}

func (r *productRepo) GetByID(req *models.ProductPrimaryKey) (*models.Product, error) {
	var categoryName models.CategoryData
	var (
		query = `
			SELECT
				p.id,
				p.title,
				p.category_id,
				c.title,
				c.image_url,
				c.parent_id,
				p.description,
				p.price,
				p.image_url,
				p.updated_at,
				p.created_at
			FROM "product" as p
			JOIN "category" as c on c.id = p.category_id
			WHERE p.id = $1
		`
	)
	var (
		id          sql.NullString
		title       sql.NullString
		description sql.NullString
		price       sql.NullFloat64
		image_url   sql.NullString
		category_id sql.NullString
		updated_at  sql.NullString
		created_at  sql.NullString
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&id,
		&title,
		&category_id,
		&categoryName.Title,
		&categoryName.ImageUrl,
		&categoryName.ParentID,
		&description,
		&price,
		&image_url,
		&updated_at,
		&created_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:          id.String,
		Title:       title.String,
		Description: description.String,
		Price:       price.Float64,
		ImageUrl:    image_url.String,
		CategoryId:  category_id.String,
		Category:    categoryName,
		UpdatedAt:   updated_at.String,
		CreatedAt:   created_at.String,
	}, nil
}

func (r *productRepo) GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	var (
		categoryName = models.CategoryData{}
		resp         models.GetListProductResponse
		where        = " WHERE TRUE"
		offset       = " OFFSET 0"
		limit        = " LIMIT 10"
		sort         = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND p.title ILIKE" + " '%" + req.Search + "%'"+ " OR c.title ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
	SELECT
		COUNT(*) OVER(),
		p.id,
		p.title,
		p.description,
		p.price,
		p.image_url,
		p.category_id,
		c.title,
		c.image_url,
		COALESCE(cast(c.parent_id as varchar), '') as parent_id,
		p.updated_at,
		p.created_at
	FROM product as p
	JOIN "category" as c on c.id = p.category_id
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			title       sql.NullString
			description sql.NullString
			price       sql.NullFloat64
			image_url   sql.NullString
			category_id sql.NullString
			updated_at  sql.NullString
			created_at  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&title,
			&description,
			&price,
			&image_url,
			&category_id,
			&categoryName.Title,
			&categoryName.ImageUrl,
			&categoryName.ParentID,
			&updated_at,
			&created_at,
		)
		if err != nil {
			return nil, err
		}
		resp.Products = append(resp.Products, &models.Product{
			Id:          id.String,
			Title:       title.String,
			CategoryId:  category_id.String,
			Category:    categoryName,
			Description: description.String,
			Price:       price.Float64,
			ImageUrl:    image_url.String,
			UpdatedAt:   updated_at.String,
			CreatedAt:   created_at.String,
		})
	}

	return &resp, nil
}

func (r *productRepo) Update(req *models.UpdateProduct) (int64, error) {

	query := `
		UPDATE product
			SET
				title = $2,
				description = $3,
				price = $4,
				image_url = $5,
				category_id = $6,
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Title,
		req.Description,
		req.Price,
		req.ImageUrl,
		helpers.NewNullString(req.CategoryId),
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *productRepo) Delete(req *models.ProductPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM product WHERE id = $1", req.Id)
	return err
}
