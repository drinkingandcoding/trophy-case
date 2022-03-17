import React from "react";
import { Card, Col, Row } from "antd";

interface Props {
  data: object | null;
}

function Achievements(props: Props) {
  console.log(props.data);
  return (
    <div className="App">
      {/* <h2>Achievements for {data.playerstats.gameName} </h2> */}
      <Row gutter={16}>
        <Col span={8}>
          <Card title="Elden Ring">Put stuff here</Card>
        </Col>
        <Col span={8}>
          <Card title="Elden Ringer">Put stuff here again</Card>
        </Col>
      </Row>
    </div>
  );
}

export default Achievements;
