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
const RARITY_API = "https://api.steampowered.com/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v2/"

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
	params.Add("gameid", strconv.Itoa(appId))

	// Get the player's achievements for a specific game
	playerData := makeSteamRequest(params, ACHIEVEMENT_API)
	playerView := PlayerAchievements{}
	err := json.Unmarshal(playerData, &playerView)
	if err != nil {
		panic(err)
	}

	// Get the achievement schema for that game
	gameData := makeSteamRequest(params, SCHEMA_API)
	gameSchema := GameSchema{}
	jErr := json.Unmarshal(gameData, &gameSchema)
	if jErr != nil {
		panic(err)
	}

	// Get rarity of a game's achievements
	rarityData := makeSteamRequest(params, RARITY_API)
	rarityJson := Rarity{}
	err = json.Unmarshal(rarityData, &rarityJson)
	if err != nil {
		panic(err)
	}

	return populateUnlockedAchievements(gameSchema, playerView, rarityJson)
}

func populateUnlockedAchievements(schema GameSchema, playerView PlayerAchievements, rarity Rarity) []UnlockedAchivement {
	unlocked := []UnlockedAchivement{}
	// Map achievement id's for name lookup
	achIds := map[string]string{}
	for _, v := range schema.Game.AvailableStats.Achievements {
		achIds[v.Name] = v.DisplayName
	}

	// Populate the return object with our collected data
	for _, v := range playerView.Achievements.Unlockedchievements {
		g := UnlockedAchivement{}
		if v.Achieved > 0 {
			g.Name = achIds[v.ApiName]
			g.Rarity = getAchievementRarity(v.ApiName, rarity)
			g.Icon = getIconForAchievement(v.ApiName, schema)
			unlocked = append(unlocked, g)
		}
	}

	return unlocked
}

func getAchievementRarity(id string, rarity Rarity) float32 {
	for _, v := range rarity.Achievements.Percentages {
		if v.Name == id {
			return v.Percent
		}
	}
	return 0
}

func getIconForAchievement(id string, schema GameSchema) string {
	for _, v := range schema.Game.AvailableStats.Achievements {
		if v.Name == id {
			return v.Icon
		}
	}
	return ""
}

func makeSteamRequest(params url.Values, endpoint string) []byte {
	req, _ := url.Parse(endpoint)
	req.RawQuery = params.Encode()

	res, err := http.Get(req.String())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return data
}
