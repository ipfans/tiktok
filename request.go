package tiktok

type GetAccessTokenRequest struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	AuthCode  string `json:"auth_code"`
	GrantType string `json:"grant_type"`
}

type RefreshTokenRequest struct {
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}
