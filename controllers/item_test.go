package controllers

import (
	"net/http"
	"testing"

	"../daos"
	"../services"
	"../testdata"
)

func TestItem(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeItemResource(&router.RouteGroup, services.NewItemService(daos.NewItemDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an item", "GET", "/items/1", "", http.StatusOK, `{"item_id":1, "promo_id":1, "name":"Belts", "stock":10, "price":20}`},
		{"t2 - get a nonexisting item", "GET", "/items/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an item", "POST", "/items", `{"promo_id": 1, "name":"Hoodie", "stock": 1, "price": 10}`, http.StatusOK, `{"item_id": 7, "promo_id": 1, "name":"Hoodie", "stock": 1, "price": 10}`},
		{"t4 - create an item with validation error", "POST", "/items", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t5 - update an item", "PUT", "/items/7", `{"stock": 2}`, http.StatusOK, `{"item_id":7, "promo_id":1, "name":"Hoodie", "stock":2, "price":10}`},
		{"t6 - update an item with validation error", "PUT", "/items/2", `{"name":""}`, http.StatusBadRequest, nameRequiredError},
		{"t7 - update a nonexisting item", "PUT", "/items/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an item", "DELETE", "/items/7", ``, http.StatusOK, `{"item_id":7, "promo_id":1, "name":"Hoodie", "stock":2, "price":10}`},
		{"t9 - delete a nonexisting item", "DELETE", "/items/99999", "", http.StatusNotFound, notFoundError},
	})
}