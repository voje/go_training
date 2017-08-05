package main

import (
    "fmt"
    "time"
    "testing"
    "sync"
)

// Messing with goroutines

func prnt(s string) {
    fmt.Printf("[%s]\t\t%s\n", time.Now().String(), s)
}

func main() {
    prnt("tests")
}

func Test1(t *testing.T) {
    prnt("Test1 start")
    go gotest()
    time.Sleep(time.Millisecond * 10)
    prnt("Test1 end")
}

func gotest() {
    prnt("test groutine")
}

var lock sync.Mutex
func Test2(t *testing.T) {
    prnt("Test2 start")
    counter := 0
    for i:=0; i<10; i++ {
        go gotest2(&counter)
        time.Sleep(time.Millisecond*10) //remove sleep and you get random outcome (1,1,1,1,1,2,2,2,3,3,...)
        fmt.Printf("%d...", counter)
    }
    prnt("Test2 end")
}

func gotest2(counter *int) {
    lock.Lock() //if you lock this goroutine, other increments don't do anything (1,1,1,1,1,1,1,1,1,...)
    defer lock.Unlock()
    *counter ++
}
