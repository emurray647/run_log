import React from "react"

import Preferences from "../../../util/Preferences"

class ActivitySummaryComponent extends React.Component {


    render() {
        console.log(this.props)
        return (
            <div>
                <ul>
                    <li>Time: {Preferences.instance().displayDate(this.props.data.start_time)}</li>
                    <li>Distance: {Preferences.instance().displayDistance(this.props.data.distance)}</li>
                    <li>Duration: {Preferences.instance().displayDuration(this.props.data.duration)}</li>
                    <li>Heart Rate: {this.props.data.avg_heart_rate}</li>
                    <li>Cadence: {this.props.data.avg_cadence}</li>
                </ul>
            </div>
        )
    }
}

export default ActivitySummaryComponent