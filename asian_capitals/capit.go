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

func find_table(tok *html.Tokenizer) html.Token {
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            panic("find_table EOF")
        case html.StartTagToken:
            t := tok.Token()
            if t.Data == "table" {
                for _,at := range t.Attr {
                    if at.Key == "class" && at.Val == "sortable wikitable" {
                        return t
                    }
                }
            }
        }
    }
}

func find_a(tok *html.Tokenizer) html.Token {
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            panic("find_a EOF")
        case html.StartTagToken:
            t := tok.Token()
            if t.Data == "a" {
                return t
            }
        }
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

    find_table(tok)
    for {
        tt := tok.Next()
        if tt == html.StartTagToken {
            t := tok.Token()
            if t.Data == "tr" {
                skip_subtree(tok, t) //skip first row
            }
        }
    }
    var t html.Token
    var td int
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            panic("main EOF")
        case html.StartTagToken:
            t = tok.Token()
            switch {
            case t.Data == "tr":
                td = 0
            case t.Data == "td":
                td++
                switch td {
                case 2:
                    find_a(tok)
                    tok.Next()
                }
            }
        case html.EndTagToken:
            t = tok.Token()
            if t.Data == "table" {
                fmt.Println("Finished reading table")
                return
            }
        }
    }
}






