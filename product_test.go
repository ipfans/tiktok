package tiktok_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCategory(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_category.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.CategoryList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetCategory(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_GetAttribute(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		CategoryID  string `json:"category_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_attribute.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.AttributeList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetAttribute(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.CategoryID,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_GetCategoryRule(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		CategoryID  string `json:"category_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_category_rule.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.CategoryRules
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetCategoryRule(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.CategoryID,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_GetBrand(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_brand.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.BrandList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetBrand(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_UploadImg(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		File        string `json:"file"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/upload_img.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.ImageInfo
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)
			r := strings.NewReader(args.File)

			img, err := c.UploadImgReader(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				tiktok.ImgSceneAttributeImage, r,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, img)
		})
	}
}

func TestClient_UploadFile(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		File        string `json:"file"`
		Name        string `json:"name"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/upload_file.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.FileInfo
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			info, err := c.UploadFile(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.Name, []byte(args.File),
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, info)
		})
	}
}

// TODO: Fix those tests.
// func TestClient_CreateProduct(t *testing.T) {
// 	var args struct {
// 		AppKey      string          `json:"app_key"`
// 		AppSecret   string          `json:"app_secret"`
// 		AccessToken string          `json:"access_token"`
// 		ShopID      string          `json:"shop_id"`
// 		RequestJSON json.RawMessage `json:"request_json"`
// 	}
// 	var want tiktok.Product
// 	mockTests(t, "testdata/product/create_product.json", &args, &want, func() (got interface{}, err error) {
// 		c, err := tiktok.New(args.AppKey, args.AppSecret)
// 		require.NoError(t, err)
// 		var req tiktok.CreateProductRequest
// 		err = json.Unmarshal(args.RequestJSON, &req)
// 		require.NoError(t, err)
// 		product, err := c.CreateProduct(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, req)
// 		got = &product
// 		return
// 	})
// }

// func TestClient_EditProduct(t *testing.T) {
// 	var args struct {
// 		AppKey      string          `json:"app_key"`
// 		AppSecret   string          `json:"app_secret"`
// 		AccessToken string          `json:"access_token"`
// 		ShopID      string          `json:"shop_id"`
// 		RequestJSON json.RawMessage `json:"request_json"`
// 	}
// 	var want tiktok.Product
// 	mockTests(t, "testdata/product/edit_product.json", &args, &want, func() (got interface{}, err error) {
// 		c, err := tiktok.New(args.AppKey, args.AppSecret)
// 		require.NoError(t, err)
// 		var req tiktok.CreateProductRequest
// 		err = json.Unmarshal(args.RequestJSON, &req)
// 		require.NoError(t, err)
// 		product, err := c.CreateProduct(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, req)
// 		got = &product
// 		return
// 	})
// }

func TestClient_GetProductList(t *testing.T) {
	var args struct {
		AppKey      string          `json:"app_key"`
		AppSecret   string          `json:"app_secret"`
		AccessToken string          `json:"access_token"`
		ShopID      string          `json:"shop_id"`
		RequestJSON json.RawMessage `json:"request_json"`
	}
	var want tiktok.ProductSearchList
	mockTests(t, "testdata/product/get_product_list.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		var req tiktok.ProductSearchRequest
		err = json.Unmarshal(args.RequestJSON, &req)
		require.NoError(t, err)
		product, err := c.GetProductList(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, req)
		got = &product
		return
	})
}

func TestClient_GetProductDetail(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		ProductID   string `json:"product_id"`
	}
	var want tiktok.Product
	mockTests(t, "testdata/product/get_product_detail.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		product, err := c.GetProductDetail(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, args.ProductID)
		got = &product
		return
	})
}

func TestClient_UpdatePrice(t *testing.T) {
	var args struct {
		AppKey      string          `json:"app_key"`
		AppSecret   string          `json:"app_secret"`
		AccessToken string          `json:"access_token"`
		ShopID      string          `json:"shop_id"`
		RequestJSON json.RawMessage `json:"request_json"`
	}
	var want tiktok.UpdatePriceFailedSKU
	mockTests(t, "testdata/product/update_price.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		var req tiktok.UpdatePriceRequest
		err = json.Unmarshal(args.RequestJSON, &req)
		require.NoError(t, err)
		product, err := c.UpdatePrice(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, req)
		got = &product
		return
	})
}

func TestClient_UpdateStock(t *testing.T) {
	var args struct {
		AppKey      string          `json:"app_key"`
		AppSecret   string          `json:"app_secret"`
		AccessToken string          `json:"access_token"`
		ShopID      string          `json:"shop_id"`
		RequestJSON json.RawMessage `json:"request_json"`
	}
	var want tiktok.UpdateStockFailedSKU
	mockTests(t, "testdata/product/update_stock.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		var req tiktok.UpdateStockRequest
		err = json.Unmarshal(args.RequestJSON, &req)
		require.NoError(t, err)
		product, err := c.UpdateStock(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, req)
		got = &product
		return
	})
}

func TestClient_DeactivateProducts(t *testing.T) {
	var args struct {
		AppKey      string   `json:"app_key"`
		AppSecret   string   `json:"app_secret"`
		AccessToken string   `json:"access_token"`
		ShopID      string   `json:"shop_id"`
		ProductID   []string `json:"product_id"`
	}
	var want tiktok.FailedProductIDs
	mockTests(t, "testdata/product/deactivate_products.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		product, err := c.DeactivateProducts(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, args.ProductID)
		got = &product
		return
	})
}

func TestClient_DeleteProducts(t *testing.T) {
	var args struct {
		AppKey      string   `json:"app_key"`
		AppSecret   string   `json:"app_secret"`
		AccessToken string   `json:"access_token"`
		ShopID      string   `json:"shop_id"`
		ProductID   []string `json:"product_id"`
	}
	var want tiktok.FailedProductIDs
	mockTests(t, "testdata/product/delete_products.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		product, err := c.DeleteProducts(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, args.ProductID)
		got = &product
		return
	})
}

func TestClient_RecoverProduct(t *testing.T) {
	var args struct {
		AppKey      string   `json:"app_key"`
		AppSecret   string   `json:"app_secret"`
		AccessToken string   `json:"access_token"`
		ShopID      string   `json:"shop_id"`
		ProductID   []string `json:"product_id"`
	}
	var want tiktok.FailedProductIDs
	mockTests(t, "testdata/product/recover_product.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		product, err := c.RecoverProduct(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, args.ProductID)
		got = &product
		return
	})
}

func TestClient_ActivateProduct(t *testing.T) {
	var args struct {
		AppKey      string   `json:"app_key"`
		AppSecret   string   `json:"app_secret"`
		AccessToken string   `json:"access_token"`
		ShopID      string   `json:"shop_id"`
		ProductID   []string `json:"product_id"`
	}
	var want tiktok.FailedProductIDs
	mockTests(t, "testdata/product/activate_product.json", &args, &want, func() (got interface{}, err error) {
		c, err := tiktok.New(args.AppKey, args.AppSecret)
		require.NoError(t, err)
		product, err := c.ActivateProduct(context.TODO(), tiktok.Param{args.AccessToken, args.ShopID}, args.ProductID)
		got = &product
		return
	})
}
