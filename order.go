package tiktok

import (
	"context"
	"net/url"
)

func (c *Client) GetOrderList(ctx context.Context, ak, shopID string, query GetOrderListRequest) (list OrdersList, err error) {
	param := url.Values{}
	param.Set("access_token", ak)
	param.Set("shop_id", shopID)
	if query.PageSize == 0 {
		query.PageSize = 50
	}
	err = c.Post(ctx, "/api/orders/search", param, query, &list)
	return
}

func (c *Client) GetOrderDetail(ctx context.Context, ak, shopID string, ordersID []string) (list OrderDetail, err error) {
	param := url.Values{}
	param.Set("access_token", ak)
	if shopID != "" {
		param.Set("shop_id", shopID)
	}
	var req GetOrderDetailRequest
	req.OrderIDList = ordersID
	err = c.Post(ctx, "/api/orders/detail/query", param, req, &list)
	return
}
