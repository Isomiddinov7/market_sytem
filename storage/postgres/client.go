package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Isomiddinov7/exam/models"
	"github.com/google/uuid"
)

type clientRepo struct {
	db *sql.DB
}

func NewClientRepo(db *sql.DB) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (c *clientRepo) Create(req *models.CreateClient) (*models.Client, error) {

	var (
		clientId = uuid.New().String()
		query    = `
			INSERT INTO "clients"(
				"id",
				"first_name", 
				"last_name",
				"phone",
				"image_url",
				"data_of_birth",
				"updated_at"
			) VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	)

	_, err := c.db.Exec(
		query,
		clientId,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Photos,
		req.Data_of_birth,
	)

	if err != nil {
		return nil, err
	}

	return c.GetByID(&models.ClientPrimaryKey{Id: clientId})
}

func (c *clientRepo) GetByID(req *models.ClientPrimaryKey) (*models.Client, error) {

	var (
		client models.Client
		query  = `
			SELECT
			  "id",
			  "first_name", 
			  "last_name",
			  "phone",
		      "image_url",
			  "data_of_birth",
			  "created_at",
			  "updated_at"	
			FROM "clients"
			WHERE "id" = $1
		`
	)

	err := c.db.QueryRow(query, req.Id).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Phone,
		&client.Photos,
		&client.Data_of_birth,
		&client.CreatedAt,
		&client.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *clientRepo) GetList(req *models.GetListClientRequest) (*models.GetListClientResponse, error) {
	var (
		resp   models.GetListClientResponse
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
		where += " AND first_name ILIKE" + " '%" + req.Search + "%'" + " OR last_name ILIKE" + " '%" + req.Search + "%'" + " OR phone ILIKE" + " '%" + req.Search + "%'"
	}

	if len(req.Query) > 0 {
		where += req.Query
	}

	var query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"first_name", 
			"last_name",
			"phone",
			"image_url",
			"data_of_birth",
			"created_at",
			"updated_at"
		FROM "clients"
	`

	query += where + sort + offset + limit
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			client models.Client
		)

		err = rows.Scan(
			&resp.Count,
			&client.Id,
			&client.FirstName,
			&client.LastName,
			&client.Phone,
			&client.Photos,
			&client.Data_of_birth,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Clients = append(resp.Clients, client)
	}

	return &resp, nil
}

func (r *clientRepo) Update(req *models.UpdateClient) (int64, error) {

	query := `
		UPDATE clients
			SET
				first_name = $2,
				last_name = $3,
				phone = $4,
				image_url = $5,
				data_of_birth = $6
		WHERE id = $1
	`
	result, err := r.db.Exec(
		query,
		req.Id,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Photos,
		req.Data_of_birth,
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

func (r *clientRepo) Delete(req *models.ClientPrimaryKey) error {
	_, err := r.db.Exec("DELETE FROM clients WHERE id = $1", req.Id)
	return err
}
