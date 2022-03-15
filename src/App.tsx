import React from "react";
import "./App.css";
import "antd/dist/antd.css";
import Achievements from "./Achievements";
import { Input, Button } from "antd";
import { useRecoilState, atom } from "recoil";
import { setTextRange } from "typescript";

function App() {
  return (
    <div className="App">
      <h1>Trophy Case</h1>
      <IDInput></IDInput>
      <Achievements></Achievements>
    </div>
  );
}

const textState = atom({
  key: "steamID", // unique ID (with respect to other atoms/selectors)
  default: "", // default value (aka initial value)
});

function IDInput() {
  const [id, setID] = useRecoilState(textState);

  const onChange = (event: React.ChangeEvent<any>) => {
    setID(event.target.value);
    console.log(event.target.value);
  };

  const onClick = (event: React.ChangeEvent<any>) => {
    localStorage.setItem("Steam_ID", id);
    getAchievementsForUser(id);
  };

  return (
    <div>
      <Input
        style={{ width: "calc(25%)" }}
        value={id}
        onChange={onChange}
        placeholder="Steam ID"
      />
      <Button type="primary" onClick={onClick}>
        Submit
      </Button>
    </div>
  );
}

async function getAchievementsForUser(user: string) {
  const data = await fetch(
    `http://localhost:8888/.netlify/functions/steam-user`
  ).then((response) => response.json());

  console.log(data);
}

export default App;
