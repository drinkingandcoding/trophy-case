package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// TODO: Add recently played
const ACHIEVEMENT_API = "https://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1/"
const SCHEMA_API = "https://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2/"

type GameSchema struct {
	Game GameObject `json:"game"`
}

type GameObject struct {
	GameName       string             `json:"gameName"`
	GameVersion    string             `json:"gameVersion,omitempty"`
	AvailableStats AvailableGameStats `json:"availableGameStats"`
}

type AvailableGameStats struct {
	Achievements []Achievement
}

type Achievement struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue,omitempty"`
	DisplayName  string `json:"displayName"`
	Hidden       int    `json:"hidden"`
	Icon         string `json:"icon"`
	IconGray     string `json:"icongray,omitempty"`
}

type PlayerAchievements struct {
	Achievements PlayerStats `json:"playerstats"`
}

type PlayerStats struct {
	Unlockedchievements []GameAchievement `json:"achievements"`
	GameName            string            `json:"gameName"`
	SteamID             string            `json:"steamID"`
	Success             bool              `json:"success,omitempty"`
}

type GameAchievement struct {
	Achieved   int    `json:"achieved"`
	ApiName    string `json:"apiname"`
	UnlockTime int    `json:"unlocktime"`
}

type DisplayComponent struct {
	Games []Game `json:"games"`
}

type Game struct {
	Title               string               `json:"title"`
	UnlockedAchivements []UnlockedAchivement `json:"unlockedAchievement"`
}

type UnlockedAchivement struct {
	Name   string  `json:"name"`
	Rarity float32 `json:"rarity,omitempty"`
	Icon   string  `json:"icon"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")

	params := url.Values{}
	params.Add("key", os.Getenv("STEAM_KEY"))
	params.Add("steamid", "76561198086180357")
	params.Add("appid", "1245620")

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

	fmt.Printf("%+v", playerView)

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

	// 3 most recently played
	dc := DisplayComponent{}
	eldenRing := Game{
		Title:               "Elden Ring",
		UnlockedAchivements: findUnlockedAchievements(gameSchema, playerView),
	}
	dc.Games = append(dc.Games, eldenRing)

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
