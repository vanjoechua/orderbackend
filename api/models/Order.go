package models

import (
	"github.com/jinzhu/gorm"
	"math/big"
	"time"
)

type Order struct {
	OrderName 	string
	CustomerName string
	CompanyName string
	OrderDate time.Time
	DeliveredAmount big.Float
	TotalAmount big.Float
}

func (o *Order) FindAllOrders(db *gorm.DB, options struct{
	Offset int
	SearchTerm string
	FromDate time.Time
	ToDate time.Time
}) (*[]Order, *int, error) {
	var err error
	var orders []Order
	var count int
	var searchTermSql string

	if options.SearchTerm != "" {
		searchTermSql = "WHERE orders.order_name LIKE '%" + options.SearchTerm + "%' OR order_items.product LIKE '%" + options.SearchTerm + "%' "
	} else {
		searchTermSql = ""
	}

	if options.SearchTerm != "" {
		db.Table("orders").Select("count(orders.id)").Joins("LEFT JOIN order_items ON order_items.order_id = orders.id").Where("orders.order_name LIKE ?","%"+options.SearchTerm+"%").Or("order_items.product LIKE ?","%"+options.SearchTerm+"%").Group("order_items.order_id").Count(&count)
	} else {
		db.Table("orders").Select("count(id)").Count(&count)
	}
	db.Raw(
		"SELECT " +
			"orders.id AS order_id, " +
			"orders.order_name AS order_name, " +
			"customers.name AS customer_name, " +
			"customer_companies.company_name AS company_name, " +
			"orders.created_at AS order_date, " +
			"SUM(order_items.price_per_unit * order_items.quantity) AS DeliveredAmount, " +
			"SUM(order_items.price_per_unit * order_items.quantity) AS TotalAmount " +
			"FROM orders LEFT JOIN customers ON customers.user_id = orders.customer_id " +
			"LEFT JOIN customer_companies ON customer_companies.company_id = customers.company_id " +
			"LEFT JOIN order_items ON order_items.order_id = orders.id " +
			searchTermSql +
			"GROUP BY orders.id, orders.order_name, customers.name, customer_companies.company_name, orders.created_at " +
			"ORDER BY orders.id ASC " +
			"LIMIT 5 OFFSET ?", options.Offset).Scan(&orders)
	if err != nil {
		count = 0
		return &orders, &count, err
	}
	//return &orders, nil
	return &orders, &count, nil
}