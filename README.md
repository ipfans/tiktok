# tiktok

[![Go Reference](https://pkg.go.dev/badge/github.com/ipfans/tiktok-sdk.svg)](https://pkg.go.dev/github.com/ipfans/tiktok-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/ipfans/tiktok-sdk)](https://goreportcard.com/report/github.com/ipfans/tiktok-sdk)
[![build](https://github.com/ipfans/tiktok/actions/workflows/ci.yml/badge.svg)](https://github.com/ipfans/tiktok/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/ipfans/tiktok/badge.svg?branch=master)](https://coveralls.io/github/ipfans/tiktok?branch=master)

Go SDK for Tiktok Shop Open Platform.

## Feature

- [x] Authentication
  - [x] Get your auth code url
  - [x] Use auth code to get access_token
  - [x] Refresh your access token
- [ ] Webhook
- [x] Order API
  - [x] Get order list
  - [x] Get order detail
  - [x] Ship order
- [ ] Fulfillment API
- [ ] Logistics API
- [ ] Product API
- [x] Shop API
  - [x] Get authorized shop list
- [x] Finance API
  - [x] Search settlements within a certain timeframe
  - [x] Query settlement of a specific order
  - [x] Search your seller account's transactions within a certain timeframe
- [ ] Reverse Order API

### TODO

- [ ] All Feature supports
- [ ] Integraion tests

## API Call Flow

https://bytedance.feishu.cn/docs/doccnZ15f4vPK4qOOBTKhgc2xNh

## HOWTO Contribuate

Install [`taskfile`](https://taskfile.dev) before you start.

1. Setup Intergration Environment Variables

edit `.env` file in root directory:

```
APPKEY=<app id for tiktok>
APPSECRET=<app secret for tiktok>
AK=<access token>
RK=<refresh token>
OPENID=<seller's openid>
SELLER=<seller's name>
```

```
task tests  # run tests.
```

If env is not setup, all tests will be skipped.
