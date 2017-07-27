package main

import (
    "fmt"
    "html/template"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq" //import a package for its side-effects only
    //"reflect"
    "dnd/db_structs"
)

var templates = template.Must(template.ParseFiles("./templates/index.html"))
var db *sql.DB
var err error
var m map[string]interface{}

func err500(w http.ResponseWriter, err error) bool {
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) //500
        fmt.Println("err500")
        return true;
    }
    return false;
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
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
    //todo
}

func main(){

    dbname := "dnd"
    uname := "gopher"

    fmt.Println("Starting dnd.main")
    m = make(map[string]interface{})
    db_structs.init_db_structs(&m)
    fmt.Println(m["npc"])

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

    //CRUD: create, read, update, delete
    http.HandleFunc("/index", handleIndex)
    http.HandleFunc("/create", handleCreate)

    http.ListenAndServe(":8001", nil)

}
