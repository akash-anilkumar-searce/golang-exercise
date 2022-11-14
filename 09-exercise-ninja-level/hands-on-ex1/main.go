package main

import (
    "fmt"
    "sync"
)

func foo() {
    fmt.Println("Function Foo")
    wg.Done()
}


func bar() {
    fmt.Println("Function Bar")
    wg.Done()
}

var wg sync.WaitGroup

func main() {
    wg.Add(2)
    go foo()
    go bar()
    wg.Wait()
}