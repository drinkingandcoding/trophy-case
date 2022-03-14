import { Handler } from '@netlify/functions'
import fetch from 'node-fetch'

const ACHIEVEMENTS_API = "https://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1"
const SCHEMA_API = "https://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2/"

export const handler: Handler = async (event, context) => {
  let ach_params = {
    "key": process.env.STEAM_KEY,
    "steamid": "76561198086180357",
    "appid": "1245620"
  };

  let sch_params = {
    "key": process.env.STEAM_KEY,
    "appid": "1245620"
  };

  let ach_url = ACHIEVEMENTS_API
  let ach_search_params = new URLSearchParams(ach_params).toString()

  let sch_url = SCHEMA_API
  let sch_search_params = new URLSearchParams(sch_params).toString()

  const response = await fetch(sch_url + "?" + sch_search_params)
  const data = await response.json()

  return {
    statusCode: 200,
    body: JSON.stringify(data),
  }
}
