package tiktok

type SearchPreCombinePkgRequest struct {
	Cursor   string `json:"cursor"`
	PageSize int    `json:"page_size" validate:"required,min=1,max=50"`
}

type SearchPreCombinePkgData struct {
	PreCombinePkgList []struct {
		PreCombinePkgID string   `json:"pre_combine_pkg_id"`
		OrderIDList     []string `json:"order_id_list"`
	} `json:"pre_combine_pkg_list"`
	NextCursor string `json:"next_cursor"`
	More       bool   `json:"more"`
	Total      int    `json:"total"`
}

type ConfirmPreCombinePkgRequest struct {
	PreCombinePkgList []struct {
		PreCombinePkgID string   `json:"pre_combine_pkg_id" validate:"required"`
		OrderIDList     []string `json:"order_id_list"  validate:"omitempty"`
	} `json:"pre_combine_pkg_list" validate:"required,min=1"`
}

type PreCombinePkgData struct {
	PackageList []struct {
		PackageID   string   `json:"package_id"`
		OrderIDList []string `json:"order_id_list"`
	} `json:"package_list"`

	FailedPackageList []struct {
		PackageID string `json:"package_id"`
		Code      int    `json:"code"`
		Message   string `json:"message"`
	} `json:"failed_package_list"`
}

type RemovePackageOrderRequest struct {
	PackageID   string   `json:"package_id" validate:"required"`
	OrderIDList []string `json:"order_id_list" validate:"required,min=1"`
}

type PackageIDRequest struct {
	PackageID string `json:"package_id" validate:"required"`
}

type GetPackagePickupConfigData struct {
	IsPickUp       bool `json:"is_pick_up"`
	IsDropOff      bool `json:"is_drop_off"`
	PickUpTimeList []struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Available bool   `json:"available"`
	} `json:"pick_up_time_list"`
	DropOffPointURL string `json:"drop_off_point_url"`
}

type FulfillmentPickUp struct {
	PickUpStartTime int64 `json:"pick_up_start_time"`
	PickUpEndTime   int64 `json:"pick_up_end_time"`
}

/*
type T struct {
	PackageID    string            `json:"package_id" validate:"required"`
	PickUpType   int               `json:"pick_up_type"`
	PickUp       FulfillmentPickUp `json:"pick_up"`
	SelfShipment SelfShipment      `json:"self_shipment"`
}*/

type ShipPackageRequest struct {
	PackageId  string `json:"package_id"`
	PickUpType int    `json:"pick_up_type"`
	PickUp     struct {
		PickUpStartTime int `json:"pick_up_start_time"`
		PickUpEndTime   int `json:"pick_up_end_time"`
	} `json:"pick_up"`
	SelfShipment struct {
		TrackingNumber     string `json:"tracking_number"`
		ShippingProviderId string `json:"shipping_provider_id"`
	} `json:"self_shipment"`
}

type ShipPackageData struct {
	FailPackages []struct {
		FailCode   int    `json:"fail_code"`
		FailReason string `json:"fail_reason"`
		PackageID  string `json:"package_id"`
	} `json:"fail_packages"`
}

type SearchPackageRequest struct {
	CreateTimeFrom int64  `json:"create_time_from,omitempty" `
	CreateTimeTo   int64  `json:"create_time_to,omitempty"`
	UpdateTimeFrom int64  `json:"update_time_from,omitempty"`
	UpdateTimeTo   int64  `json:"update_time_to,omitempty"`
	PackageStatus  int    `json:"package_status,omitempty"`
	Cursor         string `json:"cursor,omitempty"`
	SortBy         int    `json:"sort_by,omitempty" validate:"oneof=1 2 3"  example:"Default value: 1 CREATE_TIME = 1 - ORDER_PAY_TIME = 2 - UPDATE_TME = 3"`
	SortType       int    `json:"sort_type,omitempty" validate:"oneof=1 2" example:"ASC = 1 DESC = 2"`
	PageSize       int    `json:"page_size,omitempty" validate:"min=1,max=50"`
}

type SearchPackageData struct {
	More        bool   `json:"more"`
	NextCursor  string `json:"next_cursor"`
	PackageList []struct {
		PackageId     string `json:"package_id"`
		PackageStatus int    `json:"package_status"`
		CreateTime    int64  `json:"create_time"`
		UpdateTime    int64  `json:"update_time"`
	} `json:"package_list"`
	Total int `json:"total"`
}

type GetPackageDetailData struct {
	PackageID     string `json:"package_id"`
	OrderInfoList struct {
		OrderID string `json:"order_id"`
		SkuList struct {
			SkuID    string `json:"sku_id"`
			SkuName  string `json:"sku_name"`
			SkuImage string `json:"sku_image"`
			Quantity string `json:"quantity"`
		} `json:"sku_list"`
	} `json:"order_info_list"`
	PackageStatus       int      `json:"package_status" example:"TO_FULFILL = 1 , PROCESSING = 2 , FULFILLING = 3 , COMPLETED = 4 , CANCELLED = 5"`
	PackageFreezeStatus int      `json:"package_freeze_status" example:"FREEZE = 1, When the package is pending review for cancellation. UNFREEZE = 2"`
	ScTag               int      `json:"sc_tag" example:"DEFAULT = 0 , COMBINE =1 , SPLIT = 2"`
	PrintTag            int      `json:"print_tag"`
	SkuTag              int      `json:"sku_tag" example:"Whether there are multiple SKU IDs in a package, ONE = 1 MANY = 2"`
	NoteTag             int      `json:"note_tag" example:"BUYER_UNNOTED = 0 BUYER_NOTED = 1"`
	DeliveryOption      string   `json:"delivery_option" example:"STANDARD=1 EXPRESS=2 ECONOMY=3 SEND_BY_SELLER=4"`
	ShippingProvider    string   `json:"shipping_provider"`
	ShippingProviderID  string   `json:"shipping_provider_id"`
	TrackingNumber      string   `json:"tracking_number"`
	PickUpType          int      `json:"pick_up_type"`
	PickUpStartTime     int64    `json:"pick_up_start_time"`
	PickUpEndTime       int64    `json:"pick_up_end_time"`
	CreateTime          int64    `json:"create_time"`
	UpdateTime          int64    `json:"update_time"`
	OrderLineIDList     []string `json:"order_line_id_list"`
}

type GetPackageShippingInfoData struct {
	TrackingInfoList []struct {
		UpdateTime  int    `json:"update_time"`
		Description string `json:"description"`
	} `json:"tracking_info_list"`
}

type UpdatePackageShippingInfoRequest struct {
	PackageID      string `json:"package_id" validate:"required"`
	TrackingNumber string `json:"tracking_number" validate:"required"`
	ProviderID     string `json:"provider_id" validate:"required"`
}

type UpdatePackageShippingInfoData struct {
	FailedReason  string `json:"failed_reason"`
	UpdateSuccess bool   `json:"update_success"`
}

type GetPackageShippingDocumentRequest struct {
	PackageID    string `json:"package_id" validate:"required"`
	DocumentType int    `json:"document_type" validate:"required" example:"Available value: SHIPPING_LABEL = 1 PICK_LIST = 2 SL+PL = 3 PACK_LIST is not available in this version."`
	DocumentSize int    `json:"document_size" validate:"isdefault=0,oneof=0 1" example:"Use this field to specify the size of document to obtain. Available value: A6/A5.  A6 by default. A6 = 0 A5 = 1"`
}

type GetPackageShippingDocumentData struct {
	DocURL string `json:"doc_url"`
}

type VerifyOrderSplitRequest struct {
	OrderIDList []int64 `json:"order_id_list" validate:"required,min=1,max=50"`
}

type VerifyOrderSplitData struct {
	VerifyOrderResult []struct {
		OrderID           int64 `json:"order_id"`
		VerifyOrderResult bool  `json:"verify_order_result"`
	} `json:"verify_order_result"`
	FailList []struct {
		OrderID    int64  `json:"order_id"`
		FailReason string `json:"fail_reason"`
	} `json:"fail_list"`
}

type ConfirmOrderSplitRequest struct {
	OrderID    int64 `json:"order_id" validate:"required"`
	SplitGroup []struct {
		PreSplitPkgID   int     `json:"pre_split_pkg_id" validate:"required"`
		OrderLineIDList []int64 `json:"order_line_id_list" validate:"min=1"`
	} `json:"split_group" validate:"min=1"`
}

type ConfirmOrderSplitData struct {
	ConfirmResultList []struct {
		PreSplitPkgID int  `json:"pre_split_pkg_id"`
		ConfirmResult bool `json:"confirm_result"`
	} `json:"confirm_result_list"`
	SuccessList []struct {
		PreSplitPkgID int `json:"pre_split_pkg_id"`
		PackageID     int `json:"package_id"`
	} `json:"success_list"`

	FailList []struct {
		PreSplitPkgID int64  `json:"pre_split_pkg_id"`
		FailReason    string `json:"fail_reason"`
	}
}
