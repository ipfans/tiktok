package tiktok_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestSearchPreCombinePkg(t *testing.T) {
	var args struct {
		AppKey      string                            `json:"app_key"`
		AppSecret   string                            `json:"app_secret"`
		AccessToken string                            `json:"access_token"`
		ShopID      string                            `json:"shop_id"`
		Req         tiktok.SearchPreCombinePkgRequest `json:"req"`
	}

	var response tiktok.SearchPreCombinePkgData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[12]SearchPreCombinePkg.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.SearchPreCombinePkgData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.SearchPreCombinePkg(context.Background(), tiktok.Param{
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

func TestConfirmPreCombinePkg(t *testing.T) {
	var args struct {
		AppKey      string                             `json:"app_key"`
		AppSecret   string                             `json:"app_secret"`
		AccessToken string                             `json:"access_token"`
		ShopID      string                             `json:"shop_id"`
		Req         tiktok.ConfirmPreCombinePkgRequest `json:"req"`
	}

	var response tiktok.PreCombinePkgData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[11]ConfirmPreCombinePkg.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.PreCombinePkgData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.ConfirmPreCombinePkg(context.Background(), tiktok.Param{
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

func TestRemovePackageOrder(t *testing.T) {
	var args struct {
		AppKey      string                           `json:"app_key"`
		AppSecret   string                           `json:"app_secret"`
		AccessToken string                           `json:"access_token"`
		ShopID      string                           `json:"shop_id"`
		Req         tiktok.RemovePackageOrderRequest `json:"req"`
	}

	var response tiktok.PreCombinePkgData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[10]RemovePackageOrder.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.PreCombinePkgData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.RemovePackageOrder(context.Background(), tiktok.Param{
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

func TestGetPackagePickupConfig(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		Req         struct {
			PackageID string `json:"package_id"`
		}
	}

	var response tiktok.GetPackagePickupConfigData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[9]GetPackagePickupConfig.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.GetPackagePickupConfigData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetPackagePickupConfig(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			}, tiktok.PackageIDRequest{
				PackageID: args.Req.PackageID,
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

func TestShipPackage(t *testing.T) {
	var args struct {
		AppKey      string                    `json:"app_key"`
		AppSecret   string                    `json:"app_secret"`
		AccessToken string                    `json:"access_token"`
		ShopID      string                    `json:"shop_id"`
		Req         tiktok.ShipPackageRequest `json:"req"`
	}

	var response tiktok.ShipPackageData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[8]ShipPackage.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.ShipPackageData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)
			res, err := c.ShipPackage(context.Background(), tiktok.Param{
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

func TestSearchPackage(t *testing.T) {
	var args struct {
		AppKey      string                    `json:"app_key"`
		AppSecret   string                    `json:"app_secret"`
		AccessToken string                    `json:"access_token"`
		ShopID      string                    `json:"shop_id"`
		Req         tiktok.ShipPackageRequest `json:"req"`
	}

	var response tiktok.SearchPackageData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[7]SearchPackage.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.SearchPackageData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)
			res, err := c.SearchPackage(context.Background(), tiktok.Param{
				AccessToken: args.AccessToken, ShopID: args.ShopID,
			}, tiktok.SearchPackageRequest{})
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

func TestGetPackageDetail(t *testing.T) {
	var args struct {
		AppKey      string                  `json:"app_key"`
		AppSecret   string                  `json:"app_secret"`
		AccessToken string                  `json:"access_token"`
		ShopID      string                  `json:"shop_id"`
		Req         tiktok.PackageIDRequest `json:"req"`
	}

	var response tiktok.GetPackageDetailData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[6]GetPackageDetail.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.GetPackageDetailData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)
			res, err := c.GetPackageDetail(context.Background(), tiktok.Param{
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

func TestGetPackageShippingInfo(t *testing.T) {
	var args struct {
		AppKey      string                  `json:"app_key"`
		AppSecret   string                  `json:"app_secret"`
		AccessToken string                  `json:"access_token"`
		ShopID      string                  `json:"shop_id"`
		Req         tiktok.PackageIDRequest `json:"req"`
	}

	var response tiktok.GetPackageShippingInfoData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[5]GetPackageShippingInfo.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.GetPackageShippingInfoData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetPackageShippingInfo(context.Background(), tiktok.Param{
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

func TestUpdatePackageShippingInfo(t *testing.T) {
	var args struct {
		AppKey      string                                  `json:"app_key"`
		AppSecret   string                                  `json:"app_secret"`
		AccessToken string                                  `json:"access_token"`
		ShopID      string                                  `json:"shop_id"`
		Req         tiktok.UpdatePackageShippingInfoRequest `json:"req"`
	}

	var response tiktok.UpdatePackageShippingInfoData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[4]UpdatePackageShippingInfo.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.UpdatePackageShippingInfoData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.UpdatePackageShippingInfo(context.Background(), tiktok.Param{
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

func TestGetPackageShippingDocument(t *testing.T) {
	var args struct {
		AppKey      string                                   `json:"app_key"`
		AppSecret   string                                   `json:"app_secret"`
		AccessToken string                                   `json:"access_token"`
		ShopID      string                                   `json:"shop_id"`
		Req         tiktok.GetPackageShippingDocumentRequest `json:"req"`
	}

	var response tiktok.GetPackageShippingDocumentData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[3]GetPackageShippingDocument.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.GetPackageShippingDocumentData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.GetPackageShippingDocument(context.Background(), tiktok.Param{
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

func TestVerifyOrderSplit(t *testing.T) {
	var args struct {
		AppKey      string                         `json:"app_key"`
		AppSecret   string                         `json:"app_secret"`
		AccessToken string                         `json:"access_token"`
		ShopID      string                         `json:"shop_id"`
		Req         tiktok.VerifyOrderSplitRequest `json:"req"`
	}

	var response tiktok.VerifyOrderSplitData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[2]VerifyOrderSplit.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)

			var ans tiktok.VerifyOrderSplitData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.VerifyOrderSplit(context.Background(), tiktok.Param{
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

func TestConfirmOrderSplit(t *testing.T) {
	var args struct {
		AppKey      string                          `json:"app_key"`
		AppSecret   string                          `json:"app_secret"`
		AccessToken string                          `json:"access_token"`
		ShopID      string                          `json:"shop_id"`
		Req         tiktok.ConfirmOrderSplitRequest `json:"req"`
	}

	var response tiktok.ConfirmOrderSplitData
	restore := mockTime()
	defer restore()
	tests := loadTestData(t, "testdata/fulfillment/[1]ConfirmOrderSplit.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, &response)
			var ans tiktok.ConfirmOrderSplitData
			err := json.Unmarshal(tt.Want, &ans)
			require.NoError(t, err)
			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			res, err := c.ConfirmOrderSplit(context.Background(), tiktok.Param{
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
