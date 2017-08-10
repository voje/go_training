package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
    //"io/ioutil"
    //"os"
    "sync"
    "reflect"
)

var main_url string = "https://en.wikipedia.org"
var lock sync.Mutex

type country struct {
    Name    string
    Pop     string
    City    string
    Cpop    string
}
var data = make(map[string] country)

func add_country(name string) {
    lock.Lock()
    defer lock.Unlock()
    data[name] = country{City: name}
}

func update_country(name string, key string, val string) {
    lock.Lock()
    defer lock.Unlock()
    if _,ok := data[name]; !ok {
        panic("update_country: country doesn't exist in data")
    }
    //todo
    v := reflect.ValueOf(data[name])
    f := v.FieldByName(key)
    f.SetBytes([]byte(val)) //ERR todo (unaddressable ...)
}

func print_data() {
    for _,d := range data {
        fmt.Printf("Name: \t%s\nPop: \t%s\nCity: \t%s\nCpop: \t%s\n\n", d.Name, d.Pop, d.City, d.Cpop)
    }
}

func check(e error) {
    if e != nil {
        fmt.Println(e.Error())
    }
}

func scrape_city(url string, cname string) {
    page,err := http.Get(url)
    check(err)
    tok := html.NewTokenizer(page.Body)
    find_table(tok, "infobox geography vcard")
    scraping_pop := false
    var t html.Token
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            msg := fmt.Sprintf("failed scraping city of: %s", cname)
            panic(msg)
        case html.TextToken:
            t := tok.Token()
            if t.Data == "Population " || t.Data == "Population" {
                scraping_pop = true
            }
        case html.StartTagToken:
            t = tok.Token()
            if scraping_pop && t.Data == "td" {
                tok.Next()
                t = tok.Token()
                update_country(cname, "Cpop", t.Data)
                return
            }
        }
    }
}

func get_attr_val(t html.Token, key string) string {
    for _,att := range t.Attr {
        if att.Key == key {
            return att.Val
        }
    }
    return "err"
}

func find_table(tok *html.Tokenizer, classname string) html.Token {
    for {
        tt := tok.Next()
        switch tt {
        case html.ErrorToken:
            panic("find_table EOF")
        case html.StartTagToken:
            t := tok.Token()
            if t.Data == "table" {
                for _,at := range t.Attr {
                    if at.Key == "class" && at.Val == classname {
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
    //fmt.Println(tag)
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
    page,err := http.Get(main_url + "/wiki/Asia")
    check(err)
    //f,err := os.Create("test.html")
    //page.Write(f)
    /*
    var data []byte
    data,_ = ioutil.ReadAll(page.Body)
    fmt.Println(string(data))
    */
    tok := html.NewTokenizer(page.Body)

    find_table(tok, "sortable wikitable")
    for {
        tt := tok.Next()
        if tt == html.StartTagToken {
            t := tok.Token()
            if t.Data == "tr" {
                skip_subtree(tok, t) //skip first row
                break
            }
        }
    }
    var t html.Token
    var td int
    var cname string
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
                    t = tok.Token()
                    cname = t.Data
                    add_country(t.Data)
                case 3:
                    tok.Next()
                    t = tok.Token()
                    update_country(cname, "Pop", t.Data)
                case 5:
                    t = find_a(tok)
                    url := get_attr_val(t, "href")
                    tok.Next()
                    t = tok.Token()
                    update_country(cname, "City", t.Data)
                    scrape_city(main_url + url, cname)
                }
            }
        case html.EndTagToken:
            t = tok.Token()
            if t.Data == "table" {
                fmt.Println("Finished reading table")
                print_data()
                return
            }
        }
    }
}






