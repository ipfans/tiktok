package tiktok

import (
	"context"
	"net/url"
)

func (c *Client) GetAuthorizedShop(ctx context.Context, ak string, shop_id string) (shops []Shop, err error) {
	param := url.Values{}
	param.Set("access_token", ak)
	if shop_id != "" {
		param.Set("shop_id", shop_id)
	}
	var resp ShopList
	err = c.Get(ctx, "/api/shop/get_authorized_shop", param, &resp)
	if err == nil {
		shops = resp.Shops
	}
	return
}
