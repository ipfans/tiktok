package tiktok

import (
	"fmt"
	"net/http"
)

func (c *Client) GenerateAuthURL(state string) string {
	return fmt.Sprintf(
		AuthBaseURL+"/oauth/authorize?app_key=%s&state=%s",
		c.appKey, state,
	)
}

func (c *Client) GetAccessToken(code string) (resp AccessTokenResponse, err error) {
	var req GetAccessTokenRequest
	req.AppKey = c.appKey
	req.AppSecret = c.appSecret
	req.AuthCode = code
	req.GrantType = "authorized_code"
	r := c.prepareBody(req)
	err = c.request(
		http.MethodPost, AuthBaseURL,
		"/api/token/getAccessToken",
		nil, r, &resp)
	return
}

func (c *Client) RefreshToken(rk string) (resp AccessTokenResponse, err error) {
	var req RefreshTokenRequest
	req.AppKey = c.appKey
	req.AppSecret = c.appSecret
	req.RefreshToken = rk
	req.GrantType = "refresh_token"
	r := c.prepareBody(req)
	err = c.request(
		http.MethodPost, AuthBaseURL,
		"/api/token/refreshToken",
		nil, r, &resp)
	return
}
