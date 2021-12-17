package tiktok_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_GetAuthorizedShop(t *testing.T) {
	c := newTestClient(t)
	type args struct {
		ak      string
		shop_id string
	}
	b := loadTestData(t, "testdata/shop.json")
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
			wantShops: b,
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
			gotShops, err := c.GetAuthorizedShop(tt.args.ak, tt.args.shop_id)
			require.Equal(t, tt.wantErr, err != nil, "Client.GetAuthorizedShop() error = %v, wantErr %v", err, tt.wantErr)
			if err != nil {
				return
			}
			b, _ := json.Marshal(gotShops)
			jsonData := string(b)
			require.JSONEq(t, tt.wantShops, jsonData)
		})
	}
}
