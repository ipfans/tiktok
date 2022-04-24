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
