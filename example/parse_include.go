package main

import (
	"fmt"
	"github.com/gitteamer/libconfig"
	"log"
)

func ParseInclude() {
	var p libconfig.Parser

	v, err := p.ParseFile("testdata/example4.cfg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("str"))
	fmt.Printf("int=%d\n", v.GetInt("int"))
	fmt.Printf("float=%f\n", v.GetFloat64("float"))
	fmt.Printf("bool=%v\n", v.GetBool("bool"))
	fmt.Printf("arr.1=%s\n", v.GetStringBytes("arr", "1"))

	// Output:
	// foo=bar
	// int=123
	// float=1.230000
	// bool=true
	// arr.1=foo
}
