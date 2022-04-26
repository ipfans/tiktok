//go:build integration
// +build integration

package tiktok_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetAuthorizedShop_Integration(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak      string
		shop_id string
	}
	tests := []struct {
		name      string
		args      args
		wantShops string
		wantErr   bool
	}{
		{
			name: "Request with ak",
			args: args{
				ak: os.Getenv("AK"),
			},
			wantShops: os.Getenv("SHOP"),
			wantErr:   false,
		},
		{
			name: "Request with invalid ak",
			args: args{
				ak: os.Getenv("AK") + "111",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShops, err := c.GetAuthorizedShop(context.TODO(), tt.args.ak, tt.args.shop_id)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetAuthorizedShop() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			contains := false
			for i := range gotShops.Shops {
				if gotShops.Shops[i].ShopID == tt.wantShops {
					contains = true
				}
			}
			require.True(t, contains)
		})
	}
}
