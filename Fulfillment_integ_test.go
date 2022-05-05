//go:build integration

package tiktok_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/stretchr/testify/require"
)

func TestClient_SearchPreCombinePkg_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.SearchPreCombinePkgRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " SearchPreCombinePkg Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.SearchPreCombinePkgRequest{
					PageSize: _PAGESIZE,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.SearchPreCombinePkg(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.SearchPreCombinePkg() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Less(t, len(ans.PreCombinePkgList), _PAGESIZE)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_ConfirmPreCombinePkg_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.ConfirmPreCombinePkgRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " ConfirmPreCombinePkg Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.ConfirmPreCombinePkgRequest{
					PreCombinePkgList: []struct {
						PreCombinePkgID string   `json:"pre_combine_pkg_id" validate:"required"`
						OrderIDList     []string `json:"order_id_list"  validate:"omitempty"`
					}{{
						"123", []string{_FakeOrderID},
					}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.ConfirmPreCombinePkg(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.ConfirmPreCombinePkg() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.PackageList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_RemovePackageOrder_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.RemovePackageOrderRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " RemovePackageOrder Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.RemovePackageOrderRequest{
					PackageID:   "",
					OrderIDList: []string{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.RemovePackageOrder(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.RemovePackageOrder() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.PackageList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetPackagePickupConfig_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.PackageIDRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetPackagePickupConfig Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.PackageIDRequest{
					PackageID: _PACKAGEID,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetPackagePickupConfig(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetPackagePickupConfig() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.PickUpTimeList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_ShipPackage_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.ShipPackageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " ShipPackage Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.ShipPackageRequest{PackageID: "1152946689385597164", SelfShipment: struct {
					TrackingNumber     string `json:"tracking_number"`
					ShippingProviderID string `json:"shipping_provider_id"`
				}(struct {
					TrackingNumber     string
					ShippingProviderID string
				}{TrackingNumber: _TRACKINGNUMBER, ShippingProviderID: _ShippingProviderID})},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.ShipPackage(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.ShipPackage() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.FailPackages)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_SearchPackage_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.SearchPackageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " SearchPackage Integration Test ",
			args: args{
				ak:     os.Getenv("AK"),
				shopID: os.Getenv("SHOP"),
				query: tiktok.SearchPackageRequest{
					PageSize: _PAGESIZE,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.SearchPackage(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.SearchPackage() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.PackageList)
			require.Less(t, len(ans.PackageList), tt.args.query.PageSize)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetPackageDetail_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.PackageIDRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetPackageDetail Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.PackageIDRequest{
					PackageID: _PACKAGEID,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetPackageDetail(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetPackageDetail() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.OrderInfoList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetPackageShippingInfo_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.PackageIDRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetPackageShippingInfo Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.PackageIDRequest{
					PackageID: _PACKAGEID,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetPackageShippingInfo(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetPackageShippingInfo() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.TrackingInfoList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_UpdatePackageShippingInfo_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.UpdatePackageShippingInfoRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " UpdatePackageShippingInfoRequest Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.UpdatePackageShippingInfoRequest{
					PackageID:      "1152946689385597164",
					TrackingNumber: "JX0963674147",
					ProviderID:     "6841743441349706241",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.UpdatePackageShippingInfo(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.UpdatePackageShippingInfo() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}

			require.False(t, ans.UpdateSuccess)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetPackageShippingDocument_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.GetPackageShippingDocumentRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetPackageShippingDocument Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.GetPackageShippingDocumentRequest{
					PackageID:    _PACKAGEID,
					DocumentType: 3,
					DocumentSize: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetPackageShippingDocument(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetPackageShippingDocument() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.DocURL)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_VerifyOrderSplit_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.VerifyOrderSplitRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " VerifyOrderSplit Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.VerifyOrderSplitRequest{
					OrderIDList: []int64{_ORDERID},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.VerifyOrderSplit(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.VerifyOrderSplit() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.FailList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_ConfirmOrderSplit_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.ConfirmOrderSplitRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " ConfirmOrderSplit Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.ConfirmOrderSplitRequest{
					OrderID: _ORDERID,
					SplitGroup: []struct {
						PreSplitPkgID   int     `json:"pre_split_pkg_id" validate:"required"`
						OrderLineIDList []int64 `json:"order_line_id_list" validate:"min=1"`
					}{
						{123, []int64{576474865727932968}},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.ConfirmOrderSplit(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.ConfirmOrderSplit() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.ConfirmResultList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}
