import React from "react"

import Preferences from "../../util/Preferences"
import UnitConversion from "../../util/UnitConversion"

class ActivityList extends React.Component {

    constructor() {
        super()
        this.state = {
            activities: []
        }
        this.onRowClick = this.onRowClick.bind(this)
    }

    componentDidMount() {
        console.log("component did mount")
        fetch("http://localhost:8080/api/v1/activities")
        .then(response => response.json())
        .then(response => {
            console.log("length: " + response.length)
            console.log(response)
            for (let i = 0; i < response.length; i++) {
                console.log("The response")
                console.log(response[i])
                this.setState(prevState => {
                    const activity = {
                        id: response[i].id,
                        start_time: UnitConversion.instance().convertEpochToDate(response[i].start_time),
                        // total_time: UnitConversion.instance().convertDuration(response[i].total_time),
                        total_time: response[i].duration,
                        distance: response[i].distance,
                        avg_heart_rate: response[i].avg_heart_rate,
                        avg_cadence: response[i].avg_cadence,
                    }
                    const newActivities = [...prevState.activities, activity]
                    return {
                        activities: newActivities
                    }
                })
            }
        })

    }

    onRowClick(activity) {
        // event.stopPropagation();
        // console.log("onrowclick called")
        // console.log(event.target);
        console.log(activity.id)
        // this.nextPath("/activities/" + activity.id)
        this.props.history.push("/activities/" + activity.id)
    }

    onRowHover(event) {

    }

    render() {
        this.state.activities.forEach(act => console.log(act.id))
        return (
            <div>
                <table class="table">
                    <thead>
                        <tr> 
                            <th scope="col">Date</th>
                            <th scope="col">Time</th>
                            <th scope="col">Distance</th>
                            <th scope="col">Heart Rate</th>
                            <th scope="col">Cadence</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.state.activities.map(act => {
                            return (
                                <tr 
                                    key="{act.id}" 
                                    onClick={() => this.onRowClick(act)}>
                                    <td>{Preferences.instance().displayDate(act.start_time)}</td>
                                    <td>{Preferences.instance().displayDuration(act.total_time)}</td>
                                    <td>{Preferences.instance().displayDistance(act.distance)} {Preferences.instance().getDistanceUnits()}</td>
                                    <td>{act.avg_heart_rate}</td>
                                    <td>{act.avg_cadence}</td>
                                    <td>{act.id}</td>
                                </tr>
                            )
                        })
                        
                        }

                    </tbody>
                </table>
            </div>
        )
    }
}

export default ActivityList
