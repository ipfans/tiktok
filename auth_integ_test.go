//go:build integration
// +build integration

package tiktok_test

// import (
// 	"context"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// )

// func TestClient_RefreshToken_Integration(t *testing.T) {
// 	c := newTestClient(t)
// 	tests := []struct {
// 		name    string
// 		rk      string
// 		wantErr bool
// 	}{
// 		{
// 			name: "Refresh Token with valid token",
// 			rk:   os.Getenv("RK"),
// 		},
// 		{
// 			name:    "Refresh Token with invalid token",
// 			rk:      os.Getenv("RK") + "111",
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotResp, err := c.RefreshToken(context.TODO(), tt.rk)
// 			require.Equal(t, tt.wantErr, err != nil, "Client.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
// 			if err != nil {
// 				return
// 			}
// 			require.Equal(t, tt.rk, gotResp.RefreshToken)
// 			require.Greater(t, gotResp.RefreshTokenExpireIn, int(time.Now().Unix()))
// 		})
// 	}
// }
