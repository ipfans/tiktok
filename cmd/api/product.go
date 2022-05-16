package api

import (
	"context"
)

type ProductUseCase interface{}

func (s *Client) GetCategoryTree(ctx context.Context, mp int, shopID int) {
	// daily update from marketplace open api, stored in db.
	// query from database
}
func (s *Client) GetBrandList(ctx context.Context, mp int)                              {}
func (s *Client) GetProductAttrs(ctx context.Context, mp int, shopID, categoryID int)   {}
func (s *Client) CreateProduct(ctx context.Context, req AddEditProductReq)              {}
func (s *Client) UpdateProductAttributes(ctx context.Context, req AddEditProductReq)    {}
func (s *Client) UpdateProductPrice(ctx context.Context, req UpdateProductPriceRequest) {}
func (s *Client) UpdateProductStock(ctx context.Context, req UpdateProductStockRequest) {}

func (s *Client) DeleteProduct(ctx context.Context, req ProductIDSRequest)              {}
func (s *Client) RecoverProduct(ctx context.Context, req ProductIDSRequest)             {}
func (s *Client) QueryProductInfoList(ctx context.Context, req QueryProductInfoRequest) {}

func (s *Client) ActivateProduct(ctx context.Context, req ActivateProductRequest) {}
func (s *Client) UpDown(ctx context.Context, req UpDownShelfReq)                  {}
