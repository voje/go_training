package my_str

import (
    "fmt"
)

func str1() {
    var name string
    name = "kristjan"
    fmt.Printf("%q, %x, %s\n", name, name, name)

    var bstr []byte
    bstr = []byte(name)
    name = "kr"
    fmt.Printf("%q\n", bstr)

    var newstr string
    newstr = string(bstr)
    fmt.Printf("%q\n", newstr)
}
