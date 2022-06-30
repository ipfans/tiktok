package tiktok

import (
	"context"
)

// GetOrderList get order list. Default PageSize is 50.
func (c *Client) GetOrderList(ctx context.Context, p Param, query GetOrderListRequest) (list OrdersList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if query.PageSize <= 0 || query.PageSize > 50 {
		query.PageSize = 50
	}
	err = c.Post(ctx, "/api/orders/search", param, query, &list)
	return
}

// GetOrderDetail get order detail.
func (c *Client) GetOrderDetail(ctx context.Context, p Param, ordersID []string) (list OrderDetailList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	var req GetOrderDetailRequest
	req.OrderIDList = ordersID
	err = c.Post(ctx, "/api/orders/detail/query", param, req, &list)
	return
}

// ShipOrder ship order.
func (c *Client) ShipOrder(ctx context.Context, p Param, req ShipOrderRequest) (code int, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}
	err = c.Post(ctx, "/api/order/rts", param, req, nil)
	if err != nil {
		code = ErrCode(err)
	}
	return
}

// CancelOrder to cancel the order.
func (c *Client) CancelOrder(ctx context.Context, p Param, req CancelOrderRequest) (resp CancelOrderResponse, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}
	err = c.Post(ctx, "/api/reverse/order/cancel", param, req, &resp)
	return
}
