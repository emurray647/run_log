import React from "react";

import Preferences from "../../../util/Preferences";


class LapSummary extends React.Component {

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
                            return (
                                <tr key="{index}" >
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
        // return (
        //     <div>
        //         <ul>
        //             <li>Time: {Preferences.instance().displayDate(this.props.data.start_time)}</li>
        //             <li>Distance: {Preferences.instance().displayDistance(this.props.data.distance)}</li>
        //             <li>Duration: {Preferences.instance().displayDuration(this.props.data.duration)}</li>
        //             <li>Heart Rate: {this.props.data.avg_heart_rate}</li>
        //             <li>Cadence: {this.props.data.avg_cadence}</li>
        //         </ul>
        //     </div>
        // )
    }
}

export default LapSummary
