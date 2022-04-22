# tiktok
 Go SDK for TikTok global


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
