# GetPocket API Golang SDK
## https://getpocket.com/developer/

#### You can access full documentation by this [link](https://github.com/siteddv/golang-pocket-sdk/blob/main/docs.md)

### Example usage:

```go
package main

import (
	"context"
	pocket "github.com/siteddv/golang-pocket-sdk"
	"log"
)

func main() {
	ctx := context.Background()

	client, err := pocket.NewClient("<your-consumer-key>") // you can generate key at https://getpocket.com/developer/apps/
	if err != nil {
		log.Fatalf("failed to init client: %s", err.Error())
		return
	}

	requestToken, err := client.GetRequestToken(ctx, "https://example.com/")
	if err != nil {
		log.Fatalf("failed to get request token: %s", err.Error())
	}

	authResp, err := client.Authorize(ctx, requestToken)
	if err != nil {
		log.Fatalf("failed to authorize: %s", err)
	}

	err = client.Add(ctx, pocket.AddInput{
		URL:         "https://github.com/zhashkevych/go-pocket-sdk",
		AccessToken: authResp.AccessToken,
	})

	if err != nil {
		log.Fatalf("failed to add item: %s", err)
	}
}
```

