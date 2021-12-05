# SDK documentation
#### You can access SDK by this [link](https://github.com/siteddv/golang-pocket-sdk)

## Public structures documentation
```go
AddInput struct {
    URL         string
    Title       string
    Tags        []string
    AccessToken string
}
```

**AddInput** contains note data to send it to "Pocket". 
URL and AccessToken is required. Other fields are optional.
* **

```go
Client struct {
    client      *http.Client
    consumerKey string
}
```

**Client** contains a new client instance with your app key (to generate key visit https://getpocket.com/developer/apps/)
* **

```go
AuthorizeResponse struct {
    AccessToken string `json:"access_token"`
    Username    string `json:"username"`
}
```

**AuthorizeResponse** contains response data after user authorization

## Public functions documentation

```go
func NewClient(consumerKey string) (*Client, error)
```

**NewClient** creates a new client instance with your app key (to generate key visit https://getpocket.com/developer/apps/)
* **

```go
func (c *Client) GetRequestToken(ctx context.Context, redirectUrl string) (string, error)
```

**GetRequestToken** obtains the request token that is used to authorize user in your application
* **

```go
func (c *Client) GetAuthorizationURL(requestToken, redirectUrl string) (string, error)
```

**GetAuthorizationURL** generates link to authorize user
* **

```go
func (c *Client) Authorize(ctx context.Context, requestToken string) (*AuthorizeResponse, error)
```

**Authorize** generates link to authorize user
* **

```go
func (c *Client) Add(ctx context.Context, input AddInput) error
```

**Add** creates new item in Pocket list
