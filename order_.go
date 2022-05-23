package tiktok

import (
	"encoding/json"
)

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
	PackageList            []OrderPackage   `json:"package_list"`
}

type OrderDetailList struct {
	OrderList []OrderDetail `json:"order_list"`
}
