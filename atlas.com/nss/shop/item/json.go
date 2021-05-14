package item

type JSONObject struct {
	ShopId   uint32 `json:"shop_id"`
	ItemId   uint32 `json:"item_id"`
	Price    uint32 `json:"price"`
	Pitch    uint32 `json:"pitch"`
	Position uint32 `json:"position"`
}