package tiktok_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ipfans/tiktok-sdk"
	"github.com/stretchr/testify/require"
)

func TestClient_GetOrderList(t *testing.T) {
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
				ak:    os.Getenv("AK"),
				query: tiktok.GetOrderListRequest{},
			},
			wantList: loadTestData(t, "testdata/order/list/no_filter.json"),
			wantErr:  false,
		},
		{
			name: "Request with valid ak and filter",
			args: args{
				ak: os.Getenv("AK"),
				query: tiktok.GetOrderListRequest{
					OrderStatus: 111,
				},
			},
			wantList: loadTestData(t, "testdata/order/list/filter_status_111.json"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := c.GetOrderList(tt.args.ak, tt.args.shopID, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetOrderList() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}

			b, _ := json.Marshal(gotList)
			jsonData := string(b)
			require.JSONEq(t, tt.wantList, jsonData)
		})
	}
}

func TestClient_GetOrderDetail(t *testing.T) {
	t.SkipNow()

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
				ak:      os.Getenv("AK"),
				shopID:  os.Getenv("SHOPID"),
				orderID: []string{"576467232825116133"},
			},
			wantDetail: loadTestData(t, "testdata/order/detail/id_576467232825116133.json"),
			wantErr:    false,
		},
		{
			name: "Request with valid ak and multi-id",
			args: args{
				ak:      os.Getenv("AK"),
				shopID:  os.Getenv("SHOPID"),
				orderID: []string{"576467232825116133", "576466303801001208"},
			},
			wantDetail: loadTestData(t, "testdata/order/list/filter_status_111.json"),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDetail, err := c.GetOrderDetail(tt.args.ak, tt.args.shopID, tt.args.orderID)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetOrderDetail() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}

			b, _ := json.Marshal(gotDetail)
			jsonData := string(b)
			require.JSONEq(t, tt.wantDetail, jsonData)
		})
	}
}
