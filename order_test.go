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
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	var response tiktok.OrdersList

	tests := loadTestData(t, "testdata/order_list.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			err := json.Unmarshal(tt.Args, &args)
			require.NoError(t, err)
			err = json.Unmarshal(tt.Want, &response)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			gotList, err := c.GetOrderList(context.TODO(), args.AccessToken, args.ShopID, tiktok.GetOrderListRequest{})
			if tt.WantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, response, gotList)
		})
	}
}

func TestClient_GetOrderDetail(t *testing.T) {

}
