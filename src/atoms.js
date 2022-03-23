import { atom, selector } from "recoil";

const achSelector = selector({
  key: "achSelector",
  get: async ({ get }) => {
    return await getAchievementsForGame();
  },
});

export const achievements = atom({
  key: "achievements",
  default: achSelector,
});

async function getAchievementsForGame() {
  const achievementEndpoint = ".netlify/functions/unlocked-achievements";
  const res = await fetch(`${achievementEndpoint}`);
  return res.json();
}

export async function steamAuth() {
  const authEndpoint = ".netlify/functions/auth-steam";
  const res = await fetch(`${authEndpoint}`);
  console.log(res.json());
  return res.json();
}
