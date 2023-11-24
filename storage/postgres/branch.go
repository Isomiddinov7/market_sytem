package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Isomiddinov7/exam/models"
	"github.com/google/uuid"
)

type branchRepo struct {
	db *sql.DB
}

func NewBranchRepo(db *sql.DB) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(req *models.CreateBranch) (*models.Branch, error) {
	var (
		branchId = uuid.New().String()
		query    = `
			INSERT INTO "branches"(
				"id",
				"name",
				"phone",
				"image_url",
				"work_start_hour",
				"work_end_hour",
				"address",
				"delivery_price",
				"active",
				"updated_at"
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())`
	)

	_, err := r.db.Exec(
		query,
		branchId,
		req.Name,
		req.Phone,
		req.ImageUrl,
		req.WorkStartHour,
		req.WorkEndHour,
		req.Address,
		req.DeliveryPrice,
		req.Active,
	)

	if err != nil {
		return nil, err
	}

	return r.GetByID(&models.BranchPrimaryKey{Id: branchId})
}

func (r *branchRepo) GetByID(req *models.BranchPrimaryKey) (*models.Branch, error) {
	var (
		query = `
			SELECT
			"id",
			"name",
			"phone",
			"image_url",
			"work_start_hour",
			"work_end_hour",
			"address",
			"delivery_price",
			"active",
			"created_at",
			"updated_at"
			FROM "branches"
			WHERE id = $1
		`
	)
	var (
		id            sql.NullString
		name          sql.NullString
		phone         sql.NullString
		address       sql.NullString
		image_url     sql.NullString
		workstarthour sql.NullString
		workendhour   sql.NullString
		deliveryprice sql.NullFloat64
		active        sql.NullBool
		updated_at    sql.NullString
		created_at    sql.NullString
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&id,
		&name,
		&phone,
		&image_url,
		&workstarthour,
		&workendhour,
		&address,
		&deliveryprice,
		&active,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:            id.String,
		Name:          name.String,
		Phone:         phone.String,
		ImageUrl:      image_url.String,
		Address:       address.String,
		WorkStartHour: workstarthour.String,
		WorkEndHour:   workendhour.String,
		DeliveryPrice: deliveryprice.Float64,
		Active:        active.Bool,
		CreatedAt:     created_at.String,
		UpdatedAt:     updated_at.String,
	}, nil
}

func (r *branchRepo) GetList(req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	var (
		resp   models.GetListBranchResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND name ILIKE" + " '%" + req.Search + "%'"
	}

	if len(req.Query) > 0 {
		where += req.Query
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name",
			"phone",
			"image_url",
			"work_start_hour",
			"work_end_hour",
			"address",
			"delivery_price",
			"active",
			"created_at",
			"updated_at"
		FROM "branches"
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			branch models.Branch
		)

		err = rows.Scan(
			&resp.Count,
			&branch.Id,
			&branch.Name,
			&branch.Phone,
			&branch.ImageUrl,
			&branch.WorkStartHour,
			&branch.WorkEndHour,
			&branch.Address,
			&branch.DeliveryPrice,
			&branch.Active,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Branches = append(resp.Branches, &branch)
	}

	return &resp, nil
}

func (r *branchRepo) Update(req *models.UpdateBranch) (int64, error) {

	query := `
		UPDATE product
			SET
				name = $2,
				phone = $3,
				image_url = $4,
				work_start_hour = $5,
				work_end_hour = $6,
				address = $7,
				delivery_price = $8,
				active = $9,
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.Phone,
		req.ImageUrl,
		req.WorkStartHour,
		req.WorkEndHour,
		req.Address,
		req.DeliveryPrice,
		req.Active,
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

func (r *branchRepo) Delete(req *models.BranchPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM branches WHERE id = $1", req.Id)
	return err
}

func (r *branchRepo) NoClose(req *models.BranchActive) (*models.GetListBranchResponse, error) {
	var resp models.GetListBranchResponse
	days := strings.Split(time.Now().Local().String(), " ")
	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"name",
			"phone",
			"image_url",
			"work_start_hour",
			"work_end_hour",
			"address",
			"delivery_price",
			"active",
			"created_at",
			"updated_at"
		FROM "branches"
		WHERE work_start_hour <= $1 and work_end_hour >= $2 and active = $3
	`

	rows, err := r.db.Query(query, days[1], days[1], req.Active)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			branch models.Branch
		)

		err = rows.Scan(
			&resp.Count,
			&branch.Id,
			&branch.Name,
			&branch.Phone,
			&branch.ImageUrl,
			&branch.WorkStartHour,
			&branch.WorkEndHour,
			&branch.Address,
			&branch.DeliveryPrice,
			&branch.Active,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.Branches = append(resp.Branches, &branch)
	}

	return &resp, nil
}
