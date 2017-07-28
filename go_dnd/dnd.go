package main

import (
    "fmt"
    "html/template"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq" //import a package for its side-effects only
    //"reflect"
    "strings"
    "github.com/go_training/my_mux"
    "regexp"
)

var templates = template.Must(template.ParseFiles("./templates/index.html", "./templates/read.html"))
var db *sql.DB
var err error
var m map[string]interface{}

type Post struct{
    Name    string
    Rows    interface{}
}

func err500(w http.ResponseWriter, err error) bool {
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) //500
        fmt.Println("err500")
        return true;
    }
    return false;
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handleIndex")
    //make query, list tables
    query := `SELECT table_name
    FROM information_schema.tables
    WHERE table_schema='public';`
    rows, err := db.Query(query)
    defer rows.Close() //executes when function returns
    if err500(w, err) {return}

    var srows []string
    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err500(w, err) {return}
        srows = append(srows, name)
    }

    //write tables to template index.html
    err = templates.ExecuteTemplate(w, "index.html", srows)
    if err500(w, err) {return}
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handleCreate")
    //todo
}

func handleRead(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handleRead")
    //query a table (name from url)
    arr_name := strings.Split(r.URL.Path, "/")
    table_name := arr_name[1] //[0] == ""
    query := fmt.Sprintf("SELECT id,name FROM %s;", table_name)
    rows,err := db.Query(query)
    if err500(w, err) {return}
    defer rows.Close()
    var srows []Npc
    for rows.Next() {
        np := Npc{}
        err = rows.Scan(&np.Id, &np.Name)
        if err500(w, err) {return}
        srows = append(srows, np)
    }
    err = templates.ExecuteTemplate(w, "read.html", Post{Name: table_name, Rows: srows})
    if err500(w, err) {return}
}

func main(){
    dbname := "dnd"
    uname := "gopher"

    fmt.Println("Starting dnd.main")

    //structs that represent database tables in ./db_structs.go
    m = init_db_structs()

    //connect to postgres db
    db, err = sql.Open("postgres",
        fmt.Sprintf("dbname=%s user=%s password=gopher", dbname, uname))
    if err != nil {
        panic("Couldn't connect to database.")
    }
    if err = db.Ping(); err != nil {
        panic("Couldn't ping database.")
    } else {
        fmt.Printf("%s connected to database: %s.\n", uname, dbname)
    }

    //custom multiplexer (cuts away trailing '/')
    mux := my_mux.Regex_mux{}

    //CRUD: create, read, update, delete
    regIndex,_ := regexp.Compile(`^(\/index)?$`);
    mux.HandleFunc(regIndex, handleIndex)

    regCreate,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/create$`);
    mux.HandleFunc(regCreate, handleCreate)

    regRead,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/read$`);
    mux.HandleFunc(regRead, handleRead)

    http.ListenAndServe(":8001", mux)
}
