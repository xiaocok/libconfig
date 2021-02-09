package main

import (
	"fmt"
	"github.com/gitteamer/libconfig"
	"io/ioutil"
	"log"
)

func ParseBytes() {
	data, err := ioutil.ReadFile("testdata/demo.cfg")
	if err != nil {
		log.Fatal("read config file error: %s", err.Error())
	}

	version := libconfig.GetString(data, "version")
	title := libconfig.GetString(data, "application", "window", "title")
	size_w := libconfig.GetInt(data, "application", "window", "size", "w")
	size_h := libconfig.GetInt(data, "application", "window", "size", "h")
	pos_x := libconfig.GetInt(data, "application", "window", "pos", "x")
	pos_y := libconfig.GetInt(data, "application", "window", "pos", "y")

	list_0_0 := libconfig.GetString(data, "application", "list", "0", "0")
	list_0_1 := libconfig.GetInt(data, "application", "list", "0", "1")
	list_0_2 := libconfig.GetBool(data, "application", "list", "0", "2")
	list_1 := libconfig.GetFloat64(data, "application", "list", "1")
	list_2_0 := libconfig.GetString(data, "application", "list", "2", "0")

	books_0_title := libconfig.GetString(data, "application", "books", "0", "title")
	books_0_author := libconfig.GetString(data, "application", "books", "0", "author")
	books_0_price := libconfig.GetFloat64(data, "application", "books", "0", "price")
	books_0_qty := libconfig.GetInt(data, "application", "books", "0", "qty")
	books_1_title := libconfig.GetString(data, "application", "books", "1", "title")
	books_1_author := libconfig.GetString(data, "application", "books", "1", "author")
	books_1_price := libconfig.GetFloat64(data, "application", "books", "1", "price")
	books_1_qty := libconfig.GetInt(data, "application", "books", "1", "qty")

	misc_pi := libconfig.GetFloat64(data, "application", "misc", "pi")
	misc_bigint := libconfig.GetBigint(data, "application", "misc", "bigint")
	misc_columns_0 := libconfig.GetString(data, "application", "misc", "columns", "0")
	misc_columns_1 := libconfig.GetString(data, "application", "misc", "columns", "1")
	misc_columns_2 := libconfig.GetString(data, "application", "misc", "columns", "2")
	misc_bitmask := libconfig.GetInt(data, "application", "misc", "bitmask")
	misc_bitmask_hex := libconfig.GetHex(data, "application", "misc", "bitmask")

	fmt.Printf("version = %s\n", version)
	fmt.Printf("title = %s\n", title)
	fmt.Printf("size_w = %d\n", size_w)
	fmt.Printf("size_h = %d\n", size_h)
	fmt.Printf("pos_x = %d\n", pos_x)
	fmt.Printf("pos_y = %d\n", pos_y)

	fmt.Printf("list_0_0 = %s\n", list_0_0)
	fmt.Printf("list_0_1 = %d\n", list_0_1)
	fmt.Printf("list_0_2 = %t\n", list_0_2)
	fmt.Printf("list_1 = %f\n", list_1)
	fmt.Printf("list_2_0 =%s\n", list_2_0)

	fmt.Printf("books_0_title = %s\n", books_0_title)
	fmt.Printf("books_0_author = %s\n", books_0_author)
	fmt.Printf("books_0_price = %f\n", books_0_price)
	fmt.Printf("books_0_qty = %d\n", books_0_qty)
	fmt.Printf("books_1_title = %s\n", books_1_title)
	fmt.Printf("books_1_author = %s\n", books_1_author)
	fmt.Printf("books_1_price = %f\n", books_1_price)
	fmt.Printf("books_1_qty = %d\n", books_1_qty)

	fmt.Printf("misc_pi = %.9f\n", misc_pi)
	fmt.Printf("misc_bigint = %s\n", misc_bigint.String())
	fmt.Printf("misc_columns_0 = %s\n", misc_columns_0)
	fmt.Printf("misc_columns_1 = %s\n", misc_columns_1)
	fmt.Printf("misc_columns_2 = %s\n", misc_columns_2)
	fmt.Printf("misc_bitmask = %d\n", misc_bitmask)
	fmt.Printf("misc_bitmask_hex = %s\n", misc_bitmask_hex)

	// Output:
	// version = 1.0
	// title = My Application
	// size_w = 640
	// size_h = 480
	// pos_x = 350
	// pos_y = 250
	// list_0_0 = abc
	// list_0_1 = 123
	// list_0_2 = true
	// list_1 = 1.234000
	// list_2_0 =
	// books_0_title = Treasure Island
	// books_0_author = Robert Louis Stevenson
	// books_0_price = 29.950000
	// books_0_qty = 5
	// books_1_title = Snow Crash
	// books_1_author = Neal Stephenson
	// books_1_price = 9.990000
	// books_1_qty = 8
	// misc_pi = 3.141592654
	// misc_bigint = 9223372036854775807
	// misc_columns_0 = Last Name
	// misc_columns_1 = First Name
	// misc_columns_2 = MI
	// misc_bitmask = 8131
	// misc_bitmask_hex = 0x1FC3
}