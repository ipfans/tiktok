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
	_ShippingProviderID = ""
	_TRACKINGNUMBER     = ""
	_PACKAGEID          = ""
	_ORDERID            = 0
	_ORDERIDStr         = ""
	_FakeOrderID        = ""
	_PRODUCTID          = ""
	_SKUID              = ""
	_WAREHOUSEID        = ""
	_CATEGORYID         = ""
	_SORT_TYPE          = 1 // Available values: 1 (DESC), 2 (ASC) Default value 1
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
