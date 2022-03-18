package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const ACHIEVEMENT_API = "https://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1/"
const SCHEMA_API = "https://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2/"
const RECENTLY_PLAYED = "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v1/"

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")

	serviceParam := url.Values{}
	serviceParam.Add("key", os.Getenv("STEAM_KEY"))
	serviceParam.Add("steamid", "76561198086180357")
	serviceParam.Add("count", "3")

	recentlyPlayed, _ := url.Parse(RECENTLY_PLAYED)
	recentlyPlayed.RawQuery = serviceParam.Encode()

	res, err := http.Get(recentlyPlayed.String())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	threeRecent, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	recentGames := RecentlyPlayed{}
	err = json.Unmarshal(threeRecent, &recentGames)
	if err != nil {
		panic(err)
	}

	dc := DisplayComponent{}

	for _, v := range recentGames.Response.Games {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("appid: %d\n", v.AppId)
		aC := getGameAchievements(v.AppId)
		fmt.Printf("Unlocked %d\n", len(aC))
		game := Game{
			Title:               v.Name,
			UnlockedAchivements: aC,
		}
		dc.Games = append(dc.Games, game)
	}

	gameJson, err := json.Marshal(&dc)
	if err != nil {
		panic(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            string(gameJson),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func getGameAchievements(appId int) []UnlockedAchivement {
	params := url.Values{}
	params.Add("key", os.Getenv("STEAM_KEY"))
	params.Add("steamid", "76561198086180357")
	params.Add("appid", strconv.Itoa(appId))

	playerAch, _ := url.Parse(ACHIEVEMENT_API)
	playerAch.RawQuery = params.Encode()

	res, err := http.Get(playerAch.String())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	playerData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	playerView := PlayerAchievements{}
	err = json.Unmarshal(playerData, &playerView)
	if err != nil {
		panic(err)
	}

	gameScheme, _ := url.Parse(SCHEMA_API)
	gameScheme.RawQuery = params.Encode()

	res, err = http.Get(gameScheme.String())
	if err != nil {
		panic(err)
	}

	gameData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	gameSchema := GameSchema{}
	jErr := json.Unmarshal(gameData, &gameSchema)
	if jErr != nil {
		panic(err)
	}
	return findUnlockedAchievements(gameSchema, playerView)
}

func findUnlockedAchievements(schema GameSchema, playerView PlayerAchievements) []UnlockedAchivement {
	unlocked := []UnlockedAchivement{}
	gameIds := map[string]string{}
	for _, v := range schema.Game.AvailableStats.Achievements {
		gameIds[v.Name] = v.DisplayName
	}

	for _, v := range playerView.Achievements.Unlockedchievements {
		g := UnlockedAchivement{}
		if v.Achieved > 0 {
			g.Name = gameIds[v.ApiName]
			unlocked = append(unlocked, g)
		}
	}

	return unlocked
}
