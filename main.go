package main

import (
	"net/http"
)

func main()  {
	http.Handle("/api/v1/fetch-bill", controllers.FetchBill)
}
