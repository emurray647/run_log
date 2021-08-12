import React from "react";

import Preferences from "../../../util/Preferences";

import style from "./LapSummary.css"


class LapSummary extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            clicked: null,
            hovered: null
        }
        this.handleHover = this.handleHover.bind(this);
        this.handleClick = this.handleClick.bind(this);

        this.emit = this.props.setSelectedLap;
    }

    handleClick(index) {
        this.setState({
            clicked: index
        })
        
        //propagate the call to parent
        this.emit(index);
    }

    handleHover(index) {
        this.setState({
            hovered: index
        })
    }

    render() {

        return (
            <div>
                <h1>Laps</h1>
                <table class="table">
                    <thead>
                        <tr> 
                            <th scope="col">Number</th>
                            <th scope="col">Time</th>
                            <th scope="col">Distance</th>
                            <th scope="col">Heart Rate</th>
                            <th scope="col">Cadence</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.props.data.map((lap, index) => {
                            const classNames = [
                                this.state.hovered === index ? "hovered" : "",
                                this.state.clicked === index ? "clicked" : ""
                            ].join(' ');
                            return (
                                <tr 
                                    key={index} 
                                    className={classNames}
                                    onMouseOver={() => this.handleHover(index)} 
                                    onClick={() => this.handleClick(index)}
                                    
                                >
                                    <td>{index + 1}</td>
                                    <td>{Preferences.instance().displayDuration(lap.duration)}</td>
                                    <td>{Preferences.instance().displayDistance(lap.distance)} {Preferences.instance().getDistanceUnits()}</td>
                                    <td>{lap.heartrate}</td>
                                    <td>{lap.cadence}</td>
                                </tr>
                            )
                        })}
                    </tbody>

                </table>
            </div>
        )

    }
}

export default LapSummary
