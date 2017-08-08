package main

import (
    "fmt"
    "net/http"
    "time"
    "golang.org/x/net/html"
    "os"
    "io"
    "strings"
)

func get_attr_val(fnd string, t *html.Token) (ok bool, res string) {
    ok = true
    res = ""
    for _,attr := range t.Attr {
        if attr.Key == fnd {
            res = attr.Val
            return
        }
    }
    ok = false
    return
}

func check(e error) {
    if e != nil {
        panic(e.Error())
    }
}

var visited = make(map[string]bool)

func generic_read_page(url string) io.Reader{
    filepath := "./cache/" + strings.Replace(url, "/", "_", -1)
    page,err := os.Open(filepath)
    if err != nil {
        page,err := http.Get(url)
        check(err)
        f,err := os.Create(filepath)
        defer f.Close()
        check(err)
        page.Write(f)
    } else {
        fmt.Printf("File %s found in cache.\n", filepath)
    }
    return page
}

func read_city_page(url string) {
    url = "http://forgottenrealms.wikia.com" + url
    pageFile := generic_read_page(url)
    tok := html.NewTokenizer(pageFile)
    scraping := false
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            fmt.Println("EOF")
            return
        case html.StartTagToken:
            t := tok.Token()
            if t.Data == "h3" {
                tok.Next();
                t := tok.Token()
                if t.Data == "Population" {
                    scraping = true
                }
            }
        case html.TextToken:
            if (scraping) {
                t := tok.Token()
                fmt.Println(t.Data) //TODO bump in the road.. webpage is dynamic
                return
            }
        }
    }
}

func loop_main_page() {
    var url = "http://forgottenrealms.wikia.com/wiki/Category:Large_cities"
    pageFile := generic_read_page(url)
    tok := html.NewTokenizer(pageFile)

    scraping := false
    for{
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            fmt.Println("End of page")
            return
        case html.StartTagToken:
            t := tok.Token()
            if t.Data == "div" && !scraping {
                if _,res := get_attr_val("class", &t); res == "mw-content-ltr" {
                    scraping = true
                }
            } else if t.Data == "a" && scraping {
                _,res := get_attr_val("href", &t)
                if _,ok := visited[res]; ok {
                    continue
                } else {
                    visited[res] = true
                }
                go read_city_page(res)
            }
        case html.EndTagToken:
            t := tok.Token()
            if t.Data == "div" && scraping {
                scraping = false
                return
            }
        }
    }
}

func main() {
    fmt.Println("Starting pop.go")
    loop_main_page()

    fmt.Println("Waiting for goroutines to finish")
    time.Sleep(time.Second * 2) //nope, use channels
}
