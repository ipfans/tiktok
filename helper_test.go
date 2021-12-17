package tiktok_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ipfans/tiktok-sdk"
)

func newTestClient(t *testing.T) *tiktok.Client {
	t.Helper()
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	ak := os.Getenv("AK")
	rk := os.Getenv("RK")
	if tiktok.CheckEmpty(appKey, appSecret, ak, rk) {
		t.Skip()
	}
	logger := log.Default()
	c, _ := tiktok.New(appKey, appSecret, tiktok.WithLogger(logger))
	return c
}

func loadTestData(t *testing.T, fn string) (s string) {
	t.Helper()
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		t.Skipf("%s:loaded failed.", fn)
	}
	return string(b)
}
