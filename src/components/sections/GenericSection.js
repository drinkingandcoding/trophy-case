import React from "react";
import PropTypes from "prop-types";
import classNames from "classnames";
import SectionHeader from "./partials/SectionHeader";
import { SectionProps } from "../../utils/SectionProps";
import { useRecoilValue } from 'recoil';
import { achievements } from '../../atoms';

import './Section.css';

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

  const data = useRecoilValue(achievements)
  const games = data["games"]

  return (
      <section {...props} className={outerClasses}>
      {games.map((game) =>
        <div className="container" key={game.title}>
          <div className={innerClasses}>
            <SectionHeader data={{title: game.title}} className="center-content" />
            <div className="game-section">
              <ul>
                {game.unlockedAchievements.map((ach) => (
                  <li key={ach.name}>{ach.name}</li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      )}
      </section>
  );
};

GenericSection.propTypes = propTypes;
GenericSection.defaultProps = defaultProps;

export default GenericSection;
