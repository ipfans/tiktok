package tiktok

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestProductData(t *testing.T) {
	img01 := Image{
		ID: "ACC", Height: 1, Width: 2, ThumbUrlList: []string{"01.JPG", "02.JPG"}, UrlList: []string{"www.google.com", "twitter.com"},
	}
	data := &ProductData{
		ProductID:     "A",
		ProductStatus: 1,
		ProductName:   "A",
		CategoryList: []Category{
			{
				ID: "C1", ParentID: "003", LocalDisplayName: "LAKE-NIKE", IsLeaf: false,
			},
		},
		Brand: Brand{
			ID: "P", Name: "Nike",
		},
		Images: []Image{
			{
				ID: "ACC", Height: 1, Width: 2, ThumbUrlList: []string{"01.JPG", "02.JPG"}, UrlList: []string{"www.google.com", "twitter.com"},
			},
		},
		Video: Video{
			ID: "tokyo", Duration: 12.5, MediaType: "mkv", PostUrl: "www.yyets.com", VideoInfos: []VideoInfo{
				{Height: 1, Width: 2, FileHash: "29393ns9s23=", Format: "mpv", MainUrl: "shipping.com", Size: 1, UrlExpire: 3600, BackupUrl: "www.g.com", Bitrate: 1024},
			},
		},
		Description: "nike sb series",
		WarrantyPeriod: struct {
			WarrantyID          int    `json:"warranty_id"`
			WarrantyDescription string `json:"warranty_description"`
		}(struct {
			WarrantyID          int
			WarrantyDescription string
		}{WarrantyID: 996, WarrantyDescription: "996 day"}),

		WarrantyPolicy: "free return",
		PackageHeight:  100,
		PackageLength:  99,
		PackageWidth:   7,
		PackageWeight:  "100kg",
		Skus: []SKUData{
			{
				Price:      Price{OriginalPrice: "100", PriceIncludeVat: "1", Currency: "IDR"},
				SellerSku:  "ship-id-iphone13",
				ID:         "zh-fz-01",
				StockInfos: []StockInfo{{WarehouseID: "001", AvailableStock: 10}},
				SalesAttributes: []SalesAttr{
					{ID: "001", Name: "ABC", ValueID: "9", ValueName: "COLOR", SkuImg: img01},
				},
			},
		},
		ProductCertifications: []Certification{
			{ID: "1a", Title: "disney", Files: []File{
				{ID: "1", List: []string{"P", "Q"}, Name: "QQ", Type: "PDF"},
			}},
		},
		SizeChart:         img01,
		IsCodOpen:         false,
		ProductAttributes: []ProductAttribute{{ID: "12", Name: "pptv", Values: []Value{{ID: "LL", Name: "waka"}}}},
		QcReasons:         []QCReason{{Reason: "Illegal", SubReasons: []string{"R18", "AV"}}},
		UpdateTime:        1652684619,
		CreateTime:        1652684619,
		DeliveryServices: []DeliveryService{
			{
				DeliveryServiceStatus: false,
				DeliveryServiceID:     12,
				DeliveryOptionName:    "J&T Express",
			},
		},
	}
	d, _ := json.Marshal(data)
	fmt.Println(string(d))
}
