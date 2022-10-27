# DNS Block Bypass

This project lets you capture a screenshot of any website, even when it's dns-blocked.

### Setup

1. [Register a free account at Browserless](https://cloud.browserless.io/account)
2. Grab the API Key shown on the account page
3. Clone the Repo and run `go install` ([Go](https://go.dev/) is required for this step)
4. Run `dns-block-bypass -token=YOUR_TOKEN_HERE -url=URL_HERE`

### Usage

```
Usage of dns-block-bypass:
  -out string
        The output file (default "screenshot.png")
  -token string
        Your Browserless API Token
  -url string
        The URL you want to screenshot
```
