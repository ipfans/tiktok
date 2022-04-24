package tiktok

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
)

func generateSHA256(path string, queries url.Values, secret string) string {
	keys := make([]string, len(queries))
	idx := 0
	for k := range queries {
		keys[idx] = k
		idx++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	input := path
	for _, key := range keys {
		input = input + key + queries.Get(key)
	}
	input = secret + input + secret

	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(input)); err != nil {
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

func safeGet(param url.Values, key string) string {
	if param == nil {
		return ""
	}
	vs, ok := param[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func CheckEmpty(vals ...string) bool {
	for _, v := range vals {
		if v == "" {
			return true
		}
	}
	return false
}

type Param struct {
	AccessToken string `validate:"required"`
	ShopID      string `validate:"required"`
}

func (c *Client) params(p Param) (param url.Values, err error) {
	err = c.validate.Struct(p)
	if err != nil {
		return
	}
	param = url.Values{}
	param.Set("access_token", p.AccessToken)
	param.Set("shop_id", p.ShopID)
	return
}
