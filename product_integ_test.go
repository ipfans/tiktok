//go:build integration

package tiktok_test

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCategory_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetCategory Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetCategory(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID})
			require.Equal(t, tt.wantErr, err != nil, "Client.GetCategory() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.CategoryList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetAttribute_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetAttribute Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  _CATEGORYID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetAttribute(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetAttribute() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.Attributes)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetCategoryRule_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetCategoryRule Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  "600100",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetCategoryRule(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetCategoryRule() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.CategoryRules)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetBrand_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetBrand Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetBrand(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID})
			require.Equal(t, tt.wantErr, err != nil, "Client.GetBrand() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.BrandList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

// TODO
func TestClient_UploadImg_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		File   string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " UploadImgReader Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				File:   "testdata/refresh_token.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.UploadImgReader(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tiktok.ImgSceneAttributeImage, strings.NewReader(tt.args.File))
			require.Equal(t, tt.wantErr, err != nil, "Client.UploadImgReader() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.ImgURL)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_UploadFile_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		File   string
		Name   string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " UploadFile Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				File:   "",
				Name:   "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.UploadFile(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.File, []byte(tt.args.Name))
			require.Equal(t, tt.wantErr, err != nil, "Client.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.FileURL)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_CreateProductRequest_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.CreateProductRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " CreateProductRequest Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.CreateProductRequest{
					ProductName: "teapot-jingdezheng",
					Description: "<ul><li>It is recommended to avoid using Chinese because the copy will be displayed to local users</li></ul><img src=\\\"https://p19-oec-va.ibyteimg.com/tos-maliva-i-o3syd03w52-us/8de4c52c078042589e427c681ca10d0e~tplv-o3syd03w52-origin-jpeg.jpeg?\\\" width=\\\"1920\\\" height=\\\"1080\\\">",
					CategoryID:  "600100",
					BrandID:     "", // brand_id
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.CreateProduct(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}

			require.NotEmpty(t, ans)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_EditProduct_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.EditProductRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " EditProduct Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  tiktok.EditProductRequest{
					// TODO
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.EditProduct(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.EditProduct() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans)
			require.NotEmpty(t, ans.Skus)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetProductList_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.ProductSearchRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetProductList Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.ProductSearchRequest{
					PageSize:   _PAGESIZE,
					PageNumber: 1,
					// CreateTimeFrom: 1651204520,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetProductList(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetProductList() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			if ans.Total == 0 {
				require.Empty(t, ans.Products)
				return
			}
			require.NotEmpty(t, ans.Products)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetProductDetail_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetProductDetail Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  "1729428817744857665",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetProductDetail(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetProductDetail() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.Skus)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_UpdatePrice_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.UpdatePriceRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " UpdatePrice Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.UpdatePriceRequest{
					ProductID: "1729428817744857665",
					Skus: []struct {
						ID            string `json:"id" validate:"required"`
						OriginalPrice string `json:"original_price" validate:"required"`
					}{
						{
							"1729428819288164929", "9999",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.UpdatePrice(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.UpdatePrice() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			// If update success, FailedSKUIDs should be empty
			require.Empty(t, ans.FailedSKUIDs)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_UpdateStock_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.UpdateStockRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " UpdateStock Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.UpdateStockRequest{
					ProductID: "1729428817744857665",
					Skus: []struct {
						ID         string `json:"id"`
						StockInfos []struct {
							AvailableStock int    `json:"available_stock" validate:"min=0,max=99999"`
							WarehouseID    string `json:"warehouse_id"`
						} `json:"stock_infos"`
					}([]struct {
						ID         string `json:"id"`
						StockInfos []struct {
							AvailableStock int    `json:"available_stock"`
							WarehouseID    string `json:"warehouse_id"`
						} `json:"stock_infos"`
					}{
						{
							ID: "1729428819288164929",
							StockInfos: []struct {
								AvailableStock int    `json:"available_stock"`
								WarehouseID    string `json:"warehouse_id"`
							}{
								{
									AvailableStock: 2,
									WarehouseID:    "7054379554541963010",
								},
							},
						},
					}),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.UpdateStock(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.UpdateStock() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			// if update success, failed skus is empty
			require.Empty(t, ans)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

// TestClient_DeleteProducts_Integration should be combined with TestClient_RecoverProduct_Integration
func TestClient_DeleteProducts_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  []string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " DeleteProducts Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  []string{"1729428001967868481"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.DeleteProducts(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.DeleteProducts() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_RecoverProduct_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  []string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " RecoverProduct Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  []string{"1729428001967868481"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.RecoverProduct(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.RecoverProduct() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.FailedProductIDs)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

// TestClient_DeactivateProducts_Integration should be combined with  TestClient_ActivateProduct_Integration.
func TestClient_DeactivateProducts_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  []string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " DeactivateProducts Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  []string{"1729428001967868481"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.DeactivateProducts(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.DeactivateProducts() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.FailedProductIDs)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_ActivateProduct_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  []string
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " ActivateProduct Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  []string{"1729428001967868481"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.ActivateProduct(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.ActivateProduct() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.FailedProductIDs)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}
