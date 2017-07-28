package main

import (
    "net/http"
    "fmt"
)

type route struct {
    pattern     string
    handler     http.Handler
}

type Regex_mux struct {
    routes  []*route
}

func (rm *Regex_mux) HandleFunc (pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
    rm.routes = append(rm.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (rm Regex_mux) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Println(" - ServeHTTP")
    if rm.routes == nil {return}
    for _,rout := range rm.routes {
        if rout.pattern == r.URL.Path {
            rout.handler.ServeHTTP(w,r)
            return
        }
    }
}

func handlePizza(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Pizza has arrived!")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("I am root!")
}

func main() {
    fmt.Println("Starting mux.go")
    regmux := Regex_mux{}

    regmux.HandleFunc("/pizza", handlePizza)
    regmux.HandleFunc("/", handleRoot)

    http.ListenAndServe(":8002", regmux)
}
