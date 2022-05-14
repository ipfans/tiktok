package tiktok

type GetAccessTokenRequest struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	AuthCode  string `json:"auth_code"`
	GrantType string `json:"grant_type"`
}

type RefreshTokenRequest struct {
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type GetOrderListRequest struct {
	CreateTimeFrom int    `json:"create_time_from,omitempty"`
	CreateTimeTo   int    `json:"create_time_to,omitempty"`
	UpdateTimeFrom int    `json:"update_time_from,omitempty"`
	UpdateTimeTo   int    `json:"update_time_to,omitempty"`
	OrderStatus    int    `json:"order_status,omitempty"`
	Cursor         string `json:"cursor,omitempty"`
	SortType       int    `json:"sort_type,omitempty"`
	SortBy         string `json:"sort_by,omitempty"`
	PageSize       int    `json:"page_size,omitempty"`
}

type GetOrderDetailRequest struct {
	OrderIDList []string `json:"order_id_list"`
}

type SelfShipment struct {
	TrackingNumber     string `json:"tracking_number"`
	ShippingProviderID string `json:"shipping_provider_id"`
}

type Pickup struct {
	SelfShipment SelfShipment `json:"self_shipment"`
}

type ShipOrderRequest struct {
	OrderID string  `json:"order_id" validate:"required"`
	Pickup  *Pickup `json:"pick_up,omitempty"`
}

type SearchSettlementsRequest struct {
	RequestTimeFrom int    `json:"request_time_from,omitempty"`
	RequestTimeTo   int    `json:"request_time_to,omitempty"`
	PageSize        int    `json:"page_size" validate:"required,min=1,max=100"`
	Cursor          string `json:"cursor,omitempty"`
	SortType        int    `json:"sort_type"  validate:"required" example:"Available values: 1 (DESC), 2 (ASC) Default value 1"`
}

type SearchTransactionsRequest struct {
	RequestTimeFrom int   `json:"request_time_from,omitempty"`
	RequestTimeTo   int   `json:"request_time_to,omitempty"`
	TransactionType []int `json:"transaction_type" validate:"required,min=1" example:"Withdraw:1 Settle:2 Transfer:3 Reverse:4"`
	PageSize        int   `json:"page_size" validate:"required,min=1,max=100" `
	Offset          int   `json:"offset" validate:"gte=0,lte=10000"`
}

type RejectReverseRequest struct {
	ReverseOrderID         string `json:"reverse_order_id"`
	ReverseRejectReasonKey string `json:"reverse_reject_reason_key"`
	ReverseRejectComments  string `json:"reverse_reject_comments"`
}

type GetReverseListRequest struct {
	UpdateTimeFrom int `json:"update_time_from"`
	UpdateTimeTo   int `json:"update_time_to"`
	ReverseType    int `json:"reverse_type" example:"REFUND_ONLY = 2 RETURN_AND_REFUND = 3 REQUEST_CANCEL = 4"`
	SortBy         int `json:"sort_by" example:"REQUEST_TIME=0(default) UPDATE_TIME=1 REFUND_TOTAL=2"`
	SortType       int `json:"sort_type"  example:"ASC=0 DESC=1(default)"`
	Offset         int `json:"offset" validate:"required,min=0"`
	Size           int `json:"size" validate:"required,min=1,max=100"`
}

type GetReverseReasonRequest struct {
	ReverseActionType int `json:"reverse_action_type,omitempty" url:"reverse_action_type,omitempty" example:"Available value: CANCEL = 1;REFUND = 2;RETURN_AND_REFUND = 3;REQUEST_CANCEL_REFUND = 4 "`
	ReasonType        int `json:"reason_type,omitempty" url:"reason_type,omitempty" example:"Available value: STARTE_REVERSE = 1; REJECT_APPLY = 2; REJECT_PARCEL = 3"`
}

type ImgScene int

const (
	ImgSceneProductImage ImgScene = iota + 1
	ImgSceneDescriptionImage
	ImgSceneAttributeImage
	ImgSceneCertificationImage
	ImgSceneSizeChartImage
)

type SKU struct {
	ID              string           `json:"id,omitempty"`
	SalesAttributes []SalesAttribute `json:"sales_attributes,omitempty"`
	StockInfos      []StockInfo      `json:"stock_infos,omitempty"`
	SellerSku       string           `json:"seller_sku,omitempty"`
	OriginalPrice   string           `json:"original_price,omitempty"`
}

type ImgID struct {
	ID string `json:"id"`
}

type ProductCertificationRequest struct {
	ID     string  `json:"id"`
	Files  []File  `json:"files"`
	Images []ImgID `json:"images"`
}

// CreateProductRequest is creating product
// https://developers.tiktok-shops.com/documents/document/234547
type CreateProductRequest struct {
	ProductName    string  `json:"product_name" validate:"required"`
	Description    string  `json:"description"  validate:"required"`
	CategoryID     string  `json:"category_id"  validate:"required"`
	BrandID        string  `json:"brand_id"`
	Images         []ImgID `json:"images,omitempty"`
	WarrantyPeriod int     `json:"warranty_period"`
	WarrantyPolicy string  `json:"warranty_policy"`
	PackageLength  int     `json:"package_length"`
	PackageWidth   int     `json:"package_width"`
	PackageHeight  int     `json:"package_height"`
	PackageWeight  string  `json:"package_weight"  validate:"required"`
	SizeChart      struct {
		ImgID string `json:"img_id" validate:"required"`
	} `json:"size_chart"`
	ProductCertifications []ProductCertificationRequest `json:"product_certifications"`
	IsCodOpen             bool                          `json:"is_cod_open"`
	Skus                  []SKU                         `json:"skus"  validate:"required"`
	DeliveryServiceIDS    []string                      `json:"delivery_service_ids,omitempty"`
	ProductAttributes     []struct {
		AttributeID     string `json:"attribute_id"`
		AttributeName   string `json:"attribute_name"`
		AttributeValues []struct {
			ValueID   string `json:"value_id"`
			ValueName string `json:"value_name"`
		} `json:"attribute_values"`
	} `json:"product_attributes,omitempty"`
}

type StockInfo struct {
	WarehouseID    string `json:"warehouse_id"`
	AvailableStock int    `json:"available_stock"`
}

type EditProductCertifications []struct {
	Files []struct {
		ID   string `json:"id" validate:"required"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"files"`
	ID     string `json:"id" validate:"required"`
	Images []struct {
		ID string `json:"id" validate:"required"`
	} `json:"images"`
}

type EditAttributes []struct {
	AttributeID     string `json:"attribute_id"`
	AttributeName   string `json:"attribute_name"`
	AttributeValues []struct {
		ValueID   string `json:"value_id"`
		ValueName string `json:"value_name"`
	} `json:"attribute_values"`
}

// EditProductRequest is edit product attributes
// https://developers.tiktok-shops.com/documents/document/234555
type EditProductRequest struct {
	ProductID      string  `json:"product_id" validate:"required"`
	ProductName    string  `json:"product_name" validate:"required"`
	Description    string  `json:"description" validate:"required"`
	CategoryID     string  `json:"category_id" validate:"required"`
	BrandID        string  `json:"brand_id"`
	Images         []ImgID `json:"images,omitempty"`
	WarrantyPeriod int     `json:"warranty_period" example:"Need to choose among the candidate values provided by the platform: 1: 4 weeks" 2:"2 months" 3:"3 months" 4:"4 months" 5:"5 months" 6:"6 months" 7:"7 months" 8:"8 months" 9:"9 months" 10:"10 months" 11:"11 months" 12:"12 months" 13:"2 years" 14:"3 years" 15:"1 week" 16:"2 weeks" 17:"18 months" 18:"4 years" 19:"5 years" 20:"10 years" 21:"lifetime warranty"`
	WarrantyPolicy string  `json:"warranty_policy"`
	PackageLength  int     `json:"package_length" example:"14. unit is cm"`
	PackageHeight  int     `json:"package_height" example:"14. unit is cm"`
	PackageWidth   int     `json:"package_width" example:"14. unit is cm"`
	PackageWeight  string  `json:"package_weight" validate:"required" example:"14. unit is cm"`
	SizeChart      struct {
		ImgID string `json:"img_id" validate:"required"`
	} `json:"size_chart,omitempty"`
	Certifications     EditProductCertifications `json:"product_certifications"`
	IsCodOpen          bool                      `json:"is_cod_open,omitempty"`
	Skus               []SKU                     `json:"skus"`
	DeliveryServiceIDS string                    `json:"delivery_service_ids"`
	Attribute          EditAttributes            `json:"product_attributes,omitempty"`
}

type ProductSearchRequest struct {
	PageSize       int      `json:"page_size" validate:"min=1,max=100"`
	PageNumber     int      `json:"page_number" validate:"min=1"`
	SearchStatus   int      `json:"search_status,omitempty"`
	SellerSkuList  []string `json:"seller_sku_list,omitempty"`
	UpdateTimeFrom int64    `json:"update_time_from,omitempty"`
	UpdateTimeTo   int64    `json:"update_time_to,omitempty"`
	CreateTimeFrom int64    `json:"create_time_from,omitempty" `
	CreateTimeTo   int64    `json:"create_time_to,omitempty"`
}
type UpdatePriceRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Skus      []struct {
		ID            string `json:"id" validate:"required"`
		OriginalPrice string `json:"original_price" validate:"required"`
	} `json:"skus" validate:"min=1"`
}

type UpdateStockRequest struct {
	ProductID string `json:"product_id"`
	Skus      []struct {
		ID         string `json:"id"`
		StockInfos []struct {
			AvailableStock int    `json:"available_stock" validate:"min=0,max=99999"`
			WarehouseID    string `json:"warehouse_id"`
		} `json:"stock_infos"`
	} `json:"skus"`
}
