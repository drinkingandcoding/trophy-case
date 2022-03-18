import { atom, selector } from 'recoil';

const achSelector = selector({
    key: "achSelector",
    get: async ({get}) => {
        return await getAchievementsForGame()
    }
});

export const achievements = atom({
    key: "achievements",
    default: achSelector
});

async function getAchievementsForGame() {
    const achievementEndpoint = ".netlify/functions/unlocked-achievements";
    const res = await fetch(
      // Make this not this
      `${achievementEndpoint}`
    );
    return res.json();
}