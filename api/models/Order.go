package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	OrderId		int
	OrderName 	string
	CustomerName string
	CompanyName string
	OrderDate time.Time
	DeliveredAmount json.Number
	TotalAmount json.Number
}

func (o *Order) FindAllOrders(db *gorm.DB, options struct{
	Offset int
	SearchTerm string
	FromDate time.Time
	ToDate time.Time
}) (*[]Order, *int, error) {
	orders := []Order{}
	var count int

	// db.LogMode(true)

	countObj := db.Table("order_items").Select("order_items.order_id").Joins("LEFT JOIN orders ON order_items.order_id = orders.id")
	if options.SearchTerm != "" {
		countObj = countObj.Where("orders.order_name LIKE ?","%"+options.SearchTerm+"%").Or("order_items.product LIKE ?","%"+options.SearchTerm+"%")
	}
	if options.FromDate.IsZero() == false {
		countObj = countObj.Where("orders.created_at::date >= ?",options.FromDate)
	}
	if options.ToDate.IsZero() == false {
		countObj = countObj.Where("orders.created_at::date <= ?",options.ToDate)
	}
	countObj = countObj.Group("order_items.order_id")
	countObj.Count(&count)

	db.Exec("SET TIMEZONE='Australia/Melbourne'")
	orderObj := db.Table("orders").Select(
		"orders.id AS order_id, " +
		"orders.order_name AS order_name, " +
		"customers.name AS customer_name, " +
		"customer_companies.company_name AS company_name, " +
		"orders.created_at AS order_date, " +
		"SUM(order_items.price_per_unit * order_items.quantity) AS delivered_amount, " +
		"SUM(order_items.price_per_unit * order_items.quantity) AS total_amount")
	orderObj = orderObj.Joins("LEFT JOIN customers ON customers.user_id = orders.customer_id")
	orderObj = orderObj.Joins("LEFT JOIN customer_companies ON customer_companies.company_id = customers.company_id")
	orderObj = orderObj.Joins("LEFT JOIN order_items ON order_items.order_id = orders.id")
	if options.SearchTerm != "" {
		orderObj = orderObj.Where("orders.order_name LIKE ?","%"+options.SearchTerm+"%").Or("order_items.product LIKE ?","%"+options.SearchTerm+"%")
	}
	if options.FromDate.IsZero() == false {
		orderObj = orderObj.Where("orders.created_at::date >= ?",options.FromDate)
	}
	if options.ToDate.IsZero() == false {
		orderObj = orderObj.Where("orders.created_at::date <= ?",options.ToDate)
	}
	orderObj = orderObj.Group("orders.id, orders.order_name, customers.name, customer_companies.company_name, orders.created_at")
	orderObj = orderObj.Order("orders.id ASC")
	orderObj = orderObj.Limit(5)
	orderObj = orderObj.Offset( options.Offset)
	orderObj.Find(&orders)

	return &orders, &count, nil
}