package main

type Npc struct {
    Id uint32
    Name string
}

type Race struct {
    Id uint32
    Name string
}

func init_db_structs() map[string]interface{} {
    m = make(map[string]interface{})
    m["npc"] = Npc{}
    m["race"] = Race{}
    return m
}


