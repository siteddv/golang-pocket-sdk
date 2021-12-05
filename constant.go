package golang_pocket_sdk

import "time"

const (
	host         = "https://getpocket.com/v3"
	authorizeUrl = "https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s"

	endpointAdd          = "/add"
	endpointRequestToken = "/oauth/request"
	endpointAuthorize    = "/oauth/authorize"

	// xErrorHeader used to parse error message from Headers on non-2XX responses
	xErrorHeader = "X-Error"

	defaultTimeout = 5 * time.Second
)
