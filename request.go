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
	ShoppingProviderID string `json:"shopping_provider_id"`
}

type Pickup struct {
	SelfShipment SelfShipment `json:"self_shipment"`
}

type ShipOrderRequest struct {
	OrderID string  `json:"order_id" validate:"required"`
	Pickup  *Pickup `json:"pick_up,omitempty"`
}

type SearchSettlementsRequest struct {
	RequestTimeFrom int    `json:"request_time_from"`
	RequestTimeTo   int    `json:"request_time_to"`
	PageSize        int    `json:"page_size"`
	Cursor          string `json:"cursor"`
	SortType        int    `json:"sort_type"`
}

type SearchTransactionsRequest struct {
	RequestTimeFrom int   `json:"request_time_from"`
	RequestTimeTo   int   `json:"request_time_to"`
	TransactionType []int `json:"transaction_type" validate:"required"`
	PageSize        int   `json:"page_size"`
	Offset          int   `json:"offset" validate:"gte=0,lte=1000"`
}

type RejectReverseRequest struct {
	ReverseOrderID           string `json:"reverse_order_id" validate:"required"`
	ReverseRejectReasonKey   string `json:"reverse_reject_reason_key" validate:"required"`
	ReverseRejectReasonValue string `json:"reverse_reject_comments"`
}

type GetReverseListRequest struct {
	UpdateTimeFrom int `json:"update_time_from"`
	UpdateTimeTo   int `json:"update_time_to"`
	ReverseType    int `json:"reverse_type"`
	SortBy         int `json:"sort_by"`
	SortType       int `json:"sort_type"`
	Offset         int `json:"offset"`
	Size           int `json:"size"`
}

type GetReverseReasonRequest struct {
	ReverseActionType int `json:"reverse_action_type"`
	ReasonType        int `json:"reason_type"`
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

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Image struct {
	ID string `json:"id"`
}

type ProductCertificationRequest struct {
	ID     string  `json:"id"`
	Files  []File  `json:"files"`
	Images []Image `json:"images"`
}

type CreateProductRequest struct {
	ProductName           string                        `json:"product_name"`
	Description           string                        `json:"description"`
	CategoryID            string                        `json:"category_id"`
	BrandID               string                        `json:"brand_id"`
	Images                []Image                       `json:"images"`
	WarrantyPeriod        int                           `json:"warranty_period"`
	WarrantyPolicy        string                        `json:"warranty_policy"`
	PackageLength         int                           `json:"package_length"`
	PackageWidth          int                           `json:"package_width"`
	PackageHeight         int                           `json:"package_height"`
	PackageWeight         string                        `json:"package_weight"`
	ProductCertifications []ProductCertificationRequest `json:"product_certifications"`
	IsCodOpen             bool                          `json:"is_cod_open"`
	Skus                  []SKU                         `json:"skus"`
}

type StockInfo struct {
	WarehouseID    string `json:"warehouse_id"`
	AvailableStock int    `json:"available_stock"`
}

type EditProductRequest struct {
	ProductID             string                        `json:"product_id"`
	ProductName           string                        `json:"product_name"`
	Description           string                        `json:"description"`
	CategoryID            string                        `json:"category_id"`
	Images                []Image                       `json:"images"`
	PackageWeight         string                        `json:"package_weight"`
	IsCodOpen             bool                          `json:"is_cod_open"`
	Skus                  []SKU                         `json:"skus"`
	BrandID               string                        `json:"brand_id"`
	WarrantyPeriod        int                           `json:"warranty_period"`
	WarrantyPolicy        string                        `json:"warranty_policy"`
	PackageLength         int                           `json:"package_length"`
	PackageWidth          int                           `json:"package_width"`
	PackageHeight         int                           `json:"package_height"`
	ProductCertifications []ProductCertificationRequest `json:"product_certifications"`
}

type ProductSearchRequest struct {
	PageSize      int      `json:"page_size"`
	PageNumber    int      `json:"page_number"`
	SearchStatus  int      `json:"search_status"`
	SellerSkuList []string `json:"seller_sku_list"`
}

type UpdatePriceRequest struct {
	ProductID string `json:"product_id"`
	Skus      []SKU  `json:"skus"`
}

type UpdateStockRequest struct {
	ProductID string `json:"product_id"`
	Skus      []SKU  `json:"skus"`
}
