import React from "react";
import "./App.css";
import { useRecoilValue } from 'recoil';
import { achievements } from './atoms';

function Achievements() {
  const data = useRecoilValue(achievements)

  console.log(data);

  return (
    <div>
      <h3>{data.games[0].title}</h3>
      <ul>
        {data.games[0].unlockedAchievements.map((ach) => (
          <li key={ach.name}>{ach.name}</li>
        ))}
      </ul>
    </div>
  );
}

export default Achievements;
