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

func TestClient_ConfirmReverse_Integration(t *testing.T) {
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
			name: " ConfirmReverse Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query:  "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.ConfirmReverse(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.ConfirmReverse() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
		})
	}
}

func TestClient_RejectReverse_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.RejectReverseRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " RejectReverse Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.RejectReverseRequest{
					ReverseOrderID:         "576505358041058167",
					ReverseRejectReasonKey: "reverse_reject_request_reason_1",
					ReverseRejectComments:  "Please return it ya",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.RejectReverse(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.RejectReverse() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
		})
	}
}

func TestClient_GetReverseList_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.GetReverseListRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetReverseList Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.GetReverseListRequest{
					Offset:         1,
					Size:           _PAGESIZE,
					UpdateTimeTo:   1671224452,
					UpdateTimeFrom: 1558152466,
					ReverseType:    3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetReverseList(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetReverseList() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.Empty(t, ans)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}

func TestClient_GetReverseReason_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak     string
		shopID string
		query  tiktok.GetReverseReasonRequest
	}
	tests := []struct {
		name     string
		args     args
		wantList string
		wantErr  bool
	}{
		{
			name: " GetReverseReason Integration Test ",
			args: args{
				ak:     os.Getenv(_AK_KEY),
				shopID: os.Getenv(_SHOP_KEY),
				query: tiktok.GetReverseReasonRequest{
					ReasonType:        1,
					ReverseActionType: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := c.GetReverseReason(context.TODO(), tiktok.Param{AccessToken: tt.args.ak, ShopID: tt.args.shopID}, tt.args.query)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetReverseReason() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			require.NotEmpty(t, ans)
			b, _ := json.Marshal(ans)
			jsonData := string(b)
			t.Log(jsonData)
		})
	}
}
