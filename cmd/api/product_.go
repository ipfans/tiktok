package api

type GetBrandReq struct {
	PageResult
}

type ImgID struct {
	ID string `json:"id"`
}

type StandardSKU struct {
	Weight         int     `json:"weight"`
	Price          int     `json:"price"`
	WarehouseID    int     `json:"warehouse_id"`
	AvailableStock int     `json:"available_stock"`
	SellerSku      int     `json:"seller_sku"`
	Images         []ImgID `json:"images"`
	BrandID        string  `json:"brand_id"`
	BrandName      string  `json:"brand_name"`
}

type StandardShipperOption struct {
	SizeID     int  `json:"size_id"`
	ShipFee    int  `json:"ship_fee"`
	Enable     bool `json:"enabled"`
	LogisticID int  `json:"logistic_id"`
	IsFreeShip bool `json:"is_free_ship"`
}

type AddEditProductReq struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	BrandID         string  `json:"brand_id"`
	BrandName       string  `json:"brand_name"`
	Images          []ImgID `json:"images"`
	CategoryID      string  `json:"category_id"`
	IsSingleProduct bool    `json:"is_single_product"`

	Status                 string                  `json:"status"`
	MinOrder               int                     `json:"min_order"`
	ProductWeight          int                     `json:"product_weight"`
	WeightUnit             string                  `json:"weight_unit"`
	Condition              string                  `json:"condition"`
	StandardSKUS           []StandardSKU           `json:"standard_sku_list"`
	StandardShipperOptions []StandardShipperOption `json:"standard_shipper_option_list"`

	Stock     int    `json:"stock"`
	StockUnit string `json:"stock_unit"`

	/*
		ProductAttrs  ProductAttrs `json:"product_attrs"`
		OriginalPrice int          `json:"original_price"`
		Stock         int          `json:"stock"`
		Options       map[string]interface{} `json:"options"`
	*/
}

type SKU struct {
	ID int `json:"id"`
	// SKUNo       string `json:"sku_no"`
	// SellerSKUNo string `json:"seller_sku_no"`
	Name string `json:"name"`
}

/*
type Logistic struct {
	ID         int    `json:"id"`
	LogisticNo string `json:"logistic_no"`
	Name       string `json:"name"`
}

type ShipperS3Img struct {
	ID string `json:"id"`
}

type ProductAttrs struct {
	ProductName string `json:"product_name" `
	Description string `json:"description"  `
	CategoryID  string `json:"category_id"  `
	BrandID     string `json:"brand_id" `
	BrandName   string `json:"brand_name" `

	Skus      []SKU          `json:"skus" `
	Logistics []Logistic     `json:"logistics"`
	Images    []ShipperS3Img `json:"images"`
	//	ThumbNailImg []ShipperS3Img `json:"thumb_nail_img"`

	PackageLengthType int    `json:"package_length_type" ` // example:"1m 2cm 3mm 3μm 4nm"
	PackageLength     int    `json:"package_length"`
	PackageWidth      int    `json:"package_width"`
	PackageHeight     int    `json:"package_height"`
	PackageWeight     int    `json:"package_weight"`
	PackageWeightType int    `json:"package_weight_type"` // example:"1g 2kg"
	PackageContent    string `json:"package_content"`     // example:"describe package content"

	IsCodOpen bool `json:"is_cod_open"`

	TokopediaMinOrder  int    `json:"min_order"` // from tokopedia, Minimum order required to purchase the product. Can only be a positive integer
	TokopediaCondition string `json:"condition"` // from tokopedia,  default new
	TokopediaStatus    string `json:"status"`    // from tokopedia,

	WarrantyType   int    `json:"warranty_type"`
	WarrantyPeriod int    `json:"warranty_period" `
	WarrantyPolicy string `json:"warranty_policy"`

	IsShipFree bool `json:"is_ship_free"`
	ShipFee    int  `json:"ship_fee"`

	IsDangerous  bool `json:"is_dangerous"`
	IsWaterProof bool `json:"is_water_proof"`
	IsFireProof  bool `json:"is_fire_proof"`
}*/

type SKUPrice struct {
	SKUID string `json:"sku_id"`
	Price int    `json:"price" ` // example:"default Rp"
}

type UpdateProductPriceRequest struct {
	SKUPrices []SKUPrice `json:"prices"`
}

type SKUStock struct {
	SKUID string `json:"sku_id"`
	Stock int    `json:"stock" ` // example:"default Rp"
}

type UpdateProductStockRequest struct {
	Stocks []SKUStock `json:"stocks"`
}

type ProductIDSRequest struct {
	ProductIDS []string `json:"product_ids"`
}

type ActivateProductRequest struct {
	SpuID string `json:"spu_id,omitempty"`
	Skus  []SKU  `json:"skus" validate:"required"`
}

type UpDownShelfReq struct {
	SpuID string `json:"spu_id,omitempty"`

	Skus []struct {
		SKU
		OnShelf bool `json:"on_shelf"`
	} `json:"skus" validate:"required"`
}
