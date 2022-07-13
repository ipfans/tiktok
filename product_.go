package tiktok

type ImgScene int

const (
	ImgSceneProductImage ImgScene = iota + 1
	ImgSceneDescriptionImage
	ImgSceneAttributeImage
	ImgSceneCertificationImage
	ImgSceneSizeChartImage
)

type Category struct {
	ID               string `json:"id"`
	ParentID         string `json:"parent_id"`
	LocalDisplayName string `json:"local_display_name"`
	IsLeaf           bool   `json:"is_leaf"`
}

type CategoryList struct {
	CategoryList []Category `json:"category_list"`
}

type InputType struct {
	IsMandatory        bool `json:"is_mandatory"`
	IsMultipleSelected bool `json:"is_multiple_selected"`
	IsCustomized       bool `json:"is_customized"`
}

type Attribute struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	AttributeType int       `json:"attribute_type"`
	InputType     InputType `json:"input_type"`
	Values        []Value   `json:"values"`
}

type AttributeList struct {
	Attributes []Attribute `json:"attributes"`
}

type ProductCertification struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Sample      string `json:"sample"`
	IsMandatory bool   `json:"is_mandatory"`
}

type ExemptionOfIdentifierCode struct {
	SupportIdentifierCodeExemption bool `json:"support_identifier_code_exemption"`
}

type CategoryRule struct {
	ProductCertifications     []ProductCertification    `json:"product_certifications"`
	SupportSizeChart          bool                      `json:"support_size_chart"`
	SupportCod                bool                      `json:"support_cod"`
	ExemptionOfIdentifierCode ExemptionOfIdentifierCode `json:"exemption_of_identifier_code"`
}

type CategoryRules struct {
	CategoryRules []CategoryRule `json:"category_rules"`
}

type Brand struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"` // ONLINE=1;  OFFLINE=2;
}

type BrandList struct {
	BrandList []Brand `json:"brand_list"`
}

type ImageInfo struct {
	ImgID     string `json:"img_id"`
	ImgURL    string `json:"img_url"`
	ImgHeight int    `json:"img_height"`
	ImgWidth  int    `json:"img_width"`
	ImgScene  int    `json:"img_scene"`
}

type FileInfo struct {
	FileID   string `json:"file_id"`
	FileURL  string `json:"file_url"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
}

type File struct {
	ID   string   `json:"id"`
	List []string `json:"list,omitempty"`
	Name string   `json:"name"`
	Type string   `json:"type"`
}

type ImgID struct {
	ID string `json:"id"`
}

type ProductCertificationRequest struct {
	ID     string  `json:"id"`
	Files  []File  `json:"files"`
	Images []ImgID `json:"images"`
}

type SalesAttribute struct {
	AttributeID string `json:"attribute_id"`
	CustomValue string `json:"custom_value,omitempty"`
	SKUImage    *ImgID `json:"sku_img,omitempty"`
	ValueID     string `json:"value_id,omitempty"`
}

type StockInfo struct {
	WarehouseID    string `json:"warehouse_id,omitempty"`
	AvailableStock int    `json:"available_stock"`
}

type SKU struct {
	ID              string           `json:"id,omitempty"`
	SalesAttributes []SalesAttribute `json:"sales_attributes,omitempty"`
	StockInfos      []StockInfo      `json:"stock_infos,omitempty"`
	SellerSku       string           `json:"seller_sku,omitempty"`
	OriginalPrice   string           `json:"original_price,omitempty"`
}

type SizePic struct {
	ImgID string `json:"img_id" validate:"required"`
}

type ProductAttr struct {
	AttributeID     string           `json:"attribute_id"`
	AttributeName   string           `json:"attribute_name"`
	AttributeValues []AttributeValue `json:"attribute_values"`
}

// CreateProductRequest is creating product
// https://developers.tiktok-shops.com/documents/document/234547
type CreateProductRequest struct {
	ProductName           string                        `json:"product_name" validate:"required"`
	Description           string                        `json:"description"  validate:"required"`
	CategoryID            string                        `json:"category_id"  validate:"required"`
	BrandID               string                        `json:"brand_id"`
	Images                []ImgID                       `json:"images,omitempty"`
	WarrantyPeriod        int                           `json:"warranty_period"`
	WarrantyPolicy        string                        `json:"warranty_policy"`
	PackageLength         int                           `json:"package_length"`
	PackageWidth          int                           `json:"package_width"`
	PackageHeight         int                           `json:"package_height"`
	PackageWeight         string                        `json:"package_weight"  validate:"required"`
	SizeChart             SizePic                       `json:"size_chart"`
	ProductCertifications []ProductCertificationRequest `json:"product_certifications"`
	IsCodOpen             bool                          `json:"is_cod_open"`
	Skus                  []SKU                         `json:"skus"  validate:"required"`
	DeliveryServiceIDS    []string                      `json:"delivery_service_ids,omitempty"`
	ProductAttributes     []ProductAttr                 `json:"product_attributes,omitempty"`
}

type Image struct {
	ID           string   `json:"id"`
	Height       int      `json:"height"`
	Width        int      `json:"width"`
	ThumbUrlList []string `json:"thumb_url_list"`
	UrlList      []string `json:"url_list"`
}

type Value struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductAttribute struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type Certification struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Files  []File  `json:"files"`
	Images []Image `json:"images"`
}

type QCReason struct {
	Reason     string   `json:"reason"`
	SubReasons []string `json:"sub_reasons"`
}

type DeliveryService struct {
	DeliveryServiceID     int64  `json:"delivery_service_id"`
	DeliveryOptionName    string `json:"delivery_option_name"`
	DeliveryServiceStatus bool   `json:"delivery_service_status"`
}

type VideoInfo struct {
	BackupUrl string `json:"backup_url"`
	Bitrate   int    `json:"bitrate"`
	FileHash  string `json:"file_hash"`
	Format    string `json:"format"`
	Height    int    `json:"height"`
	MainUrl   string `json:"main_url"`
	Size      int    `json:"size"`
	UrlExpire int64  `json:"url_expire"`
	Width     int    `json:"width"`
}

type Video struct {
	Duration   float64     `json:"duration"`
	ID         string      `json:"id"`
	MediaType  string      `json:"media_type"`
	PostUrl    string      `json:"post_url"`
	VideoInfos []VideoInfo `json:"video_infos"`
}

type Price struct {
	Currency        string `json:"currency"`
	OriginalPrice   string `json:"original_price"`
	PriceIncludeVat string `json:"price_include_vat"`
}

type SalesAttr struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ValueID   string `json:"value_id"`
	ValueName string `json:"value_name"`
	SkuImg    Image  `json:"sku_img"`
}

type SKUData struct {
	ID              string      `json:"id"`
	SellerSku       string      `json:"seller_sku"`
	Price           Price       `json:"price"`
	StockInfos      []StockInfo `json:"stock_infos"`
	SalesAttributes []SalesAttr `json:"sales_attributes"`
}

type WarrantyPeriod struct {
	WarrantyID          int    `json:"warranty_id"`
	WarrantyDescription string `json:"warranty_description"`
}

type ProductData struct {
	ProductID             string             `json:"product_id"`
	ProductStatus         int                `json:"product_status"` // 1-draft、2-pending、3-failed(initial creation)、4-live、5-seller_deactivated、6-platform_deactivated、7-freeze 8-deleted
	ProductName           string             `json:"product_name"`
	CategoryList          []Category         `json:"category_list"`
	Brand                 Brand              `json:"brand"`
	Images                []Image            `json:"images"`
	Video                 Video              `json:"video"`
	Description           string             `json:"description"`
	WarrantyPeriod        WarrantyPeriod     `json:"warranty_period"`
	WarrantyPolicy        string             `json:"warranty_policy"`
	PackageHeight         int                `json:"package_height"`
	PackageLength         int                `json:"package_length"`
	PackageWidth          int                `json:"package_width"`
	PackageWeight         string             `json:"package_weight"`
	Skus                  []SKUData          `json:"skus"`
	ProductCertifications []Certification    `json:"product_certifications"`
	SizeChart             Image              `json:"size_chart"`
	IsCodOpen             bool               `json:"is_cod_open"`
	ProductAttributes     []ProductAttribute `json:"product_attributes"`
	QcReasons             []QCReason         `json:"qc_reasons"`
	UpdateTime            int64              `json:"update_time"`
	CreateTime            int64              `json:"create_time"`
	DeliveryServices      []DeliveryService  `json:"delivery_services"`
}

type AttributeValue struct {
	ValueID   string `json:"value_id"`
	ValueName string `json:"value_name"`
}

type EditAttributes []struct {
	AttributeID     string           `json:"attribute_id"`
	AttributeName   string           `json:"attribute_name"`
	AttributeValues []AttributeValue `json:"attribute_values"`
}

type EditProductCertifications []struct {
	Files  []File  `json:"files"` // no type field
	ID     string  `json:"id" validate:"required"`
	Images []ImgID `json:"images"`
}

// EditProductRequest is edit product attributes
// https://developers.tiktok-shops.com/documents/document/234555
type EditProductRequest struct {
	ProductID          string                    `json:"product_id" validate:"required"`
	ProductName        string                    `json:"product_name" validate:"required"`
	Description        string                    `json:"description" validate:"required"`
	CategoryID         string                    `json:"category_id" validate:"required"`
	BrandID            string                    `json:"brand_id"`
	Images             []ImgID                   `json:"images,omitempty"`
	WarrantyPeriod     int                       `json:"warranty_period" example:"Need to choose among the candidate values provided by the platform: 1: 4 weeks" 2:"2 months" 3:"3 months" 4:"4 months" 5:"5 months" 6:"6 months" 7:"7 months" 8:"8 months" 9:"9 months" 10:"10 months" 11:"11 months" 12:"12 months" 13:"2 years" 14:"3 years" 15:"1 week" 16:"2 weeks" 17:"18 months" 18:"4 years" 19:"5 years" 20:"10 years" 21:"lifetime warranty"`
	WarrantyPolicy     string                    `json:"warranty_policy"`
	PackageLength      int                       `json:"package_length" example:"14. unit is cm"`
	PackageHeight      int                       `json:"package_height" example:"14. unit is cm"`
	PackageWidth       int                       `json:"package_width" example:"14. unit is cm"`
	PackageWeight      string                    `json:"package_weight" validate:"required" example:"14. unit is cm"`
	SizeChart          SizePic                   `json:"size_chart,omitempty"`
	Certifications     EditProductCertifications `json:"product_certifications"`
	IsCodOpen          bool                      `json:"is_cod_open,omitempty"`
	Skus               []SKU                     `json:"skus"`
	DeliveryServiceIDS string                    `json:"delivery_service_ids"`
	Attribute          EditAttributes            `json:"product_attributes,omitempty"`
}

type ProductSearchRequest struct {
	PageSize       int      `json:"page_size" validate:"min=1,max=100"`
	PageNumber     int      `json:"page_number" validate:"min=1"`
	SearchStatus   int      `json:"search_status,omitempty"`
	SellerSkuList  []string `json:"seller_sku_list,omitempty"`
	UpdateTimeFrom int64    `json:"update_time_from,omitempty"`
	UpdateTimeTo   int64    `json:"update_time_to,omitempty"`
	CreateTimeFrom int64    `json:"create_time_from,omitempty" `
	CreateTimeTo   int64    `json:"create_time_to,omitempty"`
}

type SkuItem struct {
	ID         string      `json:"id"`
	Price      Price       `json:"price"`
	SellerSku  string      `json:"seller_sku"`
	StockInfos []StockInfo `json:"stock_infos"`
}

type ProductItem struct {
	CreateTime  int       `json:"create_time"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	SaleRegions []string  `json:"sale_regions"`
	Skus        []SkuItem `json:"skus"`
	Status      int       `json:"status"`
	UpdateTime  int       `json:"update_time"`
}

type GetProductListData struct {
	Products []ProductItem `json:"products"`
	Total    int           `json:"total"`
}

type SKUPriceItem struct {
	ID            string `json:"id" validate:"required"`
	OriginalPrice string `json:"original_price" validate:"required"`
}

type UpdatePriceRequest struct {
	ProductID string         `json:"product_id" validate:"required"`
	Skus      []SKUPriceItem `json:"skus" validate:"min=1"`
}

type UpdatePriceData struct {
	FailedSKUIDs []string `json:"failed_sku_ids"`
}

type FailedSKUStock struct {
	ID                 string   `json:"id"`
	FailedWarehouseIDs []string `json:"failed_warehouse_ids"`
}

type SkuStockItem struct {
	ID         string      `json:"id"`
	StockInfos []StockInfo `json:"stock_infos"`
}

type UpdateStockRequest struct {
	ProductID string         `json:"product_id"`
	Skus      []SkuStockItem `json:"skus"`
}

type UpdateStockFailedSKU struct {
	FailedSKUs []FailedSKUStock `json:"failed_skus"`
}

type FailedProductIDs struct {
	FailedProductIDs []string `json:"failed_product_ids"`
}
