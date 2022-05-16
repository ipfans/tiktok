package api

import (
	"github.com/gin-gonic/gin"
)

var s = NewClient()

// GetCategoryTree is get category tree or dictionary from marketplace
// @Summary get category tree or dictionary from marketplace
// @Description Lazada: https://bit.ly/3wbFq64 shopee: https://bit.ly/3l7cjKT Tokopedia:  https://bit.ly/39fr2Rj Tiktok: https://bit.ly/3FI0HaM
// @Description get category tree or dictionary from marketplace
// @Description e.g. get category tree, like telecommunication->mobile phone-> iphone -> iphone 13 series
// @Tags category
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Success 200 {object} HTTPGetCategoryTreeResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/category [GET]
func GetCategoryTree(ctx *gin.Context) {
}

// GetBrandList is get brand list from marketplace
// @Summary get brand list from marketplace
// @Description Lazada: https://bit.ly/3l8loTF shopee: https://bit.ly/3L7JMQ5 Tokopedia: (NONE,not found brand api) Tiktok: https://bit.ly/3FO6Pym
// @Description need to brand_id  mapping to marketplace brand id(TODO)
// @Description support pagination query
// @Description get brand list from marketplace e.g. get brand list like Apple Google Facebook Airbnb
// @Tags brand
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param page_no query integer true "common query params: page number"
// @Param page_size query integer true "common query params: page size"
// @Success 200 {object} HTTPGetBrandListResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/brand [GET]
func GetBrandList(ctx *gin.Context) {
}

// GetCategoryAttrs is get product attributes from marketplace, in other words, category attributes
// @Summary get product attributes from marketplace, in other words, category attributes
// @Description Lazada: https://bit.ly/3Ncd0yp  shopee: https://bit.ly/3PxuYOb  Tokopedia: https://bit.ly/3wkdMCQ Tiktok: https://bit.ly/3wtQdrm
// @Description Need to category_id mapping to marketplace category id(TODO)
// @Description Get product attributes from marketplace e.g. under category_id 100, product attributes is waterproof fireproof isDangerous R18
// @Tags category
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param category_id  query integer true "category_id"
// @Success 200 {object} HTTPGetProductAttrsResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/category/attribute [GET]
func GetCategoryAttrs(ctx *gin.Context) {
	categoryID := 1
	mp := 1
	shopID := 1
	s.GetProductAttrs(ctx, mp, shopID, categoryID)
	// TODO need to category_id mapping to marketplace category id
}

// CreateProduct is create product on multiple marketplaces
// @Summary create product on multiple marketplaces, how to create a product refer link: lazada usecase-> https://bit.ly/3yym6BG
// @Description so we need do shipper standard category id and brand id mapping , many attributes to special marketplace category id, brand id, and product attributes
// @Description Lazada: https://bit.ly/3PwiIgO shopee: https://bit.ly/3LjQimY Tokopedia: https://bit.ly/3wqKD9c  Tiktok: https://bit.ly/3NaZiMr
// @Description Need to be design seriously, talked and review, especially for request body
// @Description create product on multiple marketplaces, e.g. create iphone13 256gb on tiktok shop.
// @Tags product
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param data body AddEditProductReq true "Create Product Data"
// @Success 200 {object} HTTPCreateProductResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product [POST]
func CreateProduct(ctx *gin.Context) {
	req := AddEditProductReq{}
	// brandID categoryID
	s.CreateProduct(ctx, req)
}

// EditProductAttributes is edit product attributes
// @Summary is edit product attributes
// @Description a bit different like createProduct API, we query product attributes from specific marketplace when we create product, product info stored in our database
// @Description lazada: https://bit.ly/3ld8CU5 shopee: https://bit.ly/3lg4vX7 tokopedia: https://bit.ly/3FJFEV4 tiktok: https://bit.ly/3lb5VCs
// @Description support single edie, bulk edit not supported
// @Description edit product attributes, e.g. package weight,width,height,is_dangerous,
// @Tags product
// @Accept json
// @Produce json
// @Param  data body AddEditProductReq true "Update Product attributes"
// @Success 200 {object} HTTPUpdateProductResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/edit [PATCH]
func EditProductAttributes(ctx *gin.Context) {
	req := AddEditProductReq{}
	s.UpdateProductAttributes(ctx, req)
}

// UpdateProductPrice is update product price
// @Summary update product original price
// @Description lazada: https://bit.ly/3yBQVVV shopee: https://bit.ly/3Pn4HSe tokopedia: https://bit.ly/3PoACBR tiktok: https://bit.ly/3szsTHt
// @Description update product original price, e.g. old price is 1000Rp, new price is 668Rp
// @Tags product
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param  data body UpdateProductPriceRequest true "Update Product Price"
// @Success 200 {object} HTTPEmptyResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/price [PUT]
func UpdateProductPrice(ctx *gin.Context) {
	req := UpdateProductPriceRequest{}
	s.UpdateProductPrice(ctx, req)
}

// UpdateProductStock is update product stock at specific warehouse
// @Summary update product stock at specific warehouse
// @Description lazada: https://bit.ly/3yBQVVV shopee: https://bit.ly/3wrusIL tokopedia: https://bit.ly/3NcVf1Z tiktok: https://bit.ly/3sAyf57
// @Description update product stock at specific warehouse, e.g. update stock from 1 to 100 at specific warehouse
// @Tags product
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param  data body UpdateProductStockRequest true "Update Product Stock"
// @Success 200 {object} HTTPEmptyResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/stock [PUT]
func UpdateProductStock(ctx *gin.Context) {
	req := UpdateProductStockRequest{}
	s.UpdateProductStock(ctx, req)
}

// DeleteProduct is delete product from product list
// @Summary delete product from product list
// @Description  tokopedia: https://bit.ly/38fpR4n, lazada: https://bit.ly/3yCt7l4 shopee: https://bit.ly/3wtMMAY tiktok: https://bit.ly/38qiuHe
// @Description delete product from product list, e.g. delete iphone 13 red 256GB from shop product list
// @Tags product
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param  data body ProductIDSRequest true "Delete Product From Shelf"
// @Success 200 {object} HTTPEmptyResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product [delete]
func DeleteProduct(ctx *gin.Context) {
	req := ProductIDSRequest{}
	s.DeleteProduct(ctx, req)
}

// QueryProductInfoList is fetch product id list
// @Summary fetch product info list,⚠️ dut to  tokopedia GetProduct api design style, https://bit.ly/3PdbI8e ,把普通分页查询和指定一批ProductIDList seller_sku_list 设计.
// @Summary 返回结果 参考现有 Atoor: https://dashboard.staging.atoor.com/ocms/product 返回数据 设计, 还会结合当前做一定改动,删除一些不能获取数据的字段!
// @Description Lazada: 批量查询 https://bit.ly/3MfPaBQ 单条查询 https://bit.ly/39mmw3i 数据都是一模一样得,
// @Description shopee: https://bit.ly/3syQb0h get_item_list->get_item_base_info->get_item_extra_info
// @Description Tokopedia: https://bit.ly/3PdbI8e Tiktok: https://bit.ly/3l7AkBy
// @Description fetch product id list, e.g. get product list ["123","345","456"]
// @Tags product
// @Accept json
// @Produce json
// @Param  data query QueryProductInfoRequest true "Query Product params"
// @Success 200 {object} HTTPProductListResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/info [GET]
func QueryProductInfoList(ctx *gin.Context) {
	req := QueryProductInfoRequest{}
	s.QueryProductInfoList(ctx, req)
}

// UpDown is activate/deactivate products on/from shelf, in other words, list/unlist product
// @Summary activate/deactivate products on/from shelf, in other words, list/unlist product
// @Description Lazada: https://bit.ly/37Ix6kR shopee: https://bit.ly/3wbm6G6 Tokopedia: https://bit.ly/3wkDsiH Tiktok: https://bit.ly/3sxkOTu
// @Description activate/deactivate products on/from shelf, in other words, list/unlist product. e.g. iphone 13 red 256GB is top seller goods. so we often deactivate it for the sake of shortage
// @Tags product
// @Accept json
// @Produce json
// @Param shop_id query integer true "common query params: shop_id"
// @Param mp query integer true "common query params: marketplace type"
// @Param  data body UpDownShelfReq true "deactivate Product From Shelf"
// @Success 200 {object} HTTPEmptyResp
// @Failure 400 {object} HTTPErrResp
// @Failure 401 {object} HTTPErrResp
// @Failure 403 {object} HTTPErrResp
// @Failure 422 {object} HTTPErrResp
// @Failure 429 {object} HTTPErrResp
// @Failure 500 {object} HTTPErrResp
// @Failure 503 {object} HTTPErrResp
// @Router /v1/product/shelf/up-down [PUT]
func UpDown(ctx *gin.Context) {
	req := UpDownShelfReq{}
	s.UpDown(ctx, req)
}
