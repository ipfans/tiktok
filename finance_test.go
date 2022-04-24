package tiktok_test

import (
	"context"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestClient_SearchSettlements(t *testing.T) {
	var args struct {
		AppKey          string `json:"app_key"`
		AppSecret       string `json:"app_secret"`
		AccessToken     string `json:"access_token"`
		ShopID          string `json:"shop_id"`
		RequestTimeFrom int    `json:"request_time_from"`
		RequestTimeTo   int    `json:"request_time_to"`
		PageSize        int    `json:"page_size"`
		Cursor          string `json:"cursor"`
		SortType        int    `json:"sort_type"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/finance/search_settlements.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.SettlementsList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			shops, err := c.SearchSettlements(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				tiktok.SearchSettlementsRequest{
					RequestTimeFrom: args.RequestTimeFrom,
					RequestTimeTo:   args.RequestTimeTo,
					PageSize:        args.PageSize,
					Cursor:          args.Cursor,
					SortType:        args.SortType,
				},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, shops)
		})
	}
}

func TestClient_GetOrderSettlements(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		OrderID     string `json:"order_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/finance/search_transactions.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.SettlementsList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			shops, err := c.GetOrderSettlements(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.OrderID,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, shops)
		})
	}
}

func TestClient_SearchTransactions(t *testing.T) {
	var args struct {
		AppKey          string `json:"app_key"`
		AppSecret       string `json:"app_secret"`
		AccessToken     string `json:"access_token"`
		ShopID          string `json:"shop_id"`
		RequestTimeFrom int    `json:"request_time_from"`
		RequestTimeTo   int    `json:"request_time_to"`
		TransactionType []int  `json:"transaction_type" validate:"required"`
		PageSize        int    `json:"page_size"`
		Offset          int    `json:"offset" validate:"required,gte=0,lte=1000"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/finance/search_transactions.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.TransactionsList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			shops, err := c.SearchTransactions(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				tiktok.SearchTransactionsRequest{
					RequestTimeFrom: args.RequestTimeFrom,
					RequestTimeTo:   args.RequestTimeTo,
					TransactionType: args.TransactionType,
					PageSize:        args.PageSize,
					Offset:          args.Offset,
				},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, shops)
		})
	}
}
