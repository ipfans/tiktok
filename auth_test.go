package tiktok_test

import (
	"context"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestClient_GenerateAuthURL(t *testing.T) {
	client, err := tiktok.New("123abc", "456def")
	require.NoError(t, err)
	link := client.GenerateAuthURL("123")
	require.Equal(t, "https://auth.tiktok-shops.com/oauth/authorize?app_key=123abc&state=123", link)
}

func TestClient_GetAccessToken(t *testing.T) {
	type request struct {
		Code string `json:"code"`
	}

	type want struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		OpenID       string `json:"open_id"`
	}

	client, err := tiktok.New("123abc", "456def")
	require.NoError(t, err)
	tests := loadTestData(t, "testdata/get_access_token.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			var w want
			var req request
			setupMock(t, tt, &req, &w)

			resp, err := client.GetAccessToken(context.TODO(), req.Code)
			if tt.WantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, w.AccessToken, resp.AccessToken)
			require.Equal(t, w.RefreshToken, resp.RefreshToken)
			require.Equal(t, w.OpenID, resp.OpenID)
		})
	}

}

func TestClient_RefreshToken(t *testing.T) {
	type request struct {
		RefreshToken string `json:"refresh_token"`
	}

	type want struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		OpenID       string `json:"open_id"`
	}

	client, err := tiktok.New("123abc", "456def")
	require.NoError(t, err)
	tests := loadTestData(t, "testdata/refresh_token.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			var req request
			var w want
			setupMock(t, tt, &req, &w)

			resp, err := client.RefreshToken(context.TODO(), req.RefreshToken)
			if tt.WantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, w.AccessToken, resp.AccessToken)
			require.Equal(t, w.RefreshToken, resp.RefreshToken)
			require.Equal(t, w.OpenID, resp.OpenID)
		})
	}
}
