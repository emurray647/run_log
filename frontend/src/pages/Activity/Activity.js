import React from "react"

import ActivitySummaryComponent from "./components/ActivitySummaryComponent"
import LapSummary from "./components/LapSummary"

import UnitConversion from "../../util/UnitConversion"
import Chart from "./components/Chart"
import Graphs from "./components/Graphs"
import MapComponent from "./components/MapComponent"

import apiCall from "../../util/Api"

class Activity extends React.Component {
    
    constructor(props) {
        super()
        this.state = {
            activity_id: props.match.params.id,
            loaded: false
        }

        this.setSelectedLap = this.setSelectedLap.bind(this);
    }

    componentDidMount() {
        apiCall("http://localhost:8080/api/v1/users/1/activities/" + this.state.activity_id)
        .then(response => {
            response.summary.start_time = UnitConversion.instance().convertEpochToDate(response.summary.start_time)
            this.setState({
                summary: response.summary,
                records: response.records,
                laps: response.laps,
            })

            this.setState({
                loaded: true
            })
        })

    }

    setSelectedLap(index) {

        if (index === null) {
            this.setState({
                focusedData: null
            })
        } else {
            const start_time = this.state.laps[index].start_time;
            const end_time = this.state.laps[index].end_time;
            this.setState({
                focusedData: {
                    start_time: start_time,
                    end_time: end_time
                }
            })
        }

    }
    
    render() {
        console.log(this.state.records)
        return (
            <div>
                {this.state.loaded ? 
                    <div> 
                        <ActivitySummaryComponent data={this.state.summary}/>
                        <LapSummary data={this.state.laps} setSelectedLap={this.setSelectedLap}/>
                        {/* <Chart data={this.state.records} /> */}
                        <Graphs 
                            data={this.state.records} 
                            // highlightRange={{start_time: this.state.focusedData.start_time, end_time: TouchList.state.focusedData.end_time}}
                            highlightRange={this.state.focusedData ?? undefined}
                        />
                        <MapComponent 
                            data={this.state.records}
                            focusTimerange={this.state.focusedData ?? undefined}
                        />
                    </div>
                    :
                    <h1>Loading</h1>
                }

            </div>
        
        )
    }
}

export default Activity