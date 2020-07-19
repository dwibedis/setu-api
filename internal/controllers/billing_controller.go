package controllers

import (
	"net/http"
)

func FetchBill(writer http.ResponseWriter, request *http.Request) {
	bill := service.NewBilling()
}
