package tiktok_test

import (
	"context"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestClient_ConfirmReverse(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		OrderID     string `json:"order_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/reverseorder/confirm_reverse.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, nil)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			err = c.ConfirmReverse(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.OrderID,
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

func TestClient_RejectReverse(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		OrderID     string `json:"order_id"`
		ReasonKey   string `json:"reason_key"`
		Comments    string `json:"comments"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/reverseorder/reject_reverse.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			setupMock(t, tt, &args, nil)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			err = c.RejectReverse(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				tiktok.RejectReverseRequest{
					ReverseOrderID:         args.OrderID,
					ReverseRejectReasonKey: args.ReasonKey,
					ReverseRejectComments:  args.Comments,
				},
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

func TestClient_GetReverseList(t *testing.T) {
	var args struct {
		AppKey      string                       `json:"app_key"`
		AppSecret   string                       `json:"app_secret"`
		AccessToken string                       `json:"access_token"`
		ShopID      string                       `json:"shop_id"`
		Req         tiktok.GetReverseListRequest `json:"req"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/reverseorder/get_reverse_list.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var list tiktok.ReverseOrdersList
			setupMock(t, tt, &args, &list)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			gotList, err := c.GetReverseList(context.TODO(),
				tiktok.Param{AccessToken: args.AccessToken, ShopID: args.ShopID}, args.Req,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, list, gotList)
		})
	}
}

func TestClient_GetReverseReason(t *testing.T) {
	var args struct {
		AppKey         string `json:"app_key"`
		AppSecret      string `json:"app_secret"`
		AccessToken    string `json:"access_token"`
		ShopID         string `json:"shop_id"`
		UpdateTimeFrom int    `json:"update_time_from"`
		UpdateTimeTo   int    `json:"update_time_to"`
		ReverseType    int    `json:"reverse_type"`
		SortBy         int    `json:"sort_by"`
		SortType       int    `json:"sort_type"`
		Offset         int    `json:"offset"`
		Size           int    `json:"size"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/reverseorder/get_reverse_reason.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var list tiktok.ReverseReasonList
			setupMock(t, tt, &args, &list)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			gotList, err := c.GetReverseReason(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				tiktok.GetReverseReasonRequest{},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, list, gotList)
		})
	}
}
