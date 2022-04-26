//go:build integration
// +build integration

package tiktok_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ipfans/tiktok"
	"github.com/stretchr/testify/require"
)

func TestClient_RefreshToken_Integration(t *testing.T) {
	c := newTestClient(t)
	RK := os.Getenv("RK")
	if tiktok.CheckEmpty(RK) {
		t.Skip("AK or RK is empty")
		return
	}
	tests := []struct {
		name    string
		rk      string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "Refresh Token with valid token",
			rk:      RK,
			wantErr: require.NoError,
		},
		{
			name:    "Refresh Token with invalid token",
			rk:      RK + "111",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.RefreshToken(context.TODO(), tt.rk)
			tt.wantErr(t, err)
			if err == nil {
				require.Greater(t, got.RefreshTokenExpireIn, int(time.Now().Unix()))
			}
		})
	}
}
