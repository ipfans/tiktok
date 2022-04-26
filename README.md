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
  - [x] GetOrderList
  - [x] GetOrderDetail
  - [x] ShipOrder
- [x] Fulfillment API
  - [x] SearchPreCombinePkg
  - [x] ConfirmPreCombinePkg
  - [x] RemovePackageOrder
  - [x] GetPackagePickupConfig
  - [x] ShipPackage
  - [x] SearchPackage
  - [x] GetPackageDetail
  - [x] GetPackageShippingInfo
  - [x] UpdatePackageShippingInfo
  - [x] GetPackageShippingDocument
  - [x] VerifyOrderSplit
  - [x] ConfirmOrderSplit
- [x] Logistics API
  - [x] GetShippingInfo
  - [x] UpdateShippingInfo
  - [x] GetShippingDocument
  - [x] GetWarehouseList
  - [x] GetShippingProvider
- [x] Product API
  - [x] GetCategory
  - [x] GetAttribute
  - [x] GetCategoryRule
  - [x] GetBrand
  - [x] UploadImg
  - [x] UploadFile
  - [x] CreateProduct
  - [x] EditProduct
  - [x] GetProductList
  - [x] GetProductDetail
  - [x] UpdatePrice
  - [x] UpdateStock
  - [x] DeactivateProducts
  - [x] DeleteProducts
  - [x] RecoverProduct
  - [x] ActivateProduct
- [x] Shop API
  - [x] GetAuthorizedShop
- [x] Finance API
  - [x] SearchSettlements
  - [x] GetOrderSettlements
  - [x] SearchTransactions
- [x] Reverse Order API
  - [x] ConfirmReverse
  - [x] RejectReverse
  - [x] GetReverseList
  - [x] GetReverseReason

### TODO

- [ ] All Feature supports
- [ ] Integraion tests

## API Call Flow

https://bytedance.feishu.cn/docs/doccnZ15f4vPK4qOOBTKhgc2xNh

## HOWTO Contribuate

Install [`taskfile`](https://taskfile.dev) before you start.

### Setup Intergration Environment Variables

edit `.env` file in root directory:

```env
APPKEY=<app id for tiktok>
APPSECRET=<app secret for tiktok>
AK=<access token>
RK=<refresh token>
OPENID=<seller's openid>
SELLER=<seller's name>
```
### Run Tests

```bash
task  # run tests.
task integration # run integration tests.
```

If env is not setup, all integration tests will be skipped.
