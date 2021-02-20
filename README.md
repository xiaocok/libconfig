# libconfig
A golang language library for reading libconfig file

## Note
* Modified based on [valyala/fastjson](https://github.com/valyala/fastjson), fastjson is a json library.
* Refer to scan character mode to realize parsing configuration file.<br/>
* Implement the function of parsing configuration file similar to [hyperrealm/libconfig](https://github.com/hyperrealm/libconfig), that is a C/C++ library.<br/>
* The project is still in the development stage, please do not use it in the production environment.<br/>

## support
* skip # comment
* skip // comment
* skip /* */ comment
* scalarvalue
* Hexadecimal data
* big int
* array
* group
* list
* @include

## example
### parse strings
```go
package main
import (
    "fmt"
    "github.com/gitteamer/libconfig"
)
func main()  {
    data := []byte(`foo="bar"; baz=1234;`)
    foo := libconfig.GetString(data, "foo")
    fmt.Println("foo = %s\n", foo)
}
```

### parse file
```go
package main
import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/gitteamer/libconfig"
)
func main()  {
    data, err := ioutil.ReadFile("testdata/demo.cfg")
    if err != nil {
        log.Fatal("read config file error: ", err.Error())
    }
    
    fmt.Printf("version = %s\n", libconfig.GetString(data, "version"))
}
```

### parse with object
> #### strings
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main()  {
    var p libconfig.Parser
    
    v, err := p.Parse(`foo="bar"; baz=1234;`)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("foo = %s\n", v.Get("foo").String())
    fmt.Printf("foo = %s\n", string(v.GetStringBytes("foo")))
}
```

> #### file
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main()  {
    var p libconfig.Parser
    
    v, err := p.ParseFile("testdata/example4.cfg")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("books[0].title=%s\n", v.GetArray("books")[0].GetStringBytes("title"))
}
```

## parse element
### scalarvalue、Hexadecimal data、big int
```go
package main
import (
    "fmt"
    "github.com/gitteamer/libconfig"
)
func main()  {
    data := []byte(`foo="bar"; baz=1234; bigint=9223372036854775807L; float=1.0; bool=false;`)
    
    // string
    fmt.Printf("foo = %s\n", libconfig.GetString(data, "foo"))
    
    // bytes
    fmt.Printf("foo = %s\n", libconfig.GetBytes(data, "foo"))
    
    // int
    fmt.Printf("baz = %d\n", libconfig.GetInt(data, "baz"))
    
    // hex value
    fmt.Printf("baz = %s\n", libconfig.GetHex(data, "baz"))
    
    //big int
    fmt.Printf("bigint = %s\n", libconfig.GetBigint(data, "bigint").String())
    
    // float 64
    fmt.Printf("float = %v\n", libconfig.GetFloat64(data, "float"))
    
    // bool
    fmt.Printf("bool = %v\n", libconfig.GetBool(data, "bool"))
}
```

### group
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main()  {
    data := []byte(`foo={bar= 1234; baz=0;};`)
    
    // handy parse
    fmt.Printf("foo.bar = %s\n", libconfig.GetString(data, "foo", "bar"))

    //object parse
    var (
        p libconfig.Parser
        v *libconfig.Value    
    )
    v, err := p.ParseBytes(data)
    if err != nil {
    	log.Fatal(err)
    }

    // use get with multiple parameters and check error
    item, err := v.Get("foo", "bar").StringBytes()
    if err != nil {
    	log.Fatal(err)
    }
    fmt.Printf("foo.bar = %s\n", string(item))
    
    // use get with call chaining
    fmt.Printf("foo.bar = %s\n", string(v.Get("foo").Get("bar").GetStringBytes()))
    
    // use get array object
    foo := v.GetObject("foo")
    bar := foo.Get("bar")
    fmt.Printf("my_array[0] = %s\n", bar.String())
}
```

### array
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main()  {
	data := []byte(`my_array = ["CT","CA","TX","NV","FL"];`)

    // handy parse
    fmt.Printf("my_array[0] = %s\n", libconfig.GetString(data, "my_array", "0"))

    // object parse
    var (
        p libconfig.Parser
        v *libconfig.Value
    )
    
    v, err := p.ParseBytes(data)
    if err != nil {
          log.Fatal(err)
    }

    // use get with multiple parameters and check error
    item, err := v.Get("my_array", "0").StringBytes()
    if err != nil {
          log.Fatal(err)
    }
    fmt.Printf("my_array[0] = %s\n", string(item))
    
    // use get with call chaining
    fmt.Printf("my_array[0] = %s\n", v.Get("my_array").Get("0").String())
    
    // use get array object
    my_array := v.GetArray("my_array")
    fmt.Printf("my_array[0] = %s\n", my_array[0].String())
}
```

### list
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main() {
    data := []byte(`list = ( ( "abc", 123, true ), 1.234, ( /* an empty list */) );`)
    
    // handy parse
    fmt.Printf("list[0][0] = %s\n", libconfig.GetString(data, "list", "0", "0"))
    
    // object parse
    var (
        p libconfig.Parser
        v *libconfig.Value
    )
    
    v, err := p.ParseBytes(data)
    if err != nil {
        log.Fatal(err)
    }
    
    // use get with multiple parameters and check error
    item, err := v.Get("list", "0", "0").StringBytes()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("list[0][0] = %s\n", string(item))
    
    // use get with call chaining
    fmt.Printf("list[0][0] = %s\n", v.Get("list").Get("0").Get("0").String())
    
    // use get array object
    list := v.GetArray("list")
    first_array := list[0].GetArray()
    fmt.Printf("list[0][0] = %s\n", first_array[0].String())
}
```

### mix
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main() {
    data := []byte(`foo=([{bar=1234; baz=0;}],)`)
    
    // handy parse
    fmt.Printf("foo[0][0].bar = %d\n", libconfig.GetInt(data, "foo", "0", "0", "bar"))
    
    // object parse
    var (
        p libconfig.Parser
        v *libconfig.Value
    )
    
    v, err := p.ParseBytes(data)
    if err != nil {
        log.Fatal(err)
    }
    
    // use get with multiple parameters and check error
    item, err := v.Get("foo", "0", "0", "bar").Int()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("foo[0][0].bar = %d\n", item)
    
    // use get with call chaining
    fmt.Printf("foo[0][0].bar = %d\n", v.Get("foo").Get("0").Get("0").GetInt("bar"))
    
    // use get list object
    foo := v.GetArray("foo")
    first_list := foo[0].GetArray()
    fmt.Printf("foo[0][0].bar = %s\n", first_list[0].GetObject().Get("bar"))
}
```

### include
```go
package main
import (
    "fmt"
    "log"
    "github.com/gitteamer/libconfig"
)
func main(){
    var p libconfig.Parser
    
    // testdata/example4.cfg
    v, err := p.ParseFile("testdata/example4.cfg")
    if err != nil {
        log.Fatal(err)
    }
    
    // @include "cfg_includes/book*.cfg"
    fmt.Printf("books[0].title=%s\n", v.GetArray("books")[0].GetStringBytes("title"))
    fmt.Printf("books[0].title=%s\n", v.Get("books", "0").GetStringBytes("title"))
    fmt.Printf("books[0].title=%s\n", v.Get("books").Get("0").GetStringBytes("title"))
    fmt.Printf("books[0].author=%s\n", v.GetArray("books")[2].GetStringBytes("author"))
    
    //@include "cfg_includes/cfg_subincludes/*.cfg"	
    fmt.Printf("books[0].extra1=%s\n", v.GetArray("books")[0].GetStringBytes("extra1"))
    fmt.Printf("books[3].extra1=%s\n", v.GetArray("books")[3].GetStringBytes("extra1"))
    fmt.Printf("books[3].extra2=%d\n", v.GetArray("books")[3].GetInt("extra2"))
    fmt.Printf("books[3].extra2=%d\n", v.GetInt("books", "3", "extra2"))
}
```

### a complex demo
```go
package main
import (
    "fmt"
    "log"
    "io/ioutil"
    "github.com/gitteamer/libconfig"
)
func main(){
    data, err := ioutil.ReadFile("testdata/demo.cfg")
    if err != nil {
        log.Fatal("read config file error: ", err.Error())
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
```