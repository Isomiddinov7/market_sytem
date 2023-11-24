package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Isomiddinov7/exam/models"
	"github.com/Isomiddinov7/exam/pkg/helpers"
)

func (c *Handler) Product(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateProduct(w, r)
	case "GET":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.GetByIDProduct(w, r)
		} else {
			c.GetListProduct(w, r)
		}
	case "PUT":
		c.UpdateProduct(w, r)
	case "DELETE":
		c.DeleteProduct(w, r)
	}
}

func (c *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var createProduct models.CreateProduct
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if createProduct.CategoryId != "" {
		if !helpers.IsValidUUID(createProduct.CategoryId) {
			handleResponse(w, http.StatusBadRequest, "category id is not uuid")
			return
		}
	}

	resp, err := c.storage.Product().Create(&createProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDProduct(w http.ResponseWriter, r *http.Request) {

	var id = r.URL.Query().Get("id")
	resp, err := c.storage.Product().GetByID(&models.ProductPrimaryKey{Id: id})
	if err == sql.ErrNoRows {
		handleResponse(w, http.StatusBadRequest, "no rows in result set")
		return
	}

	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) GetListProduct(w http.ResponseWriter, r *http.Request) {
	limit, err := getIntegerOrDefaultValue(r.URL.Query().Get("limit"), 10)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(r.URL.Query().Get("offset"), 0)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := r.URL.Query().Get("search")
	if err != nil {
		handleResponse(w, http.StatusBadRequest, "invalid query search")
		return
	}

	resp, err := c.storage.Product().GetList(&models.GetListProductRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var updateProduct models.UpdateProduct
	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if updateProduct.CategoryId != "" {
		if !helpers.IsValidUUID(updateProduct.CategoryId) {
			handleResponse(w, http.StatusBadRequest, "category id is not uuid")
			return
		}
	}

	rowsAffected, err := c.storage.Product().Update(&updateProduct)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Product().GetByID(&models.ProductPrimaryKey{Id: updateProduct.Id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusAccepted, resp)
}

func (c *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	err := c.storage.Product().Delete(&models.ProductPrimaryKey{Id: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusNoContent, nil)
}
