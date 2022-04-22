package tiktok_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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

	var want struct {
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

			err = json.Unmarshal([]byte(tt.Want), &want)
			require.NoError(t, err)

			httpmock.RegisterResponder(
				tt.Request.Method, tt.Request.URL,
				func(r *http.Request) (res *http.Response, err error) {
					defer r.Body.Close()
					b, _ := io.ReadAll(r.Body)
					require.JSONEq(t, string(tt.Request.Body), string(b))
					res, err = httpmock.NewJsonResponse(200, tt.Response.Body)
					return
				},
			)

			var req request
			err := json.Unmarshal(tt.Args, &req)
			require.NoError(t, err)

			resp, err := client.GetAccessToken(context.TODO(), req.Code)
			if tt.WantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, want.AccessToken, resp.AccessToken)
			require.Equal(t, want.RefreshToken, resp.RefreshToken)
			require.Equal(t, want.OpenID, resp.OpenID)
		})
	}

}

func TestClient_RefreshToken(t *testing.T) {
	type request struct {
		RefreshToken string `json:"refresh_token"`
	}

	var want struct {
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

			err = json.Unmarshal([]byte(tt.Want), &want)
			require.NoError(t, err)

			httpmock.RegisterResponder(
				tt.Request.Method, tt.Request.URL,
				func(r *http.Request) (res *http.Response, err error) {
					defer r.Body.Close()
					b, _ := io.ReadAll(r.Body)
					require.JSONEq(t, string(tt.Request.Body), string(b))
					res, err = httpmock.NewJsonResponse(200, tt.Response.Body)
					return
				},
			)

			var req request
			err := json.Unmarshal(tt.Args, &req)
			require.NoError(t, err)

			resp, err := client.RefreshToken(context.TODO(), req.RefreshToken)
			if tt.WantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, want.AccessToken, resp.AccessToken)
			require.Equal(t, want.RefreshToken, resp.RefreshToken)
			require.Equal(t, want.OpenID, resp.OpenID)
		})
	}
}
