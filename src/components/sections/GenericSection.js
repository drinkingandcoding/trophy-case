import React from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import SectionHeader from "./partials/SectionHeader";
import { SectionProps } from "../../utils/SectionProps";
import { useRecoilValue } from "recoil";
import { achievements } from "../../atoms";

import "./Section.css";

const propTypes = {
  children: PropTypes.node,
  ...SectionProps.types,
};

const defaultProps = {
  children: null,
  ...SectionProps.defaults,
};

const GenericSection = ({
  className,
  children,
  topOuterDivider,
  bottomOuterDivider,
  topDivider,
  bottomDivider,
  hasBgColor,
  invertColor,
  ...props
}) => {
  const outerClasses = classNames(
    "section",
    topOuterDivider && "has-top-divider",
    bottomOuterDivider && "has-bottom-divider",
    hasBgColor && "has-bg-color",
    invertColor && "invert-color",
    className
  );

  const innerClasses = classNames(
    "section-inner pt-0",
    topDivider && "has-top-divider",
    bottomDivider && "has-bottom-divider"
  );

  const data = useRecoilValue(achievements);
  const games = data["games"];
  console.log(games);

  return (
    <section {...props} className={outerClasses}>
      {games.map((game) => (
        <div className="container" key={game.title}>
          <div className={innerClasses}>
            <SectionHeader
            data={{
              title: game.title,
              paragraph: getBreakdown(game)
            }}
            className="center-content"
            />
            <div>
              <div className="game-section">
                {game.unlockedAchievements.map((ach) => (
                  <div key={ach.name} className="achievement">
                    <img src={ach.icon} alt='achievement img' className="achievement-img" style={getRarity(ach)}/> {ach.name}
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      ))}
    </section>
  );
};

const getBreakdown = (game) => {
  const achList = game.unlockedAchievements;
  var [common, uncommon, rare, ultra] = [0, 0, 0, 0]

  for (const ach of achList) {
    switch (true) {
      case (ach.rarity < 5):
        ultra++
        break
      case (ach.rarity < 20):
        rare++
        break
      case (ach.rarity < 50):
        uncommon++
        break
      default:
        common++
    }
  }

  return (
    <p>
      Achievement Breakdown: <span style={{color: '#f9a62b'}}>{ultra}</span> Ultra Rare / <span style={{color: '#583694'}}>{rare}</span> Rare / <span style={{color: '#3da560'}}>{uncommon}</span> Uncommon / {common} Common
    </p>
  )
}

const getRarity = (ach) => {
  const rarity = ach.rarity
  var border
  switch (true) {
    case (rarity < 5):
      border = '2px solid #f9a62b'
      break
    case (rarity < 20):
      border = '2px solid #583694'
      break
    case (rarity < 50):
      border = '2px solid #3da560'
      break
    default:
      border = ''
  }

  return {
    border: border
  }
}

GenericSection.propTypes = propTypes;
GenericSection.defaultProps = defaultProps;

export default GenericSection;
