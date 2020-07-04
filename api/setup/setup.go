package setup

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"os"
)

func PopulateDb(db *gorm.DB) error {

	// _,err := CopyFile("/tmp/customers.csv","customers.csv")
	err := File("customers.csv","/tmp/customers.csv")
	if err != nil {
		fmt.Println(err)
	}
	customersTmpfile := "/tmp/customers.csv"
	db.Exec("TRUNCATE TABLE customers")
	db.Exec("COPY customers FROM '" + customersTmpfile + "' WITH (format csv, header)")

	err2 := File("customer_companies.csv","/tmp/customer_companies.csv")
	if err2 != nil {
		fmt.Println(err2)
	}
	db.Exec("TRUNCATE TABLE customer_companies")
	customerCompaniesTmpfile := "/tmp/customer_companies.csv"
	db.Exec("COPY customer_companies FROM '" + customerCompaniesTmpfile + "' WITH (format csv, header)")

	err3 := File("orders.csv","/tmp/orders.csv")
	if err3 != nil {
		fmt.Println(err3)
	}
	db.Exec("TRUNCATE TABLE orders")
	ordersTmpfile := "/tmp/orders.csv"
	db.Exec("COPY orders FROM '" + ordersTmpfile + "' WITH (format csv, header)")

	err4 := File("order_items.csv","/tmp/order_items.csv")
	if err4 != nil {
		fmt.Println(err4)
	}
	db.Exec("TRUNCATE TABLE order_items")
	orderItemsTmpfile := "/tmp/order_items.csv"
	db.Exec("COPY order_items FROM '" + orderItemsTmpfile + "' WITH (format csv, header)")

	err5 := File("deliveries.csv","/tmp/deliveries.csv")
	if err5 != nil {
		fmt.Println(err5)
	}
	db.Exec("TRUNCATE TABLE deliveries")
	deliveriesTmpfile := "/tmp/deliveries.csv"
	db.Exec("COPY deliveries FROM '" + deliveriesTmpfile + "' WITH (format csv, header)")

	return nil
}

func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
