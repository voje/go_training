package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

func Walk(t *tree.Tree, ch chan int) {
  defer close(ch)
  var walk func(t *tree.Tree)
  walk = func(t *tree.Tree) {
    if t == nil {return}
    walk(t.Left)
    ch <- t.Value
    walk(t.Right)
  }
  walk(t)
  //close(ch) possible instead of defer
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
   ch1 := make(chan int)
   go Walk(t1, ch1)
   ch2 := make(chan int)
   go Walk(t2, ch2)
   for {
     v1,ok := <- ch1
     if !ok {
       fmt.Println("ch1 closed")
       break
     }
     v2 := <- ch2
     if v1 != v2 {return false}
   }
   return true
}

func main() {
  // Workaround for closing channel when recursion ends: closure in Walk: recursive part is walk()
  fmt.Print("Tree traversal\n")
  res := Same(tree.New(1), tree.New(1))
  fmt.Printf("Trees are the same: %v.\n", res)
}
