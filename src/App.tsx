import "./App.css";
import "antd/dist/antd.css";
import { Input, Button } from "antd";
import { useRecoilState, atom } from "recoil";
import { useQuery } from "react-query";

function App() {
  return (
    <div className="App">
      <h1>Trophy Case</h1>
      <IDInput></IDInput>
      <h2>Achievements</h2>
      <Achievements></Achievements>
    </div>
  );
}

const textState = atom({
  key: "steamID", // unique ID (with respect to other atoms/selectors)
  default: "", // default value (aka initial value)
});

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
      <h3>{data.game.gameName}</h3>
      <ul>
        {data.game.availableGameStats.achievements.map((ach) => (
          <p key={ach.name}>{ach.displayName}</p>
        ))}
      </ul>
    </div>
  );
}

function IDInput() {
  const [id, setID] = useRecoilState(textState);

  const onChange = (event: React.ChangeEvent<any>) => {
    setID(event.target.value);
  };

  const onClick = (event: React.ChangeEvent<any>) => {
    localStorage.setItem("Steam_ID", id);
    console.log("Getting achievements for %s", id);
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
async function getAchievementsForGame() {
  const achievementEndpoint = ".netlify/functions/steam-game";
  const res = await fetch(
    // Make this not this
    `${achievementEndpoint}`
  );
  return res.json();
}

export default App;
