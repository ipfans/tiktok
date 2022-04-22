package tiktok_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ipfans/tiktok"
	"github.com/stretchr/testify/require"
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

type TestRecord struct {
	Name    string          `json:"name"`
	Args    json.RawMessage `json:"args"`
	Request struct {
		Method string          `json:"method"`
		URL    string          `json:"url"`
		Body   json.RawMessage `json:"body"`
	} `json:"request"`
	Response struct {
		Status  int               `json:"status"`
		Headers map[string]string `json:"headers"`
		Body    json.RawMessage   `json:"body"`
	} `json:"response"`
	Want    json.RawMessage `json:"want"`
	WantErr bool            `json:"want_err"`
}

func loadTestData(t *testing.T, fn string) (records []TestRecord) {
	t.Helper()
	b, err := ioutil.ReadFile(fn)
	require.NoError(t, err)
	err = json.Unmarshal(b, &records)
	require.NoError(t, err)
	return
}
