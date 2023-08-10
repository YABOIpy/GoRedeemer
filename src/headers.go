package src

import (
	http "github.com/Danny-Dasilva/fhttp"
)

var (
	in Instance
)

func (Hd *Header) MainHeader(req *http.Request, Token string, Cookie string, link string) {

	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en;q=0.9,nl;q=0.8",
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             Cookie,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/billing/promotions/" + link,
		"sec-ch-ua":          `"Chromium";v="110", "Not A(Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": in.Xprop(),
	} {
		req.Header.Set(x, o)
	}
}

func (Hd *Header) Discord(req *http.Request, headers map[string]string) {

	for x, o := range map[string]string{
		"authority":          "discord.com",
		"accept":             "*/*",
		"accept-language":    "en-US,en;q=0.9",
		"sec-ch-ua":          `"Chromium";v="110", "Not A(Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
		"sec-fetch-dest":     "document",
		"sec-fetch-mode":     "navigate",
		"sec-fetch-site":     "none",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
	} {
		req.Header.Set(x, o)
	}
	for x, o := range headers {
		req.Header.Set(x, o)
	}
}

func (Hd *Header) SecHeader(req *http.Request, Token string, Cookie string, link string) {

	for x, o := range map[string]string{
		"accept":             "*/*",
		"accept-encoding":    "gzip, deflate, br",
		"accept-language":    "en-US,en;q=0.9,nl;q=0.8",
		"authorization":      Token,
		"cookie":             Cookie,
		"origin":             "https://discord.com",
		"referer":            "https://discord.com/billing/promotions/" + link,
		"sec-ch-ua":          `"Chromium";v="110", "Not A(Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": in.Xprop(),
	} {
		req.Header.Set(x, o)
	}
}

func (Hd *Header) StripeHeader(req *http.Request, headers map[string]string) {
	for x, o := range map[string]string{
		"accept":             "application/json",
		"accept-language":    "en-US,en;q=0.9",
		"content-type":       "application/x-www-form-urlencoded",
		"dnt":                "1",
		"origin":             "https://m.stripe.network",
		"referer":            "https://m.stripe.network/",
		"sec-ch-ua":          `"Chromium";v="110", "Not A(Brand";v="24`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "cross-site",
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
	} {
		req.Header.Set(x, o)
	}
	for x, o := range headers {
		req.Header.Set(x, o)
	}
}
