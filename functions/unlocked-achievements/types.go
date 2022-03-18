package main

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
	Achieved    int    `json:"achieved"`
	ApiName     string `json:"apiname"`
	UnlockTime  int    `json:"unlocktime"`
	Description string `json:"description,omitempty"`
}

type DisplayComponent struct {
	Games []Game `json:"games"`
}

type Game struct {
	Title               string               `json:"title"`
	UnlockedAchivements []UnlockedAchivement `json:"unlockedAchievements"`
}

type UnlockedAchivement struct {
	Name        string  `json:"name"`
	Rarity      float32 `json:"rarity,omitempty"`
	Icon        string  `json:"icon"`
	Description string  `json:"description,omitempty"`
}

type RecentlyPlayed struct {
	Response RecentlyPlayedResponse `json:"response"`
}

type RecentlyPlayedResponse struct {
	TotalCount int           `json:"total_count,omitempty"`
	Games      []RecentGames `json:"games"`
}

type RecentGames struct {
	AppId           int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeRecent  int    `json:"playtime_2weeks,omitempty"`
	PlaytimeLinux   int    `json:"playtime_linux_forever,omitempty"`
	PlaytimeMac     int    `json:"playtime_mac_forever,omitempty"`
	PlaytimeWindows int    `json:"playtime_windows_forever,omitempty"`
	PlaytimeTotal   int    `json:"playtime_forever,omitempty"`
	ImageIcon       string `json:"img_icon_url,omitempty"`
}

type Rarity struct {
	Achievements AchievementPercents `json:"achievementpercentages"`
}

type AchievementPercents struct {
	Percentages []AchievementPercent `json:"achievements"`
}

type AchievementPercent struct {
	Name    string  `json:"name"`
	Percent float32 `json:"percent"`
}
