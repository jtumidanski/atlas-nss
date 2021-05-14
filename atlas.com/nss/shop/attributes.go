package shop

type DataListContainer struct {
	Data []DataBody `json:"data"`
}

type DataContainer struct {
	Data DataBody `json:"data"`
}

type DataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	NPC   uint32           `json:"npc"`
	Items []ItemAttributes `json:"items"`
}

type ItemAttributes struct {
	ItemId   uint32 `json:"itemId"`
	Price    uint32 `json:"price"`
	Pitch    uint32 `json:"pitch"`
	Position uint32 `json:"position"`
}
