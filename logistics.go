package tiktok

import (
	"context"
)

const (
	GetShippingInfoPATH     = "/api/logistics/ship/get"
	UpdateShippingInfoPATH  = "/api/logistics/tracking"
	GetShippingDocumentPATH = "/api/logistics/shipping_document"
	GetAddressListPATH      = "/api/logistics/addresses"
	GetWarehouseListPATH    = "/api/logistics/get_warehouse_list"
	GetShippingProviderPATH = "/api/logistics/shipping_providers"
)

func (c *Client) GetShippingInfo(ctx context.Context, p Param, req OrderIDReq) (data LogisticsGetShippingInfoData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	param.Set("order_id", req.OrderID)
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

	param.Set("order_id", req.OrderID)
	param.Set("document_type", req.DocumentType)
	if req.DocumentSize != "" {
		param.Set("document_size", req.DocumentSize)
	}

	err = c.Get(ctx, GetShippingDocumentPATH, param, &data)
	return
}

func (c *Client) GetAddressList(ctx context.Context, p Param, req GetAddressListRequest) (data GetAddressListData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	if req.AddressType != "" {
		param.Set("address_type", req.AddressType)
	}
	err = c.Get(ctx, GetAddressListPATH, param, &data)
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
