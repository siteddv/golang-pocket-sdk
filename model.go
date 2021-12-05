package golang_pocket_sdk

import "net/http"

type (
	requestTokenRequest struct {
		ConsumerKey string `json:"consumer_key"`
		RedirectURI string `json:"redirect_uri"`
	}

	authorizeRequest struct {
		ConsumerKey string `json:"consumer_key"`
		Code        string `json:"code"`
	}

	AuthorizeResponse struct {
		AccessToken string `json:"access_token"`
		Username    string `json:"username"`
	}

	addRequest struct {
		URL         string `json:"url"`
		Title       string `json:"title,omitempty"`
		Tags        string `json:"tags,omitempty"`
		AccessToken string `json:"access_token"`
		ConsumerKey string `json:"consumer_key"`
	}

	AddInput struct {
		URL         string
		Title       string
		Tags        []string
		AccessToken string
	}

	Client struct {
		client      *http.Client
		consumerKey string
	}
)
