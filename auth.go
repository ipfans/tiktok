package tiktok

import (
	"context"
	"fmt"
	"net/http"
)

// GenerateAuthURL generate auth url for user to login.
// Doc: https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg#ckvNFO
func (c *Client) GenerateAuthURL(state string) string {
	return fmt.Sprintf(
		AuthBaseURL+"/oauth/authorize?app_key=%s&state=%s",
		c.appKey, state,
	)
}

// GetAccessToken get access token from tiktok server.
// Doc: https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg#qYtWHF
func (c *Client) GetAccessToken(ctx context.Context, code string) (resp AccessTokenResponse, err error) {
	var req GetAccessTokenRequest
	req.AppKey = c.appKey
	req.AppSecret = c.appSecret
	req.AuthCode = code
	req.GrantType = "authorized_code"
	r := c.prepareBody(req)
	err = c.request(
		ctx, http.MethodPost, AuthBaseURL,
		"/api/token/getAccessToken",
		nil, r, &resp)
	return
}

// RefreshToken refresh access token.
// Doc: https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg#bG2h09
func (c *Client) RefreshToken(ctx context.Context, rk string) (resp AccessTokenResponse, err error) {
	var req RefreshTokenRequest
	req.AppKey = c.appKey
	req.AppSecret = c.appSecret
	req.RefreshToken = rk
	req.GrantType = "refresh_token"
	r := c.prepareBody(req)
	err = c.request(
		ctx, http.MethodPost, AuthBaseURL,
		"/api/token/refreshToken",
		nil, r, &resp)
	return
}
