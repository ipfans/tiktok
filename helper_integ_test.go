package tiktok_test

// import (
// 	"log"
// 	"os"
// 	"testing"

// 	"github.com/ipfans/tiktok"
// )

// func newTestClient(t *testing.T) *tiktok.Client {
// 	t.Helper()
// 	appKey := os.Getenv("APPKEY")
// 	appSecret := os.Getenv("APPSECRET")
// 	ak := os.Getenv("AK")
// 	rk := os.Getenv("RK")
// 	if tiktok.CheckEmpty(appKey, appSecret, ak, rk) {
// 		t.Skip()
// 	}
// 	logger := log.Default()
// 	c, _ := tiktok.New(appKey, appSecret, tiktok.WithLogger(logger))
// 	return c
// }
