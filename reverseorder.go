package tiktok

import (
	"context"
	"net/url"
	"strconv"
)

func (c *Client) ConfirmReverse(ctx context.Context, p Param, orderID string) (err error) {
	var params url.Values
	if params, err = c.params(p); err != nil {
		return
	}
	m := map[string]string{
		"reverse_order_id": orderID,
	}
	err = c.Post(ctx, "/api/reverse/reverse_request/confirm", params, m, nil)
	return
}

func (c *Client) RejectReverse(ctx context.Context, p Param, query RejectReverseRequest) (err error) {
	var params url.Values
	if params, err = c.params(p); err != nil {
		return
	}
	if err = c.validate.Struct(&query); err != nil {
		return
	}
	err = c.Post(ctx, "/api/reverse/reverse_request/reject", params, query, nil)
	return
}

func (c *Client) GetReverseList(ctx context.Context, p Param, query GetReverseListRequest) (list ReverseOrdersList, err error) {
	var params url.Values
	if params, err = c.params(p); err != nil {
		return
	}
	if err = c.validate.Struct(&query); err != nil {
		return
	}
	err = c.Post(ctx, "/api/reverse/reverse_order/list", params, query, &list)
	return
}

func (c *Client) GetReverseReason(ctx context.Context, p Param, query GetReverseReasonRequest) (list ReverseReasonList, err error) {
	var params url.Values
	if params, err = c.params(p); err != nil {
		return
	}
	if err = c.validate.Struct(&query); err != nil {
		return
	}
	if query.ReverseActionType != 0 {
		params.Set("reverse_action_type", strconv.Itoa(query.ReverseActionType))
	}
	if query.ReasonType != 0 {
		params.Set("reason_type​", strconv.Itoa(query.ReasonType))
	}

	err = c.Get(ctx, "/api/reverse/reverse_reason/list", params, &list)
	return
}
