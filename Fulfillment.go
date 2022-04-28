package tiktok

import (
	"context"

	"github.com/google/go-querystring/query"
)

const (
	SearchPreCombinePkgPATH        = "/api/fulfillment/pre_combine_pkg/list"
	ConfirmPreCombinePkgPATH       = "/api/fulfillment/pre_combine_pkg/confirm"
	RemovePackageOrderPATH         = "/api/fulfillment/package/remove"
	GetPackagePickupConfigPATH     = "/api/fulfillment/package_pickup_config/list"
	ShipPackagePATH                = "/api/fulfillment/rts"
	SearchPackagePATH              = "/api/fulfillment/search"
	GetPackageDetailPATH           = "/api/fulfillment/detail"
	GetPackageShippingInfoPATH     = "/api/fulfillment/shipping_info"
	UpdatePackageShippingInfoPATH  = "/api/fulfillment/shipping_info/update"
	GetPackageShippingDocumentPATH = "/api/fulfillment/shipping_document"
	VerifyOrderSplitPATH           = "/api/fulfillment/order_split/verify"
	ConfirmOrderSplitPATH          = "/api/fulfillment/order_split/confirm"
)

func (c *Client) SearchPreCombinePkg(ctx context.Context, p Param, req SearchPreCombinePkgRequest) (data SearchPreCombinePkgData, err error) {
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

	err = c.Get(ctx, SearchPreCombinePkgPATH, param, &data)
	return
}

func (c *Client) ConfirmPreCombinePkg(ctx context.Context, p Param, req ConfirmPreCombinePkgRequest) (data PreCombinePkgData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, ConfirmPreCombinePkgPATH, param, req, &data)
	return
}

func (c *Client) RemovePackageOrder(ctx context.Context, p Param, req RemovePackageOrderRequest) (data PreCombinePkgData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, RemovePackageOrderPATH, param, req, &data)
	return
}

func (c *Client) GetPackagePickupConfig(ctx context.Context, p Param, req PackageIDRequest) (data GetPackagePickupConfigData, err error) {
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
	err = c.Get(ctx, GetPackagePickupConfigPATH, param, &data)
	return
}

func (c *Client) ShipPackage(ctx context.Context, p Param, req ShipPackageRequest) (data ShipPackageData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, ShipPackagePATH, param, req, &data)
	return
}

func (c *Client) SearchPackage(ctx context.Context, p Param, req SearchPackageRequest) (data SearchPackageData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, SearchPackagePATH, param, req, &data)
	return
}

func (c *Client) GetPackageDetail(ctx context.Context, p Param, req PackageIDRequest) (data GetPackageDetailData, err error) {
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
	err = c.Get(ctx, GetPackageDetailPATH, param, &data)
	return
}

func (c *Client) GetPackageShippingInfo(ctx context.Context, p Param, req PackageIDRequest) (data GetPackageShippingInfoData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	param.Set("package_id", req.PackageID)
	err = c.Get(ctx, GetPackageShippingInfoPATH, param, &data)
	return
}

func (c *Client) UpdatePackageShippingInfo(ctx context.Context, p Param, req UpdatePackageShippingInfoRequest) (data UpdatePackageShippingInfoData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, UpdatePackageShippingInfoPATH, param, req, &data)
	return
}

func (c *Client) GetPackageShippingDocument(ctx context.Context, p Param, req GetPackageShippingDocumentRequest) (data GetPackageShippingDocumentData, err error) {
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
	err = c.Get(ctx, GetPackageShippingDocumentPATH, param, &data)
	return
}

func (c *Client) VerifyOrderSplit(ctx context.Context, p Param, req VerifyOrderSplitRequest) (data VerifyOrderSplitData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, VerifyOrderSplitPATH, param, req, &data)
	return
}

func (c *Client) ConfirmOrderSplit(ctx context.Context, p Param, req ConfirmOrderSplitRequest) (data ConfirmOrderSplitData, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(req); err != nil {
		return
	}

	err = c.Post(ctx, ConfirmOrderSplitPATH, param, req, &data)
	return
}
