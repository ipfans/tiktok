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
