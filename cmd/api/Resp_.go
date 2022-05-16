package api

import (
	"time"
)

type HTTPErrResp struct {
	Meta Meta `json:"metadata"`
}

type HTTPEmptyResp struct {
	Meta Meta `json:"metadata"`
}

type Meta struct {
	Path       string    `json:"path"`
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	Error      *AppError `json:"error,omitempty" swaggertype:"primitive,integer"`
	Timestamp  string    `json:"timestamp"`
}

type ErrorCode uint32

// AppError - Application Error Structure
type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	DebugError *string   `json:"debug,omitempty"`
	sys        error
}

type ShipperCategoryTreeData struct {
	ID          string `json:"id"`
	Leaf        bool   `json:"is_leaf"`
	DisplayName string `json:"display_name"`
	ParentID    string `json:"parent_id"`
}

type HTTPGetCategoryTreeResp struct {
	Meta Meta                    `json:"metadata"`
	Data ShipperCategoryTreeData `json:"data"`
}

type Brand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetBrandListData struct {
	PageResult
	BrandList []Brand `json:"brand_list"`
}

type HTTPGetBrandListResp struct {
	Meta Meta             `json:"metadata"`
	Data GetBrandListData `json:"data"`
}

type Value struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type InputType struct {
	IsMandatory        bool `json:"is_mandatory"`
	IsMultipleSelected bool `json:"is_multiple_selected"`
	IsCustomized       bool `json:"is_customized"`
}

type Attribute struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Label         string    `json:"label"`
	AttributeType int       `json:"attribute_type"`
	InputType     InputType `json:"input_type"`
	Values        []Value   `json:"values"`
}

type GetProductAttrsData struct {
	Attributes []Attribute `json:"attributes"`
}

type HTTPGetProductAttrsResp struct {
	Meta Meta                `json:"metadata"`
	Data GetProductAttrsData `json:"data"`
}

type BrandName2BrandData struct {
	BrandList []Brand `json:"brand_list"`
}

type HTTPBrandName2BrandResp struct {
	Meta Meta                `json:"metadata"`
	Data BrandName2BrandData `json:"data"`
}

type BrandID2MPBrandIDData struct {
	Brand Brand `json:"brand"`
}

type HTTPBrandID2MPBrandIDResp struct {
	Meta Meta                  `json:"metadata"`
	Data BrandID2MPBrandIDData `json:"data"`
}

type ShipperAttrs2MPAttrsData struct {
	TODO string `json:"todo"`
}

type HTTPShipperAttrs2MPAttrsResp struct {
	Meta Meta                     `json:"metadata"`
	Data ShipperAttrs2MPAttrsData `json:"data"`
}

type ShipperLogisticsProviderData struct {
	ID         int    `json:"id"`
	LogisticNo string `json:"logistic_no"`
	Name       string `json:"name"`
}

type HTTPShipperLogisticProviderCategoryResp struct {
	Meta Meta                         `json:"metadata"`
	Data ShipperLogisticsProviderData `json:"data"`
}

type ProductIDData struct {
	ProductID string `json:"product_id"`
}

type CreateProductData struct {
	TotalCount      int `json:"total_count"`
	SuccessCount    int `json:"success_count"`
	SuccessRowsData []struct {
		ProductID int64 `json:"product_id"`
	} `json:"success_rows_data"`
	FailRowsData []struct {
		RequestID  string `json:"request_id"`
		SellerSKU  string `json:"seller_sku"`
		FailReason string `json:"fail_reason"`
	} `json:"fail_rows_data"`
}

type HTTPCreateProductResp struct {
	Meta Meta              `json:"metadata"`
	Data CreateProductData `json:"data"`
}

type HTTPUpdateProductResp struct {
	Meta Meta              `json:"metadata"`
	Data CreateProductData `json:"data"`
}

type HTTPProductListResp struct {
	Meta Meta                  `json:"metadata"`
	Data ProductPageResultData `json:"data"`
}

// PageArg page arg
type PageArg struct {
	Page int `form:"page_no"`
	Size int `form:"page_size"`
}

// CheckPageValidation check the page validate, return limit offset
func (arg *PageArg) CheckPageValidation() (limit, offset int) {
	if arg.Page < 1 {
		arg.Page = 1
	}
	if arg.Size > 1000 || arg.Size <= 0 {
		arg.Size = 10
	}
	limit = arg.Size
	offset = (arg.Page - 1) * limit
	return
}

// PageResult page result
type PageResult struct {
	Page       int `json:"page"`
	TotalCount int `json:"total_count"`
}

type ProductPageResultData struct {
	PageResult
	Products []struct {
		ProductItem struct {
			ProductId string `json:"productId"`
			// MerchantId           int       `json:"merchantId"`
			MarketplaceShopId    int       `json:"marketplaceShopId"`
			MarketplaceProductId int       `json:"marketplaceProductId"`
			Channel              string    `json:"channel"`
			CreatedAt            time.Time `json:"createdAt"`
			CreatedBy            int       `json:"createdBy"`
			UpdatedAt            time.Time `json:"updatedAt"`
			Status               string    `json:"status"`
			//	ReasonUpload         string    `json:"reasonUpload"`
			//	StatusUpload         string    `json:"statusUpload"`
			ShopName        string `json:"shopName"`
			ProductName     string `json:"productName"`
			ProductImage    string `json:"productImage"`
			InternalShopId  string `json:"internalShopId"`
			IsSingleProduct bool   `json:"isSingleProduct"` // one productID(spuID) only have one skuID
			Variants        []struct {
				MerchantId int `json:"merchantId"`
				//	VariantId                   string `json:"variantId"`
				ShopName                    string `json:"shopName"`
				Variation                   string `json:"variation"`
				SkuMaster                   string `json:"skuMaster"`
				SkuVariant                  string `json:"skuVariant"`
				ProductId                   string `json:"productId"`
				Price                       int    `json:"price"`
				AvailableStock              int    `json:"availableStock"`
				Stock                       int    `json:"stock"`
				MarketplaceProductVariantId int    `json:"marketplaceProductVariantId"`
				Channel                     string `json:"channel"`
				Status                      string `json:"status"`
			} `json:"variants"`
		} `json:"node"`
	} `json:"edges"`
}
