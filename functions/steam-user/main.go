package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const ACHIEVEMENT_API = "https://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1/"

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")

	params := url.Values{}
	params.Add("key", os.Getenv("STEAM_KEY"))
	params.Add("steamid", "76561198086180357")
	params.Add("appid", "1245620")

	url, _ := url.Parse(ACHIEVEMENT_API)
	url.RawQuery = params.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bodyB, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            string(bodyB),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
