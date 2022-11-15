
package main

import "fmt"

type customErr struct {
    s string
}

func foo(e error) {
    fmt.Println(e)
}

func (e customErr) Error() string {
    return "call Custom Error Function"
}

func main() {
    ce := customErr{"custom Err"}
    foo(ce)
}