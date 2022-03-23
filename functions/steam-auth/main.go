package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yohcop/openid-go"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")

	url, err := openid.RedirectURL(
		"http://steamcommunity.com/openid",
		"http://localhost:8888/.netlify/functions/steam-callback",
		"http:/localhost:8888/",
	)
	fmt.Printf("%s", url)

	if err != nil {
		log.Printf("Error creating redirect URL: %q\n", err)
	}
	fmt.Println("Redirecting")
	return &events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"location":                    url,
			"cache-control":               "private",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
