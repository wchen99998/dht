package dht

import (
	"golang.org/x/time/rate"
)

// https://github.com/wchen99998/torrent/issues/1005#issuecomment-2856881633
var DefaultSendLimiter = rate.NewLimiter(250, 25)
