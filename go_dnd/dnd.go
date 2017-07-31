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
    "strconv"
)

var templates = template.Must(template.ParseFiles(
  "templates/index.html",
  "templates/read.html",
  "templates/create.html",
  "templates/toggle_update.html",
  "templates/update.html",
  "templates/toggle_delete.html" ))
var db *sql.DB
var err error

type Post struct{
    Name    string
    Id      int32
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
    table_name := strings.Split(r.URL.Path, "/")[1] //[0] == ""
    //fmt.Println(r.Method)
    if r.Method == "GET" {
      err = templates.ExecuteTemplate(w, "create.html", table_name)
      if err500(w, err) {return}
    } else if r.Method == "POST" {
      contents := r.FormValue("name")
      query := fmt.Sprintf("INSERT INTO %s (name) VALUES ('%s')", table_name, contents)
      _,err := db.Query(query)
      if err500(w, err) {return}
      http.Redirect(w, r, fmt.Sprintf("/%s/read", table_name), http.StatusFound)
    }
}

func handleRead(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handleRead")
    //query a table (name from url)
    table_name := strings.Split(r.URL.Path, "/")[1] //[0] == ""
    query := fmt.Sprintf("SELECT id,name FROM %s;", table_name)
    rows,err := db.Query(query)
    if err500(w, err) {return}
    defer rows.Close()
    var srows []TableRow
    for rows.Next() {
        trow := new_db_struct(table_name)
        //err = rows.Scan(&np.Id, &np.Name)
        err = trow.Scan(rows)
        if err500(w, err) {return}
        srows = append(srows, trow)
    }
    err = templates.ExecuteTemplate(w, "read.html", Post{Name: table_name, Rows: srows})
    if err500(w, err) {return}
}

func handleToggleUpdate(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handleToggleUpdate")
    table_name := strings.Split(r.URL.Path, "/")[1] //[0] == ""
    if r.Method == "GET" {
      //query a table (name from url)
      query := fmt.Sprintf("SELECT id,name FROM %s;", table_name)
      rows,err := db.Query(query)
      if err500(w, err) {return}
      defer rows.Close()
      var srows []TableRow
      for rows.Next() {
          trow := new_db_struct(table_name)
          //err = rows.Scan(&np.Id, &np.Name)
          err = trow.Scan(rows)
          if err500(w, err) {return}
          srows = append(srows, trow)
      }
      err = templates.ExecuteTemplate(w, "toggle_update.html", Post{Name: table_name, Rows: srows})
      if err500(w, err) {return}
    }
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
  fmt.Println("handleUpdate")
  table_name := strings.Split(r.URL.Path, "/")[1] //[0] == ""
  tid,_ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
  id := int32(tid)
  if r.Method == "GET" {
    err := templates.ExecuteTemplate(w, "update.html", Post{Name: table_name, Id: id})
    if err500(w, err) {return}
  } else if r.Method == "POST" {
    query := fmt.Sprintf("UPDATE %s SET name = '%s' WHERE id = %d;", table_name, r.FormValue("name"), id)
    _,err := db.Query(query)
    if err500(w, err) {return}
    http.Redirect(w, r, fmt.Sprintf("/%s/read/", table_name), http.StatusFound)
  }
}

func handleToggleDelete(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handleToggleDelete")
    table_name := strings.Split(r.URL.Path, "/")[1] //[0] == ""
    if r.Method == "GET" {
      //query a table (name from url)
      query := fmt.Sprintf("SELECT id,name FROM %s;", table_name)
      rows,err := db.Query(query)
      if err500(w, err) {return}
      defer rows.Close()
      var srows []TableRow
      for rows.Next() {
          trow := new_db_struct(table_name)
          //err = rows.Scan(&np.Id, &np.Name)
          err = trow.Scan(rows)
          if err500(w, err) {return}
          srows = append(srows, trow)
      }
      err = templates.ExecuteTemplate(w, "toggle_delete.html", Post{Name: table_name, Rows: srows})
      if err500(w, err) {return}
    }
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
  fmt.Println("handleDelete")
  table_name := strings.Split(r.URL.Path, "/")[1]
  tid,_ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
  id := int32(tid)
  query := fmt.Sprintf("DELETE FROM %s WHERE id = %d", table_name, id)
  _, err := db.Query(query)
  if err500(w, err) {return}
  http.Redirect(w, r, fmt.Sprintf("/%s/read", table_name), http.StatusFound)
}

func main(){
    dbname := "dnd"
    uname := "gopher"

    fmt.Println("Starting dnd.main")

    //structs that represent database tables in ./db_structs.go

    //connect to postgres db
    db, err = sql.Open("postgres",
        fmt.Sprintf("dbname=%s user=%s password=gopher, sslmode=disable", dbname, uname)) //couldn't ping on arch with ssl mode enabled
    if err != nil {
        panic("Couldn't connect to database.")
    }
    if err = db.Ping(); err != nil {
        panic(err.Error())
    } else {
        fmt.Printf("%s connected to database: %s.\n", uname, dbname)
    }

    //custom multiplexer (cuts away trailing '/')
    mux := my_mux.Regex_mux{}

    //CRUD: create, read, update, delete
    regIndex,_ := regexp.Compile(`^(\/index)?$`)
    mux.HandleFunc(regIndex, handleIndex)

    regCreate,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/create$`)
    mux.HandleFunc(regCreate, handleCreate)

    regRead,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/read$`)
    mux.HandleFunc(regRead, handleRead)

    regToggleUpdate,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/toggle_update$`)
    mux.HandleFunc(regToggleUpdate, handleToggleUpdate)

    regToggleDelete,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/toggle_delete$`)
    mux.HandleFunc(regToggleDelete, handleToggleDelete)

    regDelete,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/[0-9]*\/delete$`)
    mux.HandleFunc(regDelete, handleDelete)

    regUpdate,_ := regexp.Compile(`^\/[a-zA-Z_0-9]+\/[0-9]*\/update$`)
    mux.HandleFunc(regUpdate, handleUpdate)

    http.ListenAndServe(":8001", mux)
}
