package setup

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"os"
)

func PopulateDb(db *gorm.DB) error {

	_,err := CopyFile("/tmp","customers.csv")
	if err != nil {
		fmt.Println(err)
	}
	customersTmpfile := "/tmp/customers.csv"
	db.Exec("COPY customers FROM '" + customersTmpfile + "' WITH (format csv, header)")

	_,err2 := CopyFile("/tmp","customer_companies.csv")
	if err2 != nil {
		fmt.Println(err2)
	}
	customerCompaniesTmpfile := "/tmp/customer_companies.csv"
	db.Exec("COPY customer_companies FROM '" + customerCompaniesTmpfile + "' WITH (format csv, header)")

	_,err3 := CopyFile("/tmp","orders.csv")
	if err3 != nil {
		fmt.Println(err3)
	}
	ordersTmpfile := "/tmp/orders.csv"
	db.Exec("COPY orders FROM '" + ordersTmpfile + "' WITH (format csv, header)")

	_,err4 := CopyFile("/tmp","order_items.csv")
	if err4 != nil {
		fmt.Println(err4)
	}
	orderItemsTmpfile := "/tmp/order_items.csv"
	db.Exec("COPY order_items FROM '" + orderItemsTmpfile + "' WITH (format csv, header)")

	_,err5 := CopyFile("/tmp","deliveries.csv")
	if err5 != nil {
		fmt.Println(err5)
	}
	deliveriesTmpfile := "/tmp/deliveries.csv"
	db.Exec("COPY deliveries FROM '" + deliveriesTmpfile + "' WITH (format csv, header)")

	return nil
}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Open(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}


//func CopyFile(src, dst string) (err error) {
//	sfi, err := os.Stat(src)
//	if err != nil {
//		return
//	}
//	if !sfi.Mode().IsRegular() {
//		// cannot copy non-regular files (e.g., directories,
//		// symlinks, devices, etc.)
//		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
//	}
//	dfi, err := os.Stat(dst)
//	if err != nil {
//		if !os.IsNotExist(err) {
//			return
//		}
//	} else {
//		if !(dfi.Mode().IsRegular()) {
//			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
//		}
//		if os.SameFile(sfi, dfi) {
//			return
//		}
//	}
//	if err = os.Link(src, dst); err == nil {
//		return
//	}
//	err = copyFileContents(src, dst)
//	return
//}
//
//func copyFileContents(src, dst string) (err error) {
//	in, err := os.Open(src)
//	if err != nil {
//		return
//	}
//	defer in.Close()
//	out, err := os.Create(dst)
//	if err != nil {
//		return
//	}
//	defer func() {
//		cerr := out.Close()
//		if err == nil {
//			err = cerr
//		}
//	}()
//	if _, err = io.Copy(out, in); err != nil {
//		return
//	}
//	err = out.Sync()
//	return
//}