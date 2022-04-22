package tiktok_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ipfans/tiktok-sdk"
	"github.com/stretchr/testify/require"
)

func TestClient_GenerateAuthURL(t *testing.T) {
	client, err := tiktok.New("123abc", "456def")
	require.NoError(t, err)
	link := client.GenerateAuthURL("123")
	require.Equal(t, "https://auth.tiktok-shops.com/oauth/authorize?app_key=123abc&state=123", link)
}

func TestClient_GetAccessToken(t *testing.T) {
	c, close := mockClient(t)
	defer close()
	b, _ := os.ReadFile("testdata/get_access_token.json")

	res := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBuffer(b)),
		ContentLength: int64(len(b)),
		Request:       nil,
		Header:        make(http.Header, 0),
	}

	c.EXPECT().Do(gomock.Any()).Return(res, nil)
	client, err := tiktok.New("123abc", "456def", tiktok.WithHTTPClient(c))
	require.NoError(t, err)
	resp, err := client.GetAccessToken(context.TODO(), "123")
	require.NoError(t, err)
	require.Equal(t, tiktok.AccessTokenResponse{
		AccessToken:          "RLM6CIADWF606TZGFO5XGA",
		AccessTokenExpireIn:  1630401330,
		RefreshToken:         "C2XWDN63ON-FOHJSMR0WSG",
		RefreshTokenExpireIn: 1630401510,
		OpenID:               "7000636243100829446",
		SellerName:           "fanfan",
	}, resp)
}

func TestClient_RefreshToken(t *testing.T) {
}
