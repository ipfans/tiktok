package tiktok

type OrderIDReq struct {
	OrderID string `json:"order_id" url:"order_id,omitempty" validate:"required"`
}

type LogisticsGetShippingInfoData struct {
	TrackingInfoList []struct {
		TrackingInfo []struct {
			UpdateTime  int64  `json:"update_time"`
			Description string `json:"description"`
		} `json:"tracking_info"`
	} `json:"tracking_info_list"`
}

type UpdateShippingInfoReq struct {
	OrderID        string `json:"order_id" validate:"required"`
	TrackingNumber string `json:"tracking_number" validate:"required"`
	ProviderID     string `json:"provider_id" validate:"required"`
}

type GetShippingDocumentRequest struct {
	OrderID      string `json:"order_id"   url:"order_id,omitempty" validate:"required"`
	DocumentType string `json:"document_type" url:"document_type,omitempty" validate:"required,oneof=SHIPPING_LABEL PICK_LIST SL_PL" example:"Use this field to specify the type of document to obtain. Available value: SHIPPING_LABEL/ PICK_LIST/ SL_PL SL_PL is to print both SHIPPING_LABEL and PICK_LIST in one pdf. "`
	DocumentSize string `json:"document_size" url:"document_size,omitempty" validate:"oneof=A5 A6" example:"Use this field to specify the size of document to obtain. Available value: A6/A5.  A6 by default. A6 = 0 A5 = 1"`
}

type GetShippingDocumentData struct {
	DocURL string `json:"doc_url"`
}

type GetSubscribedDeliveryRequest struct {
	WarehouseIDList []string `json:"warehouse_id_list"`
}

type GetSubscribedDeliveryData struct {
	WarehouseList []struct {
		WarehouseDelivery []struct {
			DeliveryOption string `json:"delivery_option"`
			ServiceID      string `json:"service_id"`
			ServiceName    string `json:"service_name"`
		} `json:"warehouse_delivery"`
		WarehouseID string `json:"warehouse_id"`
	} `json:"warehouse_list"`
}

type GetWarehouseListData struct {
	WarehouseList []struct {
		WarehouseID           string `json:"warehouse_id"`
		WarehouseName         string `json:"warehouse_name"`
		WarehouseType         int    `json:"warehouse_type"`
		WarehouseSubType      int    `json:"warehouse_sub_type"`
		WarehouseEffectStatus int    `json:"warehouse_effect_status"`
		WarehouseAddress      struct {
			Region        string `json:"region"`
			State         string `json:"state"`
			City          string `json:"city"`
			District      string `json:"district"`
			Town          string `json:"town"`
			FullAddress   string `json:"full_address"`
			Zipcode       string `json:"zipcode"`
			Phone         string `json:"phone"`
			ContactPerson string `json:"contact_person"`
		} `json:"warehouse_address"`
	} `json:"warehouse_list"`
}

type GetShippingProviderData struct {
	DeliveryOptionList []struct {
		DeliveryOptionID   string `json:"delivery_option_id"`
		DeliveryOptionName string `json:"delivery_option_name"`
		ItemWeightLimit    struct {
			MaxWeight int `json:"max_weight"`
			MinWeight int `json:"min_weight"`
		} `json:"item_weight_limit"`
		ItemDimensionLimit struct {
			Length int `json:"length"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"item_dimension_limit"`
		ShippingProviderList []struct {
			ShippingProviderID   string `json:"shipping_provider_id"`
			ShippingProviderName string `json:"shipping_provider_name"`
		} `json:"shipping_provider_list"`
	} `json:"delivery_option_list"`
}
