package tiktok

import "encoding/json"

type Response struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
}

type AccessTokenResponse struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpireIn  int    `json:"access_token_expire_in"`
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireIn int    `json:"refresh_token_expire_in"`
	OpenID               string `json:"open_id"`
	SellerName           string `json:"seller_name"`
}

type Shop struct {
	ShopID   string      `json:"shop_id"`
	ShopName string      `json:"shop_name"`
	Region   string      `json:"region"`
	Type     json.Number `json:"type"`
}

type ShopList struct {
	Shops []Shop `json:"shop_list"`
}

type Order struct {
	OrderID     string `json:"order_id"`
	OrderStatus int    `json:"order_status"`
	UpdateTime  int    `json:"update_time"`
}

type OrdersList struct {
	OrderList  []Order `json:"order_list,omitempty"`
	More       bool    `json:"more"`
	NextCursor string  `json:"next_cursor"`
	Total      int     `json:"total"`
}

type RecipientAddress struct {
	FullAddress     string   `json:"full_address"`
	Region          string   `json:"region"`
	State           string   `json:"state"`
	City            string   `json:"city"`
	District        string   `json:"district"`
	Town            string   `json:"town"`
	Phone           string   `json:"phone"`
	Name            string   `json:"name"`
	Zipcode         string   `json:"zipcode"`
	AddressDetail   string   `json:"address_detail"`
	AddressLineList []string `json:"address_line_list"`
}

type PaymentInfo struct {
	Currency                    string `json:"currency"`
	SubTotal                    int    `json:"sub_total"`
	ShippingFee                 int    `json:"shipping_fee"`
	SellerDiscount              int    `json:"seller_discount"`
	TotalAmount                 int    `json:"total_amount"`
	OriginalTotalProductPrice   int    `json:"original_total_product_price"`
	OriginalShippingFee         int    `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   int    `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount int    `json:"shipping_fee_platform_discount"`
}

type Item struct {
	SkuID            string `json:"sku_id"`
	ProductID        string `json:"product_id"`
	ProductName      string `json:"product_name"`
	SkuName          string `json:"sku_name"`
	SkuImage         string `json:"sku_image"`
	Quantity         int    `json:"quantity"`
	SellerSku        string `json:"seller_sku"`
	SkuOriginalPrice int    `json:"sku_original_price"`
	SkuSalePrice     int    `json:"sku_sale_price"`
}

type OrderDetail struct {
	OrderID                string           `json:"order_id"`
	OrderStatus            int              `json:"order_status"`
	PaymentMethod          string           `json:"payment_method"`
	DeliveryOption         string           `json:"delivery_option"`
	ShippingProvider       string           `json:"shipping_provider"`
	ShippingProviderID     string           `json:"shipping_provider_id"`
	CreateTime             string           `json:"create_time"`
	PaidTime               json.Number      `json:"paid_time"`
	BuyerMessage           string           `json:"buyer_message"`
	PaymentInfo            PaymentInfo      `json:"payment_info,omitempty"`
	RecipientAddress       RecipientAddress `json:"recipient_address,omitempty"`
	TrackingNumber         string           `json:"tracking_number"`
	ItemList               []Item           `json:"item_list,omitempty"`
	RtsTime                int              `json:"rts_time"`
	RtsSLA                 int              `json:"rts_sla"`
	TtsSLA                 int              `json:"tts_sla"`
	CancelOrderSLA         int              `json:"cancel_order_sla"`
	ReceiverAddressUpdated int              `json:"receiver_address_updated"`
	UpdateTime             int              `json:"update_time"`
	BuyerUID               string           `json:"buyer_uid"`
	FulfillmentType        int              `json:"fulfillment_type"`
	OrderLineList          []OrderLineList  `json:"order_line_list"`
}

type OrderLineList struct {
	OrderLineID string `json:"order_line_id"`
	SkuID       string `json:"sku_id"`
}

type OrderDetailList struct {
	OrderList []OrderDetail `json:"order_list"`
}

type SettlementInfo struct {
	SettlementTime      int    `json:"settlement_time"`
	Currency            string `json:"currency"`
	UserPay             string `json:"user_pay"`
	PlatformPromotion   string `json:"platform_promotion"`
	Refund              string `json:"refund"`
	PaymentFee          string `json:"payment_fee"`
	PlatformCommission  string `json:"platform_commission"`
	AffiliateCommission string `json:"affiliate_commission"`
	Vat                 string `json:"vat"`
	ShippingFee         string `json:"shipping_fee"`
	SettlementAmount    string `json:"settlement_amount"`
}

type Settlement struct {
	UniqueKey      int64          `json:"unique_key"`
	OrderID        string         `json:"order_id"`
	AdjustmentID   string         `json:"adjustment_id"`
	RelatedOrderID string         `json:"related_order_id"`
	SkuID          string         `json:"sku_id"`
	SkuName        string         `json:"sku_name"`
	ProductName    string         `json:"product_name"`
	SettlementInfo SettlementInfo `json:"settlement_info"`
}

type SettlementsList struct {
	More           bool         `json:"more"`
	NextCursor     string       `json:"next_cursor"`
	SettlementList []Settlement `json:"settlement_list"`
}

type OrdersIDRequest struct {
	OrderID string `json:"order_id" url:"order_id" validate:"required"`
}
type Transaction struct {
	TransactionType     int    `json:"transaction_type"`
	TransactionTime     int    `json:"transaction_time"`
	TransactionAmount   string `json:"transaction_amount"`
	TransactionCurrency string `json:"transaction_currency"`
	TransactionStatus   int    `json:"transaction_status"`
}

type TransactionsList struct {
	More            bool          `json:"more"`
	Total           int           `json:"total"`
	TransactionList []Transaction `json:"transaction_list"`
}

type ReturnItem struct {
	ReturnProductID   string `json:"return_product_id"`
	ReturnProductName string `json:"return_product_name"`
	SkuID             string `json:"sku_id"`
	SellerSku         string `json:"seller_sku"`
	SkuName           string `json:"sku_name"`
	ReturnQuantity    int    `json:"return_quantity"`
	ProductImages     string `json:"product_images"`
}

type Reverse struct {
	ReverseOrderID     string       `json:"reverse_order_id"`
	OrderID            string       `json:"order_id"`
	RefundTotal        string       `json:"refund_total"`
	Currency           string       `json:"currency"`
	ReverseType        int          `json:"reverse_type"`
	ReturnReason       string       `json:"return_reason"`
	ReturnItemList     []ReturnItem `json:"return_item_list"`
	ReverseStatusValue int          `json:"reverse_status_value"`
	ReverseRequestTime int          `json:"reverse_request_time"`
	ReverseUpdateTime  int          `json:"reverse_update_time"`
	ReturnTrackingID   string       `json:"return_tracking_id"`
}

type ReverseOrdersList struct {
	ReverseList []Reverse `json:"reverse_list"`
	More        bool      `json:"more"`
}

type AutoGenerated struct {
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	RequestID string   `json:"request_id"`
	Data      struct{} `json:"data"`
}

type ReverseReason struct {
	ReverseReasonKey         string `json:"reverse_reason_key"`
	ReverseReason            string `json:"reverse_reason"`
	AvailableOrderStatusList []int  `json:"available_order_status_list"`
}

type ReverseReasonList struct {
	ReverseReasonList []ReverseReason `json:"reverse_reason_list"`
}

type Category struct {
	ID               string `json:"id"`
	ParentID         string `json:"parent_id"`
	LocalDisplayName string `json:"local_display_name"`
	IsLeaf           bool   `json:"is_leaf"`
}
type CategoryList struct {
	CategoryList []Category `json:"category_list"`
}

type InputType struct {
	IsMandatory        bool `json:"is_mandatory"`
	IsMultipleSelected bool `json:"is_multiple_selected"`
	IsCustomized       bool `json:"is_customized"`
}

type Attribute struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	AttributeType int       `json:"attribute_type"`
	InputType     InputType `json:"input_type"`
}
type AttributeList struct {
	Attributes []Attribute `json:"attributes"`
}

type ProductCertification struct {
	Name        string `json:"certification_name"`
	ID          string `json:"certification_id"`
	Sample      string `json:"certification_sample"`
	IsMandatory bool   `json:"is_mandatory"`
}
type CategoryRule struct {
	ProductCertifications []ProductCertification `json:"product_certifications"`
	SupportSizeChart      bool                   `json:"support_size_chart"`
	SupportCod            bool                   `json:"support_cod"`
}

type CategoryRules struct {
	CategoryRules []CategoryRule `json:"category_rules"`
}

type Brand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type BrandList struct {
	BrandList []Brand `json:"brand_list"`
}

type ImageInfo struct {
	ImgID     string `json:"img_id"`
	ImgURL    string `json:"img_url"`
	ImgHeight int    `json:"img_height"`
	ImgWidth  int    `json:"img_width"`
	ImgScene  int    `json:"img_scene"`
}

type FileInfo struct {
	FileID   string `json:"file_id"`
	FileURL  string `json:"file_url"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
}

type SalesAttribute struct {
	AttributeID string `json:"attribute_id"`
	CustomValue string `json:"custom_value,omitempty"`
	SKUImage    *Image `json:"sku_img,omitempty"`
	ValueID     string `json:"value_id,omitempty"`
}

type ProductSKU struct {
	ID              string           `json:"id"`
	SellerSku       string           `json:"seller_sku"`
	SalesAttributes []SalesAttribute `json:"sales_attributes"`
}

type ProductData struct {
	CategoryList []struct {
		ID               string `json:"id"`
		IsLeaf           bool   `json:"is_leaf"`
		LocalDisplayName string `json:"local_display_name"`
		ParentID         string `json:"parent_id"`
	} `json:"category_list"`
	CreateTime  int64  `json:"create_time"`
	Description string `json:"description"`
	Images      []struct {
		Height       int      `json:"height"`
		ID           string   `json:"id"`
		ThumbUrlList []string `json:"thumb_url_list"`
		UrlList      []string `json:"url_list"`
		Width        int      `json:"width"`
	} `json:"images"`
	IsCodOpen         bool   `json:"is_cod_open"`
	PackageHeight     int    `json:"package_height"`
	PackageLength     int    `json:"package_length"`
	PackageWeight     string `json:"package_weight"`
	PackageWidth      int    `json:"package_width"`
	ProductAttributes []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Values []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"product_attributes"`
	ProductID     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductStatus int    `json:"product_status"`
	Skus          []struct {
		ID    string `json:"id"`
		Price struct {
			Currency      string `json:"currency"`
			OriginalPrice string `json:"original_price"`
		} `json:"price"`
		SalesAttributes []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			ValueID   string `json:"value_id"`
			ValueName string `json:"value_name"`
		} `json:"sales_attributes"`
		SellerSku  string `json:"seller_sku"`
		StockInfos []struct {
			WarehouseID    string `json:"warehouse_id"`
			AvailableStock int    `json:"available_stock"`
		} `json:"stock_infos"`
	} `json:"skus"`
	UpdateTime     int64 `json:"update_time"`
	WarrantyPeriod struct {
		WarrantyID          int    `json:"warranty_id"`
		WarrantyDescription string `json:"warranty_description"`
	} `json:"warranty_period"`
	WarrantyPolicy string `json:"warranty_policy"`
}

type GetProductListData struct {
	Products []struct {
		CreateTime  int      `json:"create_time"`
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		SaleRegions []string `json:"sale_regions"`
		Skus        []struct {
			ID    string `json:"id"`
			Price struct {
				Currency      string `json:"currency"`
				OriginalPrice string `json:"original_price"`
			} `json:"price"`
			SellerSku  string `json:"seller_sku"`
			StockInfos []struct {
				WarehouseID    string `json:"warehouse_id"`
				AvailableStock int    `json:"available_stock"`
			} `json:"stock_infos"`
		} `json:"skus"`
		Status     int `json:"status"`
		UpdateTime int `json:"update_time"`
	} `json:"products"`
	Total int `json:"total"`
}

type UpdatePriceData struct {
	FailedSKUIDs []string `json:"failed_sku_ids"`
}

type FailedSKUStock struct {
	ID                 string   `json:"id"`
	FailedWarehouseIDs []string `json:"failed_warehouse_ids"`
}

type UpdateStockFailedSKU struct {
	FailedSKUs []FailedSKUStock `json:"failed_skus"`
}

type FailedProductIDs struct {
	FailedProductIDs []string `json:"failed_product_ids"`
}
