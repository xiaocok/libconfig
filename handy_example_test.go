package libconfig_test

import (
	"fmt"
	"github.com/gitteamer/libconfig"
	"io/ioutil"
	"log"
)

// parse bytes demo.cfg
func ExampleParseBytes() {
	data, err := ioutil.ReadFile("testdata/demo.cfg")
	if err != nil {
		log.Fatal("read config file error: %s", err.Error())
	}

	fmt.Printf("version = %s\n", libconfig.GetString(data, "version"))
	fmt.Printf("application.window.title = %s\n", libconfig.GetString(data, "application", "window", "title"))
	fmt.Printf("application.window.size.w = %d\n",  libconfig.GetInt(data, "application", "window", "size", "w"))
	fmt.Printf("application.window.size.h = %d\n", libconfig.GetInt(data, "application", "window", "size", "h"))
	fmt.Printf("application.window.pos.x = %d\n", libconfig.GetInt(data, "application", "window", "pos", "x"))
	fmt.Printf("application.window.pos.y = %d\n", libconfig.GetInt(data, "application", "window", "pos", "y"))

	fmt.Printf("application.list[0][0] = %s\n", libconfig.GetString(data, "application", "list", "0", "0"))
	fmt.Printf("application.list[0][1] = %d\n",  libconfig.GetInt(data, "application", "list", "0", "1"))
	fmt.Printf("application.list[0][2] = %t\n", libconfig.GetBool(data, "application", "list", "0", "2"))
	fmt.Printf("application.list[1] = %f\n",  libconfig.GetFloat64(data, "application", "list", "1"))
	fmt.Printf("application.list[2][0] =%s\n", libconfig.GetString(data, "application", "list", "2", "0"))

	fmt.Printf("application.books[0].title = %s\n", libconfig.GetString(data, "application", "books", "0", "title"))
	fmt.Printf("application.books[0].author = %s\n", libconfig.GetString(data, "application", "books", "0", "author"))
	fmt.Printf("application.books[0].price = %f\n", libconfig.GetFloat64(data, "application", "books", "0", "price"))
	fmt.Printf("application.books[0].qty = %d\n", libconfig.GetInt(data, "application", "books", "0", "qty"))
	fmt.Printf("application.books[1].title = %s\n", libconfig.GetString(data, "application", "books", "1", "title"))
	fmt.Printf("application.books[1].author = %s\n", libconfig.GetString(data, "application", "books", "1", "author"))
	fmt.Printf("application.books[1].price = %f\n", libconfig.GetFloat64(data, "application", "books", "1", "price"))
	fmt.Printf("application.books[1].qty = %d\n",  libconfig.GetInt(data, "application", "books", "1", "qty"))

	fmt.Printf("application.misc.pi = %.9f\n",  libconfig.GetFloat64(data, "application", "misc", "pi"))
	fmt.Printf("application.misc.bigint = %s\n", libconfig.GetBigint(data, "application", "misc", "bigint").String())
	fmt.Printf("application.misc.columns[0] = %s\n", libconfig.GetString(data, "application", "misc", "columns", "0"))
	fmt.Printf("application.misc.columns[1] = %s\n", libconfig.GetString(data, "application", "misc", "columns", "1"))
	fmt.Printf("application.misc.columns[2] = %s\n",  libconfig.GetString(data, "application", "misc", "columns", "2"))
	fmt.Printf("application.misc.bitmask = %d\n", libconfig.GetInt(data, "application", "misc", "bitmask"))
	fmt.Printf("application.misc.bitmask_hex = %s\n", libconfig.GetHex(data, "application", "misc", "bitmask"))

	// Output:
	// version = 1.0
	// application.window.title = My Application
	// application.window.size.w = 640
	// application.window.size.h = 480
	// application.window.pos.x = 350
	// application.window.pos.y = 250
	// application.list[0][0] = abc
	// application.list[0][1] = 123
	// application.list[0][2] = true
	// application.list[1] = 1.234000
	// application.list[2][0] =
	// application.books[0].title = Treasure Island
	// application.books[0].author = Robert Louis Stevenson
	// application.books[0].price = 29.950000
	// application.books[0].qty = 5
	// application.books[1].title = Snow Crash
	// application.books[1].author = Neal Stephenson
	// application.books[1].price = 9.990000
	// application.books[1].qty = 8
	// application.misc.pi = 3.141592654
	// application.misc.bigint = 9223372036854775807
	// application.misc.columns[0] = Last Name
	// application.misc.columns[1] = First Name
	// application.misc.columns[2] = MI
	// application.misc.bitmask = 8131
	// application.misc.bitmask_hex = 0x1FC3
}

// parse include example4.cfg
func ExampleParseInclude() {
	var p libconfig.Parser

	v, err := p.ParseFile("testdata/example4.cfg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("books[0].title=%s\n", v.GetArray("books")[0].GetStringBytes("title"))
	fmt.Printf("books[0].title=%s\n", v.GetArray("books")[2].GetStringBytes("author"))
	fmt.Printf("books[0].extra1=%s\n", v.GetArray("books")[0].GetStringBytes("extra1"))
	fmt.Printf("books[3].extra1=%s\n", v.GetArray("books")[3].GetStringBytes("extra1"))
	fmt.Printf("books[3].extra2=%d\n", v.GetArray("books")[3].GetInt("extra2"))
	fmt.Printf("books[3].extra2=%d\n", v.GetInt("books", "3", "extra2"))

	// Output:
	// books[0].title=Treasure Island
	// books[0].title=Robert A. Heinlein
	// books[0].extra1=
	// books[3].extra1=bar
	// books[3].extra2=12345
	// books[3].extra2=12345
}
