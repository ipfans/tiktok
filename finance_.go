package tiktok

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
