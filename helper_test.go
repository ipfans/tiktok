package tiktok_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ipfans/tiktok"
	"github.com/jarcoal/httpmock"
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
		Query  string          `json:"query"`
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

func setupMock(t *testing.T, tt TestRecord, args, want interface{}) {
	httpmock.RegisterResponder(
		tt.Request.Method, tt.Request.URL,
		func(r *http.Request) (res *http.Response, err error) {
			defer r.Body.Close()
			if tt.Request.Query != "" {
				require.Equal(t, tt.Request.Query, r.URL.RawQuery)
			}
			b, _ := io.ReadAll(r.Body)
			require.JSONEq(t, string(tt.Request.Body), string(b))
			res, err = httpmock.NewJsonResponse(200, tt.Response.Body)
			return
		},
	)

	var err error
	if want != nil {
		err = json.Unmarshal([]byte(tt.Want), want)
		require.NoError(t, err)
	}
	if args != nil {
		err = json.Unmarshal(tt.Args, args)
		require.NoError(t, err)
	}
}

func mockTime() func() {
	tiktok.Timestamp = func() string {
		return "1600000000"
	}

	return func() {
		tiktok.Timestamp = func() string {
			return fmt.Sprintf("%d", time.Now().Unix())
		}
	}
}
