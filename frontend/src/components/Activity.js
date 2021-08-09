import React from "react"

import ActivitySummaryComponent from "./ActivitySummaryComponent"

import UnitConversion from "../util/UnitConversion"

class Activity extends React.Component {
    
    constructor(props) {
        super()
        this.state = {
            activity_id: props.match.params.id,
            loaded: false
        }
    }

    componentDidMount() {
        fetch("http://localhost:8080/api/v1/activities/" + this.state.activity_id)
        .then(response => response.json())
        .then(response => {
            response.summary.start_time = UnitConversion.instance().convertEpochToDate(response.summary.start_time)
            this.setState({
                summary: response.summary,
                records: response.records,
                laps: response.laps.filter(elt => elt.start_time = UnitConversion.instance().convertEpochToDate(elt.start_time))
            })

            this.setState({
                loaded: true
            })
        })

    }
    
    render() {
        return (
            <div>
                {this.state.loaded ? 
                   <ActivitySummaryComponent data={this.state.summary}/> :
                   <h1>Loading</h1>
                }

            </div>
        
        )
    }
}

export default Activity