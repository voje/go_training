package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    //"io/ioutil"
    //"os"
)

func check(e error) {
    if e != nil {
        fmt.Println(e.Error())
    }
}

func skip_subtree(tok *html.Tokenizer, t html.Token) {
    depth := 0
    tag := t.Data
    fmt.Println(tag)
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            panic("skip_subtree EOF")
        case html.StartTagToken:
            t = tok.Token()
            if t.Data == tag {depth++}
        case html.EndTagToken:
            t = tok.Token()
            if t.Data == tag {
                if depth == 0 {
                    return
                }
                depth--
            }
        }
    }
}

func main() {
    fmt.Println("Starting capit.go")
    page,err := http.Get("https://en.wikipedia.org/wiki/Asia")
    check(err)
    //f,err := os.Create("test.html")
    //page.Write(f)
    /*
    var data []byte
    data,_ = ioutil.ReadAll(page.Body)
    fmt.Println(string(data))
    */
    tok := html.NewTokenizer(page.Body)

    inTable := false
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            fmt.Println("EOF")
            return
        case html.StartTagToken:
            t := tok.Token()
            switch {
            case t.Data == "table":
                if t.Attr[0].Val == "sortable wikitable" {
                    inTable = true
                }
            case t.Data == "tr" && inTable:
                skip_subtree(tok, t)
                ntd := 0
                for ntd <= 5 {
                    tok.Next()
                    t := tok.Token()
                    if t.Data == "td" { ntd++ }
                    //fmt.Printf("[ntd: %d] %s\n", ntd, t.Data)
                    switch {
                    case ntd == 3:
                        tok.Next() //<td>
                        tok.Next() //<a>
                        t = tok.Token()
                        fmt.Println(t.Data)
                    }
                }
            }
        case html.EndTagToken:
            t := tok.Token()
            if inTable && t.Data == "table" {
                fmt.Println("Finished reading table")
                return
            }
        }
    }
}
