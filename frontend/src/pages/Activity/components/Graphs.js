import { TimeRange, TimeSeries } from "pondjs";
import React from "react"
import ChartContainer from "react-timeseries-charts/lib/components/ChartContainer";
import ChartRow from "react-timeseries-charts/lib/components/ChartRow";
import Charts from "react-timeseries-charts/lib/components/Charts"
import LabelAxis from "react-timeseries-charts/lib/components/LabelAxis";
import LineChart from "react-timeseries-charts/lib/components/LineChart";
import Resizable from "react-timeseries-charts/lib/components/Resizable";
import ValueAxis from "react-timeseries-charts/lib/components/ValueAxis";
import styler from "react-timeseries-charts/lib/js/styler";

// much of this code was borrowed from 
// https://github.com/esnet/react-timeseries-charts/blob/master/src/website/packages/charts/examples/cycling/Index.js

const style = styler([
    { key: "distance", color: "#e2e2e2" },
    { key: "altitude", color: "#e2e2e2" },
    { key: "cadence", color: "#ff47ff" },
    { key: "power", color: "green", width: 1, opacity: 0.5 },
    { key: "temperature", color: "#cfc793" },
    { key: "speed", color: "steelblue", width: 1, opacity: 0.5 },
    { key: "heartrate", color: "steelblue", width: 1, opacity: 0.5 },
]);

class Graphs extends React.Component {
    constructor(props) {
        super(props);

        const channels = {
            distance: { units: "miles", label: "Distance", format: ",.1f", series: null, show: false },
            altitude: { units: "feet", label: "Altitude", format: "d", series: null, show: false },
            cadence: { units: "rpm", label: "Cadence", format: "d", series: null, show: true },
            heartrate: { units: "bpm", label: "Heart Rate", format: "d", series: null, show: true},
            speed: { units: "mph", label: "Speed", format: ",.1f", series: null, show: true }
        };

        const channelNames = ["speed", "heartrate", "cadence", "altitude", "distance"];
        const displayChannels = ["speed", "heartrate", "cadence"];

        this.state = {
            ready: false,
            channels,
            tracker: null,
            channelNames,
            displayChannels,
        };

        this.renderChart = this.renderChart.bind(this);
    }

    componentDidMount() {
        const points = {};
        let channels = this.state.channels;
        this.state.channelNames.forEach(name  => {
            points[name] = []
        });

        const data = this.props.data;
        const start_time = data[0].timestamp;
        const end_time = data[data.length - 1].timestamp - start_time

        for (let i = 0; i < data.length; i++) {

            const time = (data[i].timestamp - start_time) * 1000;

            points["speed"].push([time, data[i].speed])
            points["heartrate"].push([time, data[i].heartrate])
            points["cadence"].push([time, data[i].cadence])
            points["altitude"].push([time, data[i].altitude])
            points["distance"].push([time, data[i].distance])
            
        }

        for (let channelName of this.state.channelNames) {
            const series = new TimeSeries({
                name: this.state.channels[channelName].name,
                columns: ["time", channelName],
                points: points[channelName]
            })
        
            channels[channelName].series = series;
            channels[channelName].avg = parseInt(series.avg(channelName), 10)
            channels[channelName].max = parseInt(series.max(channelName), 10)
        
        }

        const minTime = channels.altitude.series.range().begin();
        const maxTime = channels.altitude.series.range().end();
        const minDuration = 10 * 60 * 1000;

        const initialRange = new TimeRange([0, end_time * 1000]);

        this.setState({
            ready: true,
            channels,
            minTime,
            maxTime,
            minDuration,
            timerange: initialRange,
            brushRange: initialRange,
        })

    } // componentDidMount

    renderChart() {
        const {timerange, displayChannels, channels, maxTime, minTime, minDuration} = this.state;

        const durationPerPixel = timerange.duration() / 800 / 1000;
        const rows = [];

        for (let channelName of displayChannels) {
            const charts = [];
            let series = channels[channelName].series;

            charts.push(
                <LineChart
                    key={`line-${channelName}`}
                    axis={`${channelName}_axis`}
                    series={series}
                    columns={[channelName]}
                    style={style}
                    breakLine
                />
            );

            let value = "--";

            const summary = [
                { label: "Max", value: 0},
                { label: "Avg", value: 0},
            ];

            rows.push(
                <ChartRow
                    height="100"
                    visible={channels[channelName].show}
                    key={`row-${channelName}`}
                >
                    <LabelAxis
                        id={`${channelName}_axis`}
                        label={channels[channelName].label}
                        values={summary}
                        min={0}
                        max={channels[channelName].max}
                        width={140}
                        type="linear"
                        format=",.1f"
                    />
                    <Charts>{charts}</Charts>
                    <ValueAxis
                        id={`${channelName}_valueaxis`}
                        value={value}
                        detail={channels[channelName].units}
                        width={80}
                        min={0}
                        max={35}
                    />
                </ChartRow>
            );
        }

        return (
            <ChartContainer
                timeRange={this.state.timerange}
                format="relative"
                showGrid={false}
                enablePanZoom
                maxTime={maxTime}
                minTime={minTime}
                minDuration={minDuration}
                trackerPosition={this.state.tracker}
                // onTimeRangeChanged={this.handleTimeRangeChange}
                // onChartResize={width => this.handleChartResize(width)}
                // onTrackerChanged={this.handleTrackerChanged}
            >
                {rows}
            </ChartContainer>
        )

        // return (
        //     <div>This is our chart</div>
        // )
    }

    render() {
        const {ready, channels, displayChannels} = this.state
        if (!ready) {
            return <div>Loading</div>
        }

        const chartStyle = {
            borderStyle: "solid",
            borderWidth: 1,
            borderColor: "#DDD",
            paddingTop: 10,
            marginBottom: 10
        };

        const brushStyle = {
            boxShadow: "inset 0px 2px 5px -2px rgba(189, 189, 189, 0.75)",
            background: "#FEFEFE",
            paddingTop: 10
        };

        const legend = displayChannels.map(name => ({
            key: name,
            label: channels[name].label,
            disabled: !channels[name].show
        }));


        return (
            <div className="row">
                <div className="col-md-12" style={chartStyle}>
                    <Resizable>
                        {ready ? this.renderChart() : <div>Loading</div>}
                    </Resizable>
                </div>
            </div>
        )
    }
}

export default Graphs