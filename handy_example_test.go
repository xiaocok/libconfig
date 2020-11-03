package libconfig_test

import (
	"fmt"
	"fastjson"
)

func ExampleGetString() {
	data := []byte(`
version = "1.0";

application:
{
  window:
  {
    title = "My Application";
    size = { w = 640; h = 480; };
    pos = { x = 350; y = 250; };
  };

  list = ( ( "abc", 123, true ), 1.234, ( /* an empty list */) );

  books = ( { title  = "Treasure Island";
              author = "Robert Louis Stevenson";
              price  = 29.95;
              qty    = 5; },
            { title  = "Snow Crash";
              author = "Neal Stephenson";
              price  = 9.99;
              qty    = 8; } );

  misc:
  {
    pi = 3.141592654;
    bigint = 9223372036854775807L;
    columns = [ "Last Name", "First Name", "MI" ];
    bitmask = 0x1FC3;
  };
};
`)

	s := fastjson.GetString(data, "version")
	fmt.Printf("root.version = %v", s)

	// Output:
	// root.version = 640
}

func ExampleGetInt() {
	data := []byte(`{"foo": [233,true, {"bar": [2343]} ]}`)

	n1 := fastjson.GetInt(data, "foo", "0")
	fmt.Printf("data.foo[0] = %d\n", n1)

	n2 := fastjson.GetInt(data, "foo", "2", "bar", "0")
	fmt.Printf("data.foo[2].bar[0] = %d\n", n2)

	// Output:
	// data.foo[0] = 233
	// data.foo[2].bar[0] = 2343
}

func ExampleExists() {
	data := []byte(`{"foo": [1.23,{"bar":33,"baz":null}]}`)

	fmt.Printf("exists(data.foo) = %v\n", fastjson.Exists(data, "foo"))
	fmt.Printf("exists(data.foo[0]) = %v\n", fastjson.Exists(data, "foo", "0"))
	fmt.Printf("exists(data.foo[1].baz) = %v\n", fastjson.Exists(data, "foo", "1", "baz"))
	fmt.Printf("exists(data.foobar) = %v\n", fastjson.Exists(data, "foobar"))
	fmt.Printf("exists(data.foo.bar) = %v\n", fastjson.Exists(data, "foo", "bar"))

	// Output:
	// exists(data.foo) = true
	// exists(data.foo[0]) = true
	// exists(data.foo[1].baz) = true
	// exists(data.foobar) = false
	// exists(data.foo.bar) = false
}
