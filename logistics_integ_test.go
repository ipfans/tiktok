//go:build integration

package tiktok_test

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/stretchr/testify/require"
)

func TestClient_GetShippingInfo_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.OrderIDReq
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetShippingInfo Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.OrderIDReq{
					OrderID: os.Getenv(_ORDER_KEY),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetShippingInfo(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetShippingInfo() error = %v, wantErr %v", err, tt.wantErr)
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

func TestClient_UpdateShippingInfo_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.UpdateShippingInfoReq
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " UpdateShippingInfo Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.UpdateShippingInfoReq{
					OrderID:        "576500005218781420",
					TrackingNumber: "JX0963674147",
					ProviderID:     "6841743441349706241",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.UpdateShippingInfo(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.UpdateShippingInfo() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
		})
	}
}

func TestClient_GetShippingDocument_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.GetShippingDocumentRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetShippingDocument Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.GetShippingDocumentRequest{
					OrderID:      strconv.Itoa(_ORDERID),
					DocumentType: "SHIPPING_LABEL",
					DocumentSize: "A5",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetShippingDocument(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetShippingDocument() error = %v, wantErr %v", err, tt.wantErr)
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

func TestClient_GetWarehouseList_Integration(t *testing.T) {
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
			name: " GetWarehouseList Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetWarehouseList(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID})
			require.Equal(t, tt.wantErr, err != nil, "Client.GetWarehouseList() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}

			require.NotEmpty(t, ans.WarehouseList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetShippingProvider_Integration(t *testing.T) {
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
			name: " GetShippingProvider Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetShippingProvider(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID})
			require.Equal(t, tt.wantErr, err != nil, "Client.GetShippingProvider() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.DeliveryOptionList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetSubscribedDelivery_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.GetSubscribedDeliveryRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetSubscribedDelivery Integration Test ",
			args: args{
				ak:     os.Getenv("AK"),
				shopID: os.Getenv("SHOP"),
				query:  tiktok.GetSubscribedDeliveryRequest{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetSubscribedDelivery(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetSubscribedDelivery() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans.WarehouseList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}
