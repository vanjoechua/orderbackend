package controllers

import (
	// "encoding/json"
	// "errors"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "strconv"

	// "github.com/gorilla/mux"
	"github.com/vanjoechua/orderbackend/api/models"
	"github.com/vanjoechua/orderbackend/api/responses"
	// "log"
	"strconv"

	"net/http"
	"time"
)

type Result struct {
	Orders *[]models.Order
	TotalOrders *int
}

type Options struct {
	Offset int
	SearchTerm string
	FromDate time.Time
	ToDate time.Time
}

func (server *Server) GetOrders(w http.ResponseWriter, r *http.Request) {

	var options Options
	// err := json.NewDecoder(r.Body).Decode(&options)
	query := r.URL.Query()
	offset := query.Get("offset")
	searchTerm := query.Get("searchterm")
	fromDate := query["fromdate"]
	toDate := query["todate"]

	layout := "2006-01-02"
	options.Offset, _ = strconv.Atoi(offset)
	options.SearchTerm = searchTerm
	fromDateStamp, _ := time.Parse(layout,fromDate[0])
	toDateStamp, _ := time.Parse(layout,toDate[0])
	options.FromDate = fromDateStamp
	options.ToDate = toDateStamp

	order := models.Order{}
	orders, count, err := order.FindAllOrders(server.DB, options)
	var result Result
	result.Orders = orders
	result.TotalOrders = count

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, result)
}
