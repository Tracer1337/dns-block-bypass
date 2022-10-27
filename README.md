# DNS Block Bypass

This project lets you capture a screenshot of any website, even when it's dns-blocked.

### Setup

1. [Register a free account at Browserless](https://cloud.browserless.io/account)
2. Grab the API Key shown on the account page
3. Download the binary for your OS from the [Releases Page](https://github.com/Tracer1337/dns-block-bypass/releases) 
4. Run the binary with the following parameters: `-token=YOUR_TOKEN_HERE -url=URL_HERE`

### Compiling

Clone the Repo and run `go install` ([Go](https://go.dev/) is required for this step)

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
