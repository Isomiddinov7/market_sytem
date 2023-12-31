package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Isomiddinov7/exam/models"
	"github.com/Isomiddinov7/exam/pkg/helpers"
)

func (c *Handler) Order(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.CreateOrder(w, r)
	case "GET":
		var values = r.URL.Query()
		if _, ok := values["id"]; ok {
			c.GetByIDOrder(w, r)
		} else {
			c.GetListOrder(w, r)
		}
	case "PUT":
		c.UpdateOrder(w, r)
	case "DELETE":
		c.DeleteOrder(w, r)
	}
}

func (c *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	var createOrder models.CreateOrder
	err := json.NewDecoder(r.Body).Decode(&createOrder)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if createOrder.ClientID != "" {
		if !helpers.IsValidUUID(createOrder.ClientID) {
			handleResponse(w, http.StatusBadRequest, "client id is not uuid")
			return
		}
	}

	if createOrder.BranchId != "" {
		if !helpers.IsValidUUID(createOrder.BranchId) {
			handleResponse(w, http.StatusBadRequest, "branch id is not uuid")
			return
		}
	}

	resp, err := c.storage.Order().Create(&createOrder)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, resp)
}

func (c *Handler) GetByIDOrder(w http.ResponseWriter, r *http.Request) {

	var id = r.URL.Query().Get("id")

	resp, err := c.storage.Order().GetByID(&models.OrderPrimaryKey{Id: id})
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

func (c *Handler) GetListOrder(w http.ResponseWriter, r *http.Request) {
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

	resp, err := c.storage.Order().GetList(&models.GetListOrderRequest{
		Offset: offset,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (c *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {

	var updateOrder models.UpdateOrder
	err := json.NewDecoder(r.Body).Decode(&updateOrder)
	if err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if updateOrder.ClientID != "" {
		if !helpers.IsValidUUID(updateOrder.ClientID) {
			handleResponse(w, http.StatusBadRequest, "client id is not uuid")
			return
		}
	}

	if updateOrder.BranchId != "" {
		if !helpers.IsValidUUID(updateOrder.BranchId) {
			handleResponse(w, http.StatusBadRequest, "branch id is not uuid")
			return
		}
	}

	rowsAffected, err := c.storage.Order().Update(&updateOrder)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(w, http.StatusBadRequest, "no rows affected")
		return
	}

	resp, err := c.storage.Order().GetByID(&models.OrderPrimaryKey{Id: updateOrder.Id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusAccepted, resp)
}

func (c *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")

	err := c.storage.Order().Delete(&models.OrderPrimaryKey{Id: id})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusNoContent, nil)
}
