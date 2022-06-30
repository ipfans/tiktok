package tiktok

type RejectReverseRequest struct {
	ReverseOrderID         string `json:"reverse_order_id"`
	ReverseRejectReasonKey string `json:"reverse_reject_reason_key"`
	ReverseRejectComments  string `json:"reverse_reject_comments"`
}

type GetReverseListRequest struct {
	UpdateTimeFrom     int `json:"update_time_from,omitempty"`
	UpdateTimeTo       int `json:"update_time_to,omitempty"`
	ReverseType        int `json:"reverse_type,omitempty" example:"REFUND_ONLY = 2 RETURN_AND_REFUND = 3 REQUEST_CANCEL = 4"`
	SortBy             int `json:"sort_by,omitempty" example:"REQUEST_TIME=0(default) UPDATE_TIME=1 REFUND_TOTAL=2"`
	SortType           int `json:"sort_type,omitempty"  example:"ASC=0 DESC=1(default)"`
	Offset             int `json:"offset" validate:"min=0"`
	Size               int `json:"size" validate:"required,min=1,max=100"`
	ReverseOrderStatus int `json:"reverse_order_status,omitempty"`
	OrderID            int `json:"order_id,omitempty"`
	ReverseOrderID     int `json:"reverse_order_id,omitempty"`
}

type GetReverseReasonRequest struct {
	ReverseActionType int `json:"reverse_action_type,omitempty" url:"reverse_action_type,omitempty" example:"Available value: CANCEL = 1;REFUND = 2;RETURN_AND_REFUND = 3;REQUEST_CANCEL_REFUND = 4 "`
	ReasonType        int `json:"reason_type,omitempty" url:"reason_type,omitempty" example:"Available value: STARTE_REVERSE = 1; REJECT_APPLY = 2; REJECT_PARCEL = 3"`
}

type ReverseReason struct {
	ReverseReasonKey         string `json:"reverse_reason_key"`
	ReverseReason            string `json:"reverse_reason"`
	AvailableOrderStatusList []int  `json:"available_order_status_list"`
}

type ReverseReasonList struct {
	ReverseReasonList []ReverseReason `json:"reverse_reason_list"`
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

type ReverseRecord struct {
	Description         string   `json:"description"`
	UpdateTime          int      `json:"update_time"`
	ReasonText          string   `json:"reason_text"`
	AdditionalMessage   string   `json:"additional_message"`
	AdditionalImageList []string `json:"additional_image_list"`
}

type Reverse struct {
	ReverseOrderID     string          `json:"reverse_order_id"`
	OrderID            string          `json:"order_id"`
	RefundTotal        string          `json:"refund_total"`
	Currency           string          `json:"currency"`
	ReverseType        int             `json:"reverse_type"`
	ReturnReason       string          `json:"return_reason"`
	ReturnItemList     []ReturnItem    `json:"return_item_list"`
	ReverseStatusValue int             `json:"reverse_status_value"`
	ReverseRequestTime int             `json:"reverse_request_time"`
	ReverseUpdateTime  int             `json:"reverse_update_time"`
	ReturnTrackingID   string          `json:"return_tracking_id"`
	ReverseRecordList  []ReverseRecord `json:"reverse_record_list"`
}

type ReverseOrdersList struct {
	ReverseList []Reverse `json:"reverse_list"`
	More        bool      `json:"more"`
}
