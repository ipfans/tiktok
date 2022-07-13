package tiktok

import (
	"context"
	"encoding/base64"
	"io"
	"net/url"
	"strings"
)

// GetCategory
// WARN: DO NOT CACHE THIS FUNCTION RESULT.
func (c *Client) GetCategory(ctx context.Context, p Param) (list CategoryList, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	err = c.Get(ctx, "/api/products/categories", param, &list)
	return
}

func (c *Client) GetAttribute(ctx context.Context, p Param, categoryID string) (list AttributeList, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	param.Set("category_id", categoryID)
	err = c.Get(ctx, "/api/products/attributes", param, &list)
	return
}

func (c *Client) GetCategoryRule(ctx context.Context, p Param, categoryID string) (list CategoryRules, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	param.Set("category_id", categoryID)
	err = c.Get(ctx, "/api/products/categories/rules", param, &list)
	return
}

func (c *Client) GetBrand(ctx context.Context, p Param, categoryID string) (list BrandList, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	if categoryID != "" {
		param.Set("category_id", categoryID)
	}
	err = c.Get(ctx, "/api/products/brands", param, &list)
	return
}

func (c *Client) UploadImgReader(ctx context.Context, p Param, scene ImgScene, r io.Reader) (img ImageInfo, err error) {
	var b []byte
	if b, err = io.ReadAll(r); err != nil {
		return
	}
	body := base64.StdEncoding.EncodeToString(b)
	img, err = c.UploadImg(ctx, p, scene, body)
	return
}

func (c *Client) UploadImg(ctx context.Context, p Param, scene ImgScene, body string) (img ImageInfo, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	m := map[string]interface{}{
		"img_data":  body,
		"img_scene": int(scene),
	}
	err = c.Post(ctx, "/api/products/upload_imgs", param, m, &img)
	return
}

func (c *Client) UploadFile(ctx context.Context, p Param, fileName string, body []byte) (info FileInfo, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	suffix := ".pdf"
	if !strings.HasSuffix(fileName, suffix) {
		fileName = fileName + suffix
	}
	m := map[string]interface{}{
		"file_name": fileName,
		"file_data": base64.StdEncoding.EncodeToString(body),
	}
	err = c.Post(ctx, "/api/products/upload_files", param, m, &info)
	return
}

func (c *Client) CreateProduct(ctx context.Context, p Param, req CreateProductRequest) (product ProductData, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	if err = c.validate.Struct(&req); err != nil {
		return
	}
	err = c.Post(ctx, "/api/products", param, req, &product)
	return
}

func (c *Client) EditProduct(ctx context.Context, p Param, req EditProductRequest) (product ProductData, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}

	if err = c.validate.Struct(&req); err != nil {
		return
	}
	err = c.Put(ctx, "/api/products", param, req, &product)
	return
}

func (c *Client) GetProductList(ctx context.Context, p Param, req ProductSearchRequest) (list GetProductListData, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	err = c.Post(ctx, "/api/products/search", param, req, &list)
	return
}

func (c *Client) GetProductDetail(ctx context.Context, p Param, productID string) (product ProductData, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	param.Set("product_id", productID)
	err = c.Get(ctx, "/api/products/details", param, &product)
	return
}

func (c *Client) UpdatePrice(ctx context.Context, p Param, req UpdatePriceRequest) (list UpdatePriceData, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	err = c.Put(ctx, "/api/products/prices", param, req, &list)
	return
}

func (c *Client) UpdateStock(ctx context.Context, p Param, req UpdateStockRequest) (list UpdateStockFailedSKU, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	err = c.Put(ctx, "/api/products/stocks", param, req, &list)
	return
}

func (c *Client) DeactivateProducts(ctx context.Context, p Param, productIDs []string) (list FailedProductIDs, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	req := map[string][]string{
		"product_ids": productIDs,
	}
	err = c.Post(ctx, "/api/products/inactivated_products", param, req, &list)
	return
}

func (c *Client) DeleteProducts(ctx context.Context, p Param, ids []string) (list FailedProductIDs, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	req := map[string][]string{
		"product_ids": ids,
	}
	err = c.Delete(ctx, "/api/products", param, req, &list)
	return
}

func (c *Client) RecoverProduct(ctx context.Context, p Param, ids []string) (list FailedProductIDs, err error) {
	param := url.Values{}
	param.Set("access_token", p.AccessToken)
	req := map[string][]string{
		"product_ids": ids,
	}
	err = c.Post(ctx, "/api/products/recover", param, req, &list)
	return
}

func (c *Client) ActivateProduct(ctx context.Context, p Param, ids []string) (list FailedProductIDs, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	req := map[string][]string{
		"product_ids": ids,
	}
	err = c.Post(ctx, "/api/products/activate", param, req, &list)
	return
}
