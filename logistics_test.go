package tiktok_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetShippingInfo(t *testing.T) {
	var args struct {
		AppKey      string            `json:"app_key"`
		AppSecret   string            `json:"app_secret"`
		AccessToken string            `json:"access_token"`
		ShopID      string            `json:"shop_id"`
		Req         tiktok.OrderIDReq `json:"req"`
	}

	var response tiktok.LogisticsGetShippingInfoData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/logistics/[1]GetShippingInfo.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.LogisticsGetShippingInfoData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetShippingInfo(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			}, args.Req)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, ans, res)
		})
	}
}

func TestUpdateShippingInfo(t *testing.T) {
	var args struct {
		AppKey      string                       `json:"app_key"`
		AppSecret   string                       `json:"app_secret"`
		AccessToken string                       `json:"access_token"`
		ShopID      string                       `json:"shop_id"`
		Req         tiktok.UpdateShippingInfoReq `json:"req"`
	}

	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/logistics/[2]UpdateShippingInfo.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, nil)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			err = c.UpdateShippingInfo(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			}, args.Req)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetShippingDocument(t *testing.T) {
	var args struct {
		AppKey      string                            `json:"app_key"`
		AppSecret   string                            `json:"app_secret"`
		AccessToken string                            `json:"access_token"`
		ShopID      string                            `json:"shop_id"`
		Req         tiktok.GetShippingDocumentRequest `json:"req"`
	}

	var response tiktok.GetShippingDocumentData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/logistics/[3]GetShippingDocument.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.GetShippingDocumentData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetShippingDocument(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			}, args.Req)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, ans, res)
		})
	}
}

func TestGetAddressList(t *testing.T) {
	var args struct {
		AppKey      string                       `json:"app_key"`
		AppSecret   string                       `json:"app_secret"`
		AccessToken string                       `json:"access_token"`
		ShopID      string                       `json:"shop_id"`
		Req         tiktok.GetAddressListRequest `json:"req"`
	}

	var response tiktok.GetAddressListData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/logistics/[4]GetAddressList(deprecated).json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.GetAddressListData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetAddressList(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			}, args.Req)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, ans, res)
		})
	}
}

func TestGetWarehouseList(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	var response tiktok.GetWarehouseListData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/logistics/[5]GetWarehouseList.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.GetWarehouseListData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetWarehouseList(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			})
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, ans, res)
		})
	}
}

func TestGetShippingProvider(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	var response tiktok.GetShippingProviderData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/logistics/[6]GetShippingProvider.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.GetShippingProviderData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetShippingProvider(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			})
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, ans, res)
		})
	}
}
