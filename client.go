package main

import (
	"net"
	"net/http"
	"time"
)

var transport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 600 * time.Second, // OG timeout is 15
	}).Dial,
	TLSHandshakeTimeout: 600 * time.Second, // OG timeout is 15
}

var client = &http.Client{
	Transport: transport,
	Timeout:   600 * time.Second, // OG timeout is 15
}
