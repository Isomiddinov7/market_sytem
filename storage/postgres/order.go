package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Isomiddinov7/exam/models"
	"github.com/spf13/cast"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Id() int {
	resp, _ := r.GetList(&models.GetListOrderRequest{Offset: 0})
	return resp.Count
}

func (r *orderRepo) Create(req *models.CreateOrder) (*models.Order, error) {
	resp := r.Id()
	nol := "0"
	orderId := cast.ToString(resp + 1)
	var son string
	lenght := 7 - len(orderId)

	for lenght > 0 {
		son += nol
		lenght--
	}
	orderId = "O-" + son + orderId

	var (
		query = `
			INSERT INTO "orders"(
				"id",
				"client_id",
				"branch_id",
				"address",
				"delivery_price",
				"total_count",
				"total_price",
				"status",
				"updated_at"
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
			`
	)

	_, err := r.db.Exec(
		query,
		orderId,
		req.ClientID,
		req.BranchId,
		req.Address,
		req.DeliveryPrice,
		req.TotalCount,
		req.TotalPrice,
		req.Status,
	)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return r.GetByID(&models.OrderPrimaryKey{Id: orderId})
}

func (r *orderRepo) GetByID(req *models.OrderPrimaryKey) (*models.Order, error) {
	var (
		branchdata models.Branch
		clientdata models.Client
		query      = `
			SELECT
				o.id,
				o.client_id,
				c.first_name,
				c.last_name,
				c.phone,
				c.image_url,
				c.data_of_birth,
				c.updated_at,
				c.created_at,
				o.branch_id,
				b.name,
				b.phone,
				b.image_url,
				b.work_start_hour,
				b.work_end_hour,
				b.address,
				b.updated_at,
				b.created_at,
				o.address as o_address,
				o.delivery_price,
				o.total_count,
				o.total_price,
				o.status,
				o.updated_at,
				o.created_at
			FROM "orders" as o
			LEFT JOIN "clients" as c on c.id = o.client_id
			LEFT JOIN "branches" as b on b.id = o.branch_id
			WHERE o.id = $1
		`
	)
	var (
		id             sql.NullString
		client_id      sql.NullString
		branch_id      sql.NullString
		address        sql.NullString
		delivery_price sql.NullFloat64
		total_count    sql.NullInt64
		total_price    sql.NullFloat64
		status         sql.NullString
		updated_at     sql.NullString
		created_at     sql.NullString
	)

	err := r.db.QueryRow(query, req.Id).Scan(
		&id,
		&client_id,
		&clientdata.FirstName,
		&clientdata.LastName,
		&clientdata.Phone,
		&clientdata.Photos,
		&clientdata.Data_of_birth,
		&clientdata.UpdatedAt,
		&clientdata.CreatedAt,
		&branch_id,
		&branchdata.Name,
		&branchdata.Phone,
		&branchdata.ImageUrl,
		&branchdata.WorkStartHour,
		&branchdata.WorkEndHour,
		&branchdata.Address,
		&branchdata.UpdatedAt,
		&branchdata.CreatedAt,
		&address,
		&delivery_price,
		&total_count,
		&total_price,
		&status,
		&updated_at,
		&created_at,
	)
	clientdata.Id = client_id.String
	branchdata.Id = branch_id.String
	if err != nil {
		return nil, err
	}
	return &models.Order{
		Id:            id.String,
		ClientID:      client_id.String,
		Client:        clientdata,
		BranchId:      branch_id.String,
		Branch:        branchdata,
		Address:       address.String,
		DeliveryPrice: delivery_price.Float64,
		TotalCount:    total_count.Int64,
		TotalPrice:    total_price.Float64,
		Status:        status.String,
		UpdatedAt:     updated_at.String,
		CreatedAt:     created_at.String,
	}, nil
}

func (r *orderRepo) GetList(req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		clientdata = models.Client{}
		branchdata = models.Branch{}
		resp       models.GetListOrderResponse
		where      = " WHERE TRUE"
		offset     = " OFFSET 0"
		limit      = " LIMIT 10"
		sort       = " ORDER BY o.id DESC"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if len(req.Search) > 0 {
		where += " AND o.id ILIKE" + " '%" + req.Search + "%'" + " OR c.first_name ILIKE" + " '%" + req.Search + "%'" + " OR b.name ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
			SELECT
				COUNT(*) OVER(),
				o.id,
				o.client_id,
				c.first_name,
				c.last_name,
				c.phone,
				c.image_url,
				c.data_of_birth,
				c.updated_at,
				c.created_at,
				o.branch_id,
				b.name,
				b.phone,
				b.image_url,
				b.work_start_hour,
				b.work_end_hour,
				b.address,
				b.updated_at,
				b.created_at,
				o.address as o_address,
				o.delivery_price,
				o.total_count,
				o.total_price,
				o.status,
				o.updated_at,
				o.created_at
			FROM "orders" as o
			LEFT JOIN "clients" as c on c.id = o.client_id
			LEFT JOIN "branches" as b on b.id = o.branch_id
	`

	query += where + sort + offset + limit
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			id             sql.NullString
			client_id      sql.NullString
			branch_id      sql.NullString
			address        sql.NullString
			delivery_price sql.NullFloat64
			total_count    sql.NullInt64
			total_price    sql.NullFloat64
			status         sql.NullString
			updated_at     sql.NullString
			created_at     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&client_id,
			&clientdata.FirstName,
			&clientdata.LastName,
			&clientdata.Phone,
			&clientdata.Photos,
			&clientdata.Data_of_birth,
			&clientdata.UpdatedAt,
			&clientdata.CreatedAt,
			&branch_id,
			&branchdata.Name,
			&branchdata.Phone,
			&branchdata.ImageUrl,
			&branchdata.WorkStartHour,
			&branchdata.WorkEndHour,
			&branchdata.Address,
			&branchdata.UpdatedAt,
			&branchdata.CreatedAt,
			&address,
			&delivery_price,
			&total_count,
			&total_price,
			&status,
			&updated_at,
			&created_at,
		)
		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, &models.Order{
			Id:            id.String,
			ClientID:      client_id.String,
			Client:        clientdata,
			BranchId:      branch_id.String,
			Branch:        branchdata,
			Address:       address.String,
			DeliveryPrice: branchdata.DeliveryPrice,
			TotalCount:    total_count.Int64,
			TotalPrice:    total_price.Float64,
			Status:        status.String,
			UpdatedAt:     updated_at.String,
			CreatedAt:     created_at.String,
		})
	}
	return &resp, nil
}

func (r *orderRepo) Update(req *models.UpdateOrder) (int64, error) {

	query := `
		UPDATE "orders"
			SET
				client_id = $2,
				branch_id = $3,
				delivery_price = $4,
				total_count = $5,
				total_price = $6,
				status = $7
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.ClientID,
		req.BranchId,
		req.Address,
		req.TotalCount,
		req.TotalPrice,
		req.Status,
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

func (r *orderRepo) Delete(req *models.OrderPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", req.Id)
	return err
}
