import React from "react";
import "./App.css";
import { useQuery } from "react-query";

function Achievements() {
  const { data, status } = useQuery("achievements", getAchievementsForGame);

  if (status === "loading") {
    return <p>Loading...</p>;
  }
  if (status === "error") {
    return <p>Error!</p>;
  }

  console.log(data);

  return (
    <div>
      <h3>{data.games[0].title}</h3>
      <ul>
        {data.games[0].unlockedAchievement.map((ach) => (
          <li key={ach.name}>{ach.name}</li>
        ))}
      </ul>
    </div>
  );
}

async function getAchievementsForGame() {
  const achievementEndpoint = ".netlify/functions/unlocked-achievements";
  const res = await fetch(
    // Make this not this
    `${achievementEndpoint}`
  );
  return res.json();
}
export default Achievements;
