package tests

import "time"

type Response struct {
	URL             string        `json:"url"`
	CustomShortUrl  string        `json:"short_url"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRatelimitReset time.Duration `json:"rate_limit_reset"`
}

type Error struct {
	Error string `json:"error"`
}
