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

func new_db_struct(s string) TableRow {
  switch s {
    case "npc":
      return new(Npc)
    case "race":
      return new(Race)
  }
  return nil
}


