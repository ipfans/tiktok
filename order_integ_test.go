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

func TestClient_GetOrderList_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.GetOrderListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: "Request with valid ak",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  tiktok.GetOrderListRequest{},
			},
			wantErr: false,
		},
		{
			name: "Request with valid ak and filter",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.GetOrderListRequest{
					OrderStatus: 111,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := c.GetOrderList(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetOrderList() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, gotList.NextCursor)
			b, _ := json.Marshal(gotList)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetOrderDetail_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak      string
		shopID  string
		orderID []string
	}
	tests := []struct {
		name       string
		args       args
		wantDetail string
		wantErr    bool
	}{
		{
			name: "Request with valid ak and id",
			args: args{
				ak:      os.Getenv(_AK_KEY),
				shopID:  os.Getenv(_SHOP_KEY),
				orderID: strings.Split(os.Getenv("ORDER"), ","),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDetail, err := c.GetOrderDetail(context.TODO(), tiktok.Param{tt.args.ak, tt.args.shopID}, tt.args.orderID)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetOrderDetail() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}

			b, _ := json.Marshal(gotDetail)
			jsonData := string(b)
			t.Log(jsonData)
			require.Contains(t, jsonData, tt.args.orderID[0])
		})
	}
}
