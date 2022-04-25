package tiktok_test

import (
	"context"
	"strings"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCategory(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_category.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.CategoryList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetCategory(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_GetAttribute(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		CategoryID  string `json:"category_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_attribute.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.AttributeList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetAttribute(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.CategoryID,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_GetCategoryRule(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		CategoryID  string `json:"category_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_category_rule.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.CategoryRules
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetCategoryRule(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.CategoryID,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_GetBrand(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/get_brand.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.BrandList
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			list, err := c.GetBrand(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, list)
		})
	}
}

func TestClient_UploadImg(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		File        string `json:"file"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/upload_img.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.ImageInfo
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)
			r := strings.NewReader(args.File)

			img, err := c.UploadImgReader(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				tiktok.ImgSceneAttributeImage, r,
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, img)
		})
	}
}

func TestClient_UploadFile(t *testing.T) {
	var args struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		AccessToken string `json:"access_token"`
		ShopID      string `json:"shop_id"`
		File        string `json:"file"`
		Name        string `json:"name"`
	}

	restore := mockTime()
	defer restore()

	tests := loadTestData(t, "testdata/product/upload_file.json")
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			var want tiktok.FileInfo
			setupMock(t, tt, &args, &want)

			c, err := tiktok.New(args.AppKey, args.AppSecret)
			require.NoError(t, err)

			info, err := c.UploadFile(context.TODO(),
				tiktok.Param{args.AccessToken, args.ShopID},
				args.Name, []byte(args.File),
			)
			if tt.WantErr {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, want, info)
		})
	}
}
