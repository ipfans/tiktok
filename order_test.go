package tiktok_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestClient_GetOrderList(t *testing.T) {
	var response tiktok.OrdersList
	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/order/order_list.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var args struct {
				AppKey      string `json:"app_key"`
				AppSecret   string `json:"app_secret"`
				AccessToken string `json:"access_token"`
				ShopID      string `json:"shop_id"`
			}
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			gotList, err := c.GetOrderList(context.TODO(),
				tiktok.Param{AccessToken: args.AccessToken, ShopID: args.ShopID},
				tiktok.GetOrderListRequest{},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, response, gotList)
		})
	}
}

func TestClient_GetOrderDetail(t *testing.T) {
	var response tiktok.OrdersList
	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/order/order_detail.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var args struct {
				AppKey      string   `json:"app_key"`
				AppSecret   string   `json:"app_secret"`
				AccessToken string   `json:"access_token"`
				ShopID      string   `json:"shop_id"`
				OrderID     []string `json:"order_id"`
			}
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var orderDetail tiktok.OrderDetail
			err := json.Unmarshal([]byte(tt.Want), &orderDetail)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			gotList, err := c.GetOrderDetail(context.TODO(),
				tiktok.Param{AccessToken: args.AccessToken, ShopID: args.ShopID},
				args.OrderID)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, orderDetail, gotList.OrderList[0])
		})
	}
}

func TestClient_ShipOrder(t *testing.T) {
	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/order/ship_order.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var args struct {
				AppKey      string `json:"app_key"`
				AppSecret   string `json:"app_secret"`
				AccessToken string `json:"access_token"`
				ShopID      string `json:"shop_id"`
				OrderID     string `json:"order_id"`
			}
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, nil)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			_, err = c.ShipOrder(context.TODO(),
				tiktok.Param{AccessToken: args.AccessToken, ShopID: args.ShopID},
				tiktok.ShipOrderRequest{OrderID: args.OrderID},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestClient_CancelOrder(t *testing.T) {
	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/order/cancel_order.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var args struct {
				AppKey          string `json:"app_key"`
				AppSecret       string `json:"app_secret"`
				AccessToken     string `json:"access_token"`
				ShopID          string `json:"shop_id"`
				OrderID         string `json:"order_id"`
				CancelReasonKey string `json:"cancel_reason_key"`
			}
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, nil)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			_, err = c.CancelOrder(context.TODO(),
				tiktok.Param{AccessToken: args.AccessToken, ShopID: args.ShopID},
				tiktok.CancelOrderRequest{OrderID: args.OrderID, CancelReasonKey: args.CancelReasonKey},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
		})
	}
}
