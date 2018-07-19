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


	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	// nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`
}