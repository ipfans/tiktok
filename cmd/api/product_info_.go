package api

// QueryProductInfoRequest
// lazada:  批量查询 https://bit.ly/3MfPaBQ 单条查询 https://bit.ly/39mmw3i 数据都是一模一样得,
// shopee:  get_item_list 获取 item_id list-> get_item_base_info 和 get_item_extra_info
// tokopedia: https://bit.ly/3PdbI8e ?shop_id=479573&page=1&per_page=2&sort=1
// tiktok:   {"create_time_from":"","create_time_to":"","page_number":"","page_size":"","search_status":"","seller_sku_list":"","update_time_from":"","update_time_to":""}
type QueryProductInfoRequest struct {
	// priority order desc
	ProductIDList []string `json:"product_id_list"`
	SELLERSKUList []string `json:"seller_sku_list"`
	SKUIDList     []string `json:"sku_id_list"`

	// 4个 MP 都有
	//  tiktok  lazada   tokopedia  shopee
	//	MarketplaceType int `json:"marketplace_type"`
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`

	// tiktok  lazada
	CreateTimeFrom int64 `json:"create_time_from"`
	CreateTimeTo   int64 `json:"create_time_to"`

	// tiktok  shopee
	UpdateTimeFrom int64 `json:"update_time_from"`
	UpdateTimeTo   int64 `json:"update_time_to"`
}

/*
type Tokopedia struct {
	Basic struct {
		ProductID       int    `json:"productID"`
		ShopID          int    `json:"shopID"`
		Status          int    `json:"status"`
		Name            string `json:"name"`
		Condition       int    `json:"condition"`
		ChildCategoryID int    `json:"childCategoryID"`
		ShortDesc       string `json:"shortDesc"`
	} `json:"basic"`
	Price struct {
		Value          int `json:"value"`
		Currency       int `json:"currency"`
		LastUpdateUnix int `json:"LastUpdateUnix"`
		Idr            int `json:"idr"`
	} `json:"price"`
	Weight struct {
		Value int `json:"value"`
		Unit  int `json:"unit"`
	} `json:"weight"`
	Stock struct {
		Value        int    `json:"value"`
		StockWording string `json:"stockWording"`
	} `json:"stock"`
	MainStock    int `json:"main_stock"`
	ReserveStock int `json:"reserve_stock"`
	Variant      struct {
		IsParent   bool  `json:"isParent"`
		IsVariant  bool  `json:"isVariant"`
		ChildrenID []int `json:"childrenID"`
	} `json:"variant"`
	Menu struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"menu"`
	Preorder       struct{} `json:"preorder"`
	ExtraAttribute struct {
		MinOrder           int  `json:"minOrder"`
		LastUpdateCategory int  `json:"lastUpdateCategory"`
		IsEligibleCOD      bool `json:"isEligibleCOD"`
	} `json:"extraAttribute"`
	CategoryTree []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		Title         string `json:"title"`
		BreadcrumbURL string `json:"breadcrumbURL"`
	} `json:"categoryTree"`
	Pictures []struct {
		PicID        int    `json:"picID"`
		FileName     string `json:"fileName"`
		FilePath     string `json:"filePath"`
		Status       int    `json:"status"`
		OriginalURL  string `json:"OriginalURL"`
		ThumbnailURL string `json:"ThumbnailURL"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		URL300       string `json:"URL300"`
	} `json:"pictures"`
	GMStats struct {
		TransactionSuccess int `json:"transactionSuccess"`
		TransactionReject  int `json:"transactionReject"`
		CountSold          int `json:"countSold"`
	} `json:"GMStats"`
	Stats struct {
		CountView int `json:"countView"`
	} `json:"stats"`
	Other struct {
		Sku       string `json:"sku"`
		Url       string `json:"url"`
		MobileURL string `json:"mobileURL"`
	} `json:"other"`
	Campaign struct {
		StartDate time.Time `json:"StartDate"`
		EndDate   time.Time `json:"EndDate"`
	} `json:"campaign"`
	Warehouses []struct {
		ProductID   int `json:"productID"`
		WarehouseID int `json:"warehouseID"`
		Price       struct {
			Value          int `json:"value"`
			Currency       int `json:"currency"`
			LastUpdateUnix int `json:"LastUpdateUnix"`
			Idr            int `json:"idr"`
		} `json:"price"`
		Stock struct {
			UseStock bool `json:"useStock"`
			Value    int  `json:"value"`
		} `json:"stock"`
	} `json:"warehouses"`
}

type Lazada struct {
	TotalProducts string `json:"total_products"`
	Products      []struct {
		CreatedTime string `json:"created_time"`
		UpdatedTime string `json:"updated_time"`
		Images      string `json:"images"`
		Skus        []struct {
			Status          string   `json:"Status"`
			Quantity        int      `json:"quantity"`
			ProductWeight   string   `json:"product_weight"`
			Images          []string `json:"Images"`
			SellerSku       string   `json:"SellerSku"`
			ShopSku         string   `json:"ShopSku"`
			Url             string   `json:"Url"`
			PackageWidth    string   `json:"package_width"`
			SpecialToTime   string   `json:"special_to_time"`
			SpecialFromTime string   `json:"special_from_time"`
			PackageHeight   string   `json:"package_height"`
			SpecialPrice    int      `json:"special_price"`
			Price           int      `json:"price"`
			PackageLength   string   `json:"package_length"`
			PackageWeight   string   `json:"package_weight"`
			Available       int      `json:"Available"`
			SkuId           int      `json:"SkuId"`
			SpecialToDate   string   `json:"special_to_date"`
		} `json:"skus"`
		ItemId          string `json:"item_id"`
		PrimaryCategory string `json:"primary_category"`
		MarketImages    string `json:"marketImages"`
		Attributes      struct {
			ShortDescription string `json:"short_description"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			WarrantyType     string `json:"warranty_type"`
			Brand            string `json:"brand"`
		} `json:"attributes"`
		SuspendedSkus []interface{} `json:"suspendedSkus"`
		SubStatus     string        `json:"subStatus"`
		Status        string        `json:"status"`
	} `json:"products"`
}

type Shopee struct {
	ItemList []struct {
		ItemId        int64  `json:"item_id"`
		CategoryId    int    `json:"category_id"`
		ItemName      string `json:"item_name"`
		ItemSku       string `json:"item_sku"`
		CreateTime    int    `json:"create_time"`
		UpdateTime    int    `json:"update_time"`
		AttributeList []struct {
			AttributeId           int    `json:"attribute_id"`
			OriginalAttributeName string `json:"original_attribute_name"`
			IsMandatory           bool   `json:"is_mandatory"`
			AttributeValueList    []struct {
				ValueId           int    `json:"value_id"`
				OriginalValueName string `json:"original_value_name"`
				ValueUnit         string `json:"value_unit"`
			} `json:"attribute_value_list"`
		} `json:"attribute_list"`
		PriceInfo []struct {
			Currency      string `json:"currency"`
			OriginalPrice int    `json:"original_price"`
			CurrentPrice  int    `json:"current_price"`
		} `json:"price_info"`
		StockInfo []struct {
			StockType     int `json:"stock_type"`
			CurrentStock  int `json:"current_stock"`
			NormalStock   int `json:"normal_stock"`
			ReservedStock int `json:"reserved_stock"`
		} `json:"stock_info"`
		Image struct {
			ImageUrlList []string `json:"image_url_list"`
			ImageIdList  []string `json:"image_id_list"`
		} `json:"image"`
		Weight    string `json:"weight"`
		Dimension struct {
			PackageLength int `json:"package_length"`
			PackageWidth  int `json:"package_width"`
			PackageHeight int `json:"package_height"`
		} `json:"dimension"`
		LogisticInfo []struct {
			LogisticId           int     `json:"logistic_id"`
			LogisticName         string  `json:"logistic_name"`
			Enabled              bool    `json:"enabled"`
			ShippingFee          int     `json:"shipping_fee,omitempty"`
			IsFree               bool    `json:"is_free"`
			EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty"`
		} `json:"logistic_info"`
		PreOrder struct {
			IsPreOrder bool `json:"is_pre_order"`
			DaysToShip int  `json:"days_to_ship"`
		} `json:"pre_order"`
		Condition   string `json:"condition"`
		SizeChart   string `json:"size_chart"`
		ItemStatus  string `json:"item_status"`
		HasModel    bool   `json:"has_model"`
		PromotionId int    `json:"promotion_id"`
		Brand       struct {
			BrandId           int    `json:"brand_id"`
			OriginalBrandName string `json:"original_brand_name"`
		} `json:"brand"`
		TaxInfo struct {
			Ncm           int `json:"ncm"`
			SameStateCfop int `json:"same_state_cfop"`
			DiffStateCfop int `json:"diff_state_cfop"`
			Csosn         int `json:"csosn"`
			Origin        int `json:"origin"`
		} `json:"tax_info"`
		DescriptionType string `json:"description_type"`
		DescriptionInfo struct {
			ExtendedDescription struct {
				FieldList []struct {
					FieldType string `json:"field_type"`
					Text      string `json:"text,omitempty"`
					ImageInfo struct {
						ImageId  string `json:"image_id"`
						ImageUrl string `json:"image_url"`
					} `json:"image_info,omitempty"`
				} `json:"field_list"`
			} `json:"extended_description"`
		} `json:"description_info"`
	} `json:"item_list"`
}

type Tiktok struct {
	CategoryList []struct {
		Id               string `json:"id"`
		IsLeaf           bool   `json:"is_leaf"`
		LocalDisplayName string `json:"local_display_name"`
		ParentId         string `json:"parent_id"`
	} `json:"category_list"`
	CreateTime  int64  `json:"create_time"`
	Description string `json:"description"`
	Images      []struct {
		Height       int      `json:"height"`
		Id           string   `json:"id"`
		ThumbUrlList []string `json:"thumb_url_list"`
		UrlList      []string `json:"url_list"`
		Width        int      `json:"width"`
	} `json:"images"`
	IsCodOpen         bool   `json:"is_cod_open"`
	PackageHeight     int    `json:"package_height"`
	PackageLength     int    `json:"package_length"`
	PackageWeight     string `json:"package_weight"`
	PackageWidth      int    `json:"package_width"`
	ProductAttributes []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Values []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"values"`
	} `json:"product_attributes"`
	ProductId     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductStatus int    `json:"product_status"`
	Skus          []struct {
		Id    string `json:"id"`
		Price struct {
			Currency      string `json:"currency"`
			OriginalPrice string `json:"original_price"`
		} `json:"price"`
		SalesAttributes []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ValueId   string `json:"value_id"`
			ValueName string `json:"value_name"`
		} `json:"sales_attributes"`
		SellerSku  string `json:"seller_sku"`
		StockInfos []struct {
			AvailableStock int    `json:"available_stock"`
			WarehouseId    string `json:"warehouse_id"`
		} `json:"stock_infos"`
	} `json:"skus"`
	UpdateTime     int64 `json:"update_time"`
	WarrantyPeriod struct {
		WarrantyDescription string `json:"warranty_description"`
		WarrantyId          int    `json:"warranty_id"`
	} `json:"warranty_period"`
	WarrantyPolicy string `json:"warranty_policy"`
}

type ProductInfoData struct{}
*/
