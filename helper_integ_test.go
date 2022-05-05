//go:build integration
// +build integration

package tiktok_test

import (
	"log"
	"os"
	"testing"

	"github.com/ipfans/tiktok"
)

const (
	_AK_KEY             = "AK"
	_SHOP_KEY           = "SHOP"
	_ORDER_KEY          = "ORDER"
	_PAGESIZE           = 10
	_ShippingProviderID = "6841743441349706241"
	_TRACKINGNUMBER     = "JX0963674147"
	_PACKAGEID          = "1152946689385597164"
	_ORDERID            = 576500005218781420
	_ORDERIDStr         = "576500005218781420"
	_FakeOrderID        = "576474865727932168"
	_PRODUCTID          = ""
	_SKUID              = ""
	_WAREHOUSEID        = ""
	_CATEGORYID         = "600100"
	_DESC               = 1 // Available values: 1 (DESC), 2 (ASC) Default value 1
	_ASC                = 2
)

func newTestClient(t *testing.T) *tiktok.Client {
	t.Helper()
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	ak := os.Getenv("AK")
	if tiktok.CheckEmpty(appKey, appSecret, ak) {
		t.Skip("APPKEY, APPSECRET or AK is empty")
	}
	logger := log.Default()
	c, _ := tiktok.New(appKey, appSecret, tiktok.WithLogger(logger))
	return c
}
