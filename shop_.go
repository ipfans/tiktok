package tiktok

import (
	"encoding/json"
)

type Shop struct {
	ShopID   string      `json:"shop_id"`
	ShopName string      `json:"shop_name"`
	Region   string      `json:"region"`
	Type     json.Number `json:"type"`
}

type ShopList struct {
	Shops []Shop `json:"shop_list"`
}
