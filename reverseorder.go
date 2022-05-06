package tiktok

import (
	"context"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	_ConfirmReversePATH   = "/api/reverse/reverse_request/confirm"
	_RejectReversePATH    = "/api/reverse/reverse_request/reject"
	_GetReverseListPATH   = "/api/reverse/reverse_order/list"
	_GetReverseReasonPATH = "/api/reverse/reverse_reason/list"
)

func (c *Client) ConfirmReverse(ctx context.Context, p Param, orderID string) (err error) {
	var params url.Values
	if params, err = c.params(p); err != nil {
		return
	}
	m := map[string]string{
		"reverse_order_id": orderID,
	}
	err = c.Post(ctx, _ConfirmReversePATH, params, m, nil)
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
	err = c.Post(ctx, _RejectReversePATH, params, query, nil)
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
	err = c.Post(ctx, _GetReverseListPATH, params, query, &list)
	return
}

func (c *Client) GetReverseReason(ctx context.Context, p Param, req GetReverseReasonRequest) (list ReverseReasonList, err error) {
	var params url.Values
	if params, err = c.params(p); err != nil {
		return
	}
	if err = c.validate.Struct(&req); err != nil {
		return
	}

	values, err := query.Values(req)
	if err != nil {
		return
	}
	for k, data := range values {
		if items, ok := params[k]; ok {
			params[k] = append(items, data...)
		} else {
			params[k] = data
		}
	}

	err = c.Get(ctx, _GetReverseReasonPATH, params, &list)
	return
}
