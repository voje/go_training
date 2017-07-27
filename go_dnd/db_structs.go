package go_dnd

func init_db_structs(m *map[string]interface{}) {
    m["npc"] = npc{}
    m["race"] = race{}
}

type npc struct {
    id uint32
    name string
}

type race struct {
    id uint32
    name string
}
