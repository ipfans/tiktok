package tiktok

import "net/url"

func (c *Client) GetAuthorizedShop(ak string, shop_id string) (shops []Shop, err error) {
	param := url.Values{}
	param.Set("access_token", ak)
	if shop_id != "" {
		param.Set("shop_id", shop_id)
	}
	var resp ShopList
	err = c.Get("/api/shop/get_authorized_shop", param, &resp)
	if err == nil {
		shops = resp.Shops
	}
	return
}
