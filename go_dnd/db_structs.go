package main

import (
  "database/sql"
  //_ "github.com/lib/pq"
)

type TableRow interface {
  Scan(rows *sql.Rows) error
}

type Npc struct {
    Id uint32
    Name string
}

func (n *Npc) Scan (rows *sql.Rows) error {
  return rows.Scan(&n.Id, &n.Name)
}

type Race struct {
    Id uint32
    Name string
}

func (r *Race) Scan (rows *sql.Rows) error {
  return rows.Scan(&r.Id, &r.Name)
}

func init_db_structs() map[string]TableRow {
  m := make(map[string]TableRow)
  //interface TableRow has methods for pointers,
  //therefore TableRow is of pointer type
  m["npc"] = &Npc{}
  m["race"] = &Race{}
  return m
}


