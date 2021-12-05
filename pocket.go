package golang_pocket_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// NewClient creates a new client instance with your app key (to generate key visit https://getpocket.com/developer/apps/)
func NewClient(consumerKey string) (*Client, error) {
	if consumerKey == "" {
		return nil, errors.New("consumer key is empty")
	}

	return &Client{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		consumerKey: consumerKey,
	}, nil
}

// GetRequestToken obtains the request token that is used to authorize user in your application
func (c *Client) GetRequestToken(ctx context.Context, redirectUrl string) (string, error) {
	req := &requestTokenRequest{
		ConsumerKey: c.consumerKey,
		RedirectURI: redirectUrl,
	}

	values, err := c.doHTTP(ctx, endpointRequestToken, req)
	if err != nil {
		return "", err
	}

	if values.Get("code") == "" {
		return "", errors.New("empty request token in API response")
	}

	return values.Get("code"), nil
}

// GetAuthorizationURL generates link to authorize user
func (c *Client) GetAuthorizationURL(requestToken, redirectUrl string) (string, error) {
	if requestToken == "" || redirectUrl == "" {
		return "", errors.New("empty params")
	}

	return fmt.Sprintf(authorizeUrl, requestToken, redirectUrl), nil
}

// Authorize generates access token for user, that authorized in your app via link
func (c *Client) Authorize(ctx context.Context, requestToken string) (*AuthorizeResponse, error) {
	if requestToken == "" {
		return nil, errors.New("empty request token")
	}

	req := &authorizeRequest{
		Code:        requestToken,
		ConsumerKey: c.consumerKey,
	}

	values, err := c.doHTTP(ctx, endpointAuthorize, req)
	if err != nil {
		return nil, err
	}

	accessToken, username := values.Get("access_token"), values.Get("username")
	if accessToken == "" {
		return nil, errors.New("empty access token in API response")
	}

	return &AuthorizeResponse{
		AccessToken: accessToken,
		Username:    username,
	}, nil
}

// Add creates new item in Pocket list
func (c *Client) Add(ctx context.Context, input AddInput) error {
	if err := input.validate(); err != nil {
		return err
	}

	req := input.generateRequest(c.consumerKey)
	_, err := c.doHTTP(ctx, endpointAdd, req)

	return err
}

func (c *Client) doHTTP(ctx context.Context, endpoint string, body interface{}) (url.Values, error) {
	marshaledBody, err := json.Marshal(body)
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "failed to marshal input body")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, host+endpoint, bytes.NewBuffer(marshaledBody))
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "failed to create new request")
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF8")

	resp, err := c.client.Do(req)
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "failed to send http request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("API Error: %s", resp.Header.Get(xErrorHeader))
		return url.Values{}, errors.New(err)
	}

	respB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "failed to read request body")
	}

	values, err := url.ParseQuery(string(respB))
	if err != nil {
		return url.Values{}, errors.WithMessage(err, "failed to parse response body")
	}

	return values, nil
}

func (ai AddInput) generateRequest(consumerKey string) addRequest {
	return addRequest{
		URL:         ai.URL,
		Title:       ai.Title,
		Tags:        strings.Join(ai.Tags, ","),
		AccessToken: ai.AccessToken,
		ConsumerKey: consumerKey,
	}
}

func (ai AddInput) validate() error {
	if ai.URL == "" {
		return errors.New("required URL values is empty")
	}

	if ai.AccessToken == "" {
		return errors.New("access token is empty")
	}

	return nil
}
