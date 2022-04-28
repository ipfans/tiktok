package tiktok

import (
	"context"

	"github.com/google/go-querystring/query"
)

const (
	GetShippingInfoPATH       = "/api/logistics/ship/get"
	UpdateShippingInfoPATH    = "/api/logistics/tracking"
	GetShippingDocumentPATH   = "/api/logistics/shipping_document"
	GetWarehouseListPATH      = "/api/logistics/get_warehouse_list"
	GetShippingProviderPATH   = "/api/logistics/shipping_providers"
	GetSubscribedDeliveryPATH = "/api/logistics/get_subscribed_deliveryoptions"
)

func (c *Client) GetShippingInfo(ctx context.Context, p Param, req OrderIDReq) (data LogisticsGetShippingInfoData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	values, err := query.Values(req)
	if err != nil {
		return
	}
	for k, data := range values {
		if items, ok := param[k]; ok {
			param[k] = append(items, data...)
		} else {
			param[k] = data
		}
	}

	err = c.Get(ctx, GetShippingInfoPATH, param, &data)
	return
}

func (c *Client) UpdateShippingInfo(ctx context.Context, p Param, req UpdateShippingInfoReq) (err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, UpdateShippingInfoPATH, param, req, nil)
	return
}

func (c *Client) GetShippingDocument(ctx context.Context, p Param, req GetShippingDocumentRequest) (data GetShippingDocumentData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	values, err := query.Values(req)
	if err != nil {
		return
	}
	for k, data := range values {
		if items, ok := param[k]; ok {
			param[k] = append(items, data...)
		} else {
			param[k] = data
		}
	}

	err = c.Get(ctx, GetShippingDocumentPATH, param, &data)
	return
}

func (c *Client) GetWarehouseList(ctx context.Context, p Param) (data GetWarehouseListData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	err = c.Get(ctx, GetWarehouseListPATH, param, &data)
	return
}

func (c *Client) GetShippingProvider(ctx context.Context, p Param) (data GetShippingProviderData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	err = c.Get(ctx, GetShippingProviderPATH, param, &data)
	return
}

func (c *Client) GetSubscribedDelivery(ctx context.Context, p Param, req GetSubscribedDeliveryRequest) (data GetSubscribedDeliveryData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, GetSubscribedDeliveryPATH, param, req, &data)
	return
}
