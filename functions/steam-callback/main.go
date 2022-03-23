package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yohcop/openid-go"
)

const (
	LOGIN_ENDPOINT   = "https://steamcommunity/openid/login"
	OPEN_ID_MODE     = "checkid_setup"
	OPENIDNS         = "http://specs.openid.net/auth/2.0"
	OPENIDIDENTIFIER = "http://specs.openid.net/auth/2.0/identifier_select"
)

// NoOpDiscoveryCache implements the DiscoveryCache interface and doesn't cache anything.
// For a simple website, I'm not sure you need a cache.
type NoOpDiscoveryCache struct{}

// Put is a no op.
func (n *NoOpDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {}

// Get always returns nil.
func (n *NoOpDiscoveryCache) Get(id string) openid.DiscoveredInfo {
	return nil
}

var nonceStore = openid.NewSimpleNonceStore()
var discoveryCache = &NoOpDiscoveryCache{}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")
	fullURL := "http://localhost:8888/.netlify/functions/steam-callback"

	id, err := openid.Verify(fullURL, discoveryCache, nonceStore)
	data := make(map[string]string)
	fmt.Printf("%s", data["user"])
	if err != nil {
		log.Printf("Error verifying: %q\n", err)
	} else {
		log.Printf("NonceStore: %+v\n", nonceStore)
		data["user"] = id
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "text/plain",
			"Access-Control-Allow-Origin": "*",
		},
		Body:            string(data["user"]),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
