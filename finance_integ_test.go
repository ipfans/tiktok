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

func TestClient_SearchSettlements_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.SearchSettlementsRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " SearchSettlements Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.SearchSettlementsRequest{
					PageSize:        _PAGESIZE,
					SortType:        _DESC_TYPE,
					RequestTimeTo:   1691224452,
					RequestTimeFrom: 1558152466,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.SearchSettlements(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.SearchSettlements() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			if len(ans.SettlementList) == 0 {
				return
			}
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetOrderSettlements_Integration(t *testing.T) {
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
			name: " GetOrderSettlements Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  _ORDERIDStr,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetOrderSettlements(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetOrderSettlements() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans.SettlementList)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_SearchTransactions_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.SearchTransactionsRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " SearchTransactions Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.SearchTransactionsRequest{
					PageSize: _PAGESIZE,
					// Withdraw:1 Settle:2 Transfer:3 Reverse:4
					TransactionType: []int{1, 2, 3, 4},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.SearchTransactions(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.SearchTransactions() error = %v, wantErr %v", err, tt.wantErr)
		})
	}
}
