# `dnsup` - a simple DNS updater client for Cloudflare

A quick hack to create or update a given DNS record via Cloudflare's API.

I use to this to update a DNS record for my personal zone as and when my home IP address changes.

## Usage

First, make sure you have the `CF_API_KEY` and `CF_API_EMAIL` environment variables set.  Then:

```
./dnsup <fqdn>
```

## Thanks

@mhayden for providing https://icanhazip.com

