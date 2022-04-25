package tiktok

import (
	"context"
	"encoding/base64"
	"io"
	"net/url"
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

func (c *Client) GetBrand(ctx context.Context, p Param) (list BrandList, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
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

func (c *Client) UploadFile(ctx context.Context, p Param, fn string, body []byte) (info FileInfo, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	m := map[string]interface{}{
		"file_name": fn,
		"file_data": base64.StdEncoding.EncodeToString(body),
	}
	err = c.Post(ctx, "/api/products/upload_files", param, m, &info)
	return
}

func (c *Client) CreateProduct(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) EditProduct(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) GetProductList(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) GetProductDetail(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) UpdatePrice(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) UpdateStock(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) DeactivateProducts(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) DeleteProducts(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) RecoverProduct(ctx context.Context, p Param) (err error) {
	return
}

func (c *Client) ActivateProduct(ctx context.Context, p Param) (err error) {
	return
}
