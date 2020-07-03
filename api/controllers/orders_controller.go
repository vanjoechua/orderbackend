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
	//fromDate,present := query["fromdate"]
	//if !present {
	//	fromDate[0] = ""
	//}
	//toDate,present := query["fromdate"]
	//if !present {
	//	fromDate[0] = ""
	//}

	// layout := "2006-01-02T15:04:05.000Z"
	options.Offset, _ = strconv.Atoi(offset)
	options.SearchTerm = searchTerm
	// options.FromDate, _ = time.Parse(layout,fromDate[0])
	// options.ToDate, _ = time.Parse(layout,toDate[0])

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
