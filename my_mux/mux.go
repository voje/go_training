package my_mux

import (
    "net/http"
    //"fmt"
    "regexp"
)

type route struct {
    pattern     *regexp.Regexp
    handler     http.Handler
}

type Regex_mux struct {
    routes  []*route
}

func (rm *Regex_mux) HandleFunc (pattern *regexp.Regexp, handler func(w http.ResponseWriter, r *http.Request)) {
    rm.routes = append(rm.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (rm Regex_mux) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    //remove trailing '/'
    if r.URL.Path[len(r.URL.Path)-1] == '/' {
        r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
    }

    for _,rout := range rm.routes {
        if rout.pattern.MatchString(r.URL.Path) {
            rout.handler.ServeHTTP(w,r)
            return
        }
    }
    http.Redirect(w, r, "/", http.StatusFound)
}

/*
func handlePizza(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Pizza has arrived!")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("I am root!")
}

func main() {
    fmt.Println("Starting mux.go")

    // regex101.com
    // Expression remove trailing / before parsing
    // for general checking
    // `^(\/[a-zA-Z_0-9]+\/(create|read|update|delete))?$`
    regmux := Regex_mux{}

    regPizza,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/pizza$`)
    regmux.HandleFunc(regPizza, handlePizza)

    regRoot,_ := regexp.Compile(`^$`)
    regmux.HandleFunc(regRoot, handleRoot)

    http.ListenAndServe(":8002", regmux)
}
*/
