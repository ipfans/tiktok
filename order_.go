package tiktok

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

type CancelOrderRequest struct {
	OrderID         string `json:"order_id" validate:"required"`
	CancelReasonKey string `json:"cancel_reason_key" validate:"required"`
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
	RegionCode      string   `json:"region_code"`
}

type PaymentInfo struct {
	Currency                    string  `json:"currency"`
	SubTotal                    float64 `json:"sub_total"`
	ShippingFee                 float64 `json:"shipping_fee"`
	SellerDiscount              float64 `json:"seller_discount"`
	PlatformDiscount            float64 `json:"platform_discount"`
	TotalAmount                 float64 `json:"total_amount"`
	OriginalTotalProductPrice   float64 `json:"original_total_product_price"`
	OriginalShippingFee         float64 `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   float64 `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount float64 `json:"shipping_fee_platform_discount"`
	Taxes                       float64 `json:"taxes"`
}

type Item struct {
	SkuID               string  `json:"sku_id"`
	ProductID           string  `json:"product_id"`
	SkuName             string  `json:"sku_name"`
	Quantity            int     `json:"quantity"`
	SellerSku           string  `json:"seller_sku"`
	ProductName         string  `json:"product_name"`
	SkuImage            string  `json:"sku_image"`
	SkuOriginalPrice    float64 `json:"sku_original_price"`
	SkuSalePrice        float64 `json:"sku_sale_price"`
	SkuPlatformDiscount float64 `json:"sku_platform_discount"`
	SkuSellerDiscount   float64 `json:"sku_seller_discount"`
	SkuExtStatus        int     `json:"sku_ext_status"`
	SkuDisplayStatus    int     `json:"sku_display_status"`
	SkuCancelReason     string  `json:"sku_cancel_reason"`
	SkuCancelUser       string  `json:"sku_cancel_user"`
	SkuRtsTime          int     `json:"sku_rts_time"`
	SkuType             int     `json:"sku_type"`
}

type OrderLineList struct {
	OrderLineID string `json:"order_line_id"`
	SkuID       string `json:"sku_id"`
}

type OrderPackage struct {
	PackageID string `json:"package_id"`
}

type OrderDetail struct {
	OrderID                string           `json:"order_id"`
	OrderStatus            int              `json:"order_status"`
	PaymentMethod          string           `json:"payment_method"`
	DeliveryOption         string           `json:"delivery_option"`
	ShippingProvider       string           `json:"shipping_provider"`
	ShippingProviderID     string           `json:"shipping_provider_id"`
	CreateTime             string           `json:"create_time"`
	PaidTime               int64            `json:"paid_time"`
	BuyerMessage           string           `json:"buyer_message"`
	PaymentInfo            PaymentInfo      `json:"payment_info,omitempty"`
	RecipientAddress       RecipientAddress `json:"recipient_address,omitempty"`
	ItemList               []Item           `json:"item_list,omitempty"`
	CancelReason           string           `json:"cancel_reason"`
	CancelUser             string           `json:"cancel_user"`
	ExtStatus              int              `json:"ext_status"`
	OrderStatusOld         string           `json:"order_status_old"`
	TrackingNumber         string           `json:"tracking_number"`
	RtsTime                int              `json:"rts_time"`
	RtsSLA                 int              `json:"rts_sla"`
	TtsSLA                 int              `json:"tts_sla"`
	CancelOrderSLA         int              `json:"cancel_order_sla"`
	UpdateTime             int              `json:"update_time"`
	PackageList            []OrderPackage   `json:"package_list"`
	ReceiverAddressUpdated int              `json:"receiver_address_updated"`
	BuyerUID               string           `json:"buyer_uid"`
	SplitOrCombineTag      string           `json:"split_or_combine_tag"`
	FulfillmentType        int              `json:"fulfillment_type"`
	SellerNote             string           `json:"seller_note"`
	OrderLineList          []OrderLineList  `json:"order_line_list"`
}

type OrderDetailList struct {
	OrderList []OrderDetail `json:"order_list"`
}

type CancelOrderResponse struct {
	ReverseMainOrderID int64 `json:"reverse_main_order_id"`
}
