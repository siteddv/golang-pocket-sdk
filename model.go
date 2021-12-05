package pocket

import "net/http"

type (
	// AddInput contains note data to send it to "Pocket".
	//URL and AccessToken is required. Other fields are optional.
	AddInput struct {
		URL         string
		Title       string
		Tags        []string
		AccessToken string
	}

	//Client contains a new client instance with your app key
	Client struct {
		client      *http.Client
		consumerKey string
	}

	// AuthorizeResponse contains response data after user authorization
	AuthorizeResponse struct {
		AccessToken string `json:"access_token"`
		Username    string `json:"username"`
	}

	requestTokenRequest struct {
		ConsumerKey string `json:"consumer_key"`
		RedirectURI string `json:"redirect_uri"`
	}

	authorizeRequest struct {
		ConsumerKey string `json:"consumer_key"`
		Code        string `json:"code"`
	}

	addRequest struct {
		URL         string `json:"url"`
		Title       string `json:"title,omitempty"`
		Tags        string `json:"tags,omitempty"`
		AccessToken string `json:"access_token"`
		ConsumerKey string `json:"consumer_key"`
	}
)
