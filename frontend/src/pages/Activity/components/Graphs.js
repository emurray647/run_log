import { TimeRange, TimeRangeEvent, TimeSeries } from "pondjs";

import React from "react"
import AreaChart from "react-timeseries-charts/lib/components/AreaChart";
import Brush from "react-timeseries-charts/lib/components/Brush"
import ChartContainer from "react-timeseries-charts/lib/components/ChartContainer";
import ChartRow from "react-timeseries-charts/lib/components/ChartRow";
import Charts from "react-timeseries-charts/lib/components/Charts"
import LabelAxis from "react-timeseries-charts/lib/components/LabelAxis";
import LineChart from "react-timeseries-charts/lib/components/LineChart";
import Resizable from "react-timeseries-charts/lib/components/Resizable";
import ValueAxis from "react-timeseries-charts/lib/components/ValueAxis";
import YAxis from "react-timeseries-charts/lib/components/YAxis";
import styler from "react-timeseries-charts/lib/js/styler";

import UnitConversion from "../../../util/UnitConversion";

// much of this code was borrowed from 
// https://github.com/esnet/react-timeseries-charts/blob/master/src/website/packages/charts/examples/cycling/Index.js

const style = styler([
    { key: "distance", color: "#e2e2e2" },
    { key: "altitude", color: "#e2e2e2" },
    { key: "cadence", color: "#ff47ff" },
    { key: "power", color: "green", width: 1, opacity: 0.5 },
    { key: "temperature", color: "#cfc793" },
    { key: "speed", color: "steelblue", width: 1, opacity: 0.5 },
    { key: "pace", color: "steelblue", width: 1, opacity: 0.5 },
    { key: "heartrate", color: "steelblue", width: 1, opacity: 0.5 },
]);

function minsToTimeString(input) {
    const min = Math.floor(input);
    const sec = (input - min) * 60;
    const minString = min.toLocaleString('en-US', {minimumIntegerDigits: 2})
    const secString = sec.toLocaleString('en-US', {
        minimumIntegerDigits: 2,
        maximumFractionDigits: 0,
    })
    return `${minString}:${secString}`
}

class Graphs extends React.Component {
    constructor(props) {
        super(props);

        const channels = {
            distance: { units: "miles", label: "Distance", format: ",.1f", series: null, show: false },
            altitude: { units: "feet", label: "Altitude", format: "d", series: null, show: true },
            cadence: { units: "spm", label: "Cadence", format: "d", series: null, show: true },
            heartrate: { units: "bpm", label: "Heart Rate", format: "d", series: null, show: true},
            speed: { units: "mph", label: "Speed", format: ",.1f", series: null, show: true },
            pace: { units: "min/mile", label: "Pace", format: ",.1f", series: null, show: true, formatter: minsToTimeString},
        };

        const channelNames = ["speed", "heartrate", "cadence", "altitude", "distance", "pace"];
        const displayChannels = ["heartrate", "cadence", "pace", "altitude"];

        this.state = {
            ready: false,
            channels,
            tracker: null,
            channelNames,
            displayChannels,
        };

        this.renderChart = this.renderChart.bind(this);
        this.renderBrush = this.renderBrush.bind(this);
        this.handleTrackerChanged = this.handleTrackerChanged.bind(this);
        this.handleTimeRangeChange = this.handleTimeRangeChange.bind(this);
        this.handleChartResize = this.handleChartResize.bind(this);
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
            points["altitude"].push([time, data[i].elevation])
            points["distance"].push([time, data[i].distance])
            points["pace"].push([time, UnitConversion.instance().convertSpeedToPace(data[i].speed)])
            
        }

        console.log(points["speed"])

        for (let channelName of this.state.channelNames) {
            const series = new TimeSeries({
                name: this.state.channels[channelName].name,
                columns: ["time", channelName],
                points: points[channelName]
            })
        
            channels[channelName].series = series;
            channels[channelName].avg = parseFloat(series.avg(channelName))
            channels[channelName].max = parseFloat(series.max(channelName))
        
        }

        const minTime = channels.altitude.series.range().begin();
        const maxTime = channels.altitude.series.range().end();
        // const minDuration = 10 * 60 * 1000;
        const minDuration = 0;

        console.log("Setting the range")
        let initialRange;
        if (this.props.highlightRange) {
            initialRange = new TimeRange([this.props.highlightRange.start_time * 1000, this.props.highlightRange.end_time * 1000]);
        } else {
            initialRange = new TimeRange([0, end_time * 1000]);
        }

        this.setState({
            ready: true,
            channels,
            minTime,
            maxTime,
            minDuration,
            timerange: initialRange,
            brushrange: initialRange,
        })

    } // componentDidMount

    static getDerivedStateFromProps(nextProps, prevState) {
        if (nextProps.highlightRange) {
            const timerange = new TimeRange([
                nextProps.highlightRange.start_time * 1000,
                nextProps.highlightRange.end_time * 1000
            ]);
            return {
                timerange: timerange,
                brushrange: timerange
            }
        }
        return null;
    }

    handleTrackerChanged(t) {
        this.setState({tracker: t});
    }

    handleTimeRangeChange(tr) {
        const {channels} = this.state;

        if (tr) {
            this.setState({timerange: tr, brushrange: tr})
        }
        else {
            console.error("No time range provided")
        }
    }

    handleChartResize(width) {
        this.setState({width});
        console.log("Chart resize")
    }

    renderChart() {
        const {timerange, displayChannels, channels, maxTime, minTime, minDuration} = this.state;

        // const durationPerPixel = timerange.duration() / 800 / 1000;
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
            if (this.state.tracker) {
                const approx = 
                    (+this.state.tracker - +timerange.begin()) /
                    (+timerange.end() - +timerange.begin());
                const ii = Math.floor(approx * series.size());
                const i = series.bisect(new Date(this.state.tracker), ii);
                const v = i < series.size() ? series.at(i).get(channelName) : null;
                if (v) {
                    // value = parseInt(v, 10);
                    value = parseFloat(v, 10);
                    if (channels[channelName].formatter) {
                        value = channels[channelName].formatter(value);
                    }
                }
            }

            const formatter = channels[channelName].formatter ?? (value => {
                const string = value.toLocaleString('en-US', {
                    maximumFractionDigits: 1,
                })
                return string
            });
            const summary = [
                { label: "Max", value: formatter(channels[channelName].max)},
                { label: "Avg", value: formatter(channels[channelName].avg)},
            ];

            rows.push(
                <ChartRow
                    height="150"
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
                        // format=",.1f"
                        format={channels[channelName].format}
                    />
                    <Charts>{charts}</Charts>
                    <ValueAxis
                        id={`${channelName}_valueaxis`}
                        value={value}
                        detail={channels[channelName].units}
                        width={80}
                        min={0}
                        max={35}
                        format={channels[channelName].format}
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
                onTimeRangeChanged={this.handleTimeRangeChange}
                onChartResize={width => this.handleChartResize(width)}
                onTrackerChanged={this.handleTrackerChanged}
            >
                {rows}
            </ChartContainer>
        )

    }

    renderBrush() {
        const { channels } = this.state;
        return (
            <ChartContainer
                timeRange={channels.altitude.series.range()}
                format="relative"
                trackerPosition={this.state.tracker}
            >
                <ChartRow height="100" debug={false}>
                    <Brush
                        timeRange={this.state.brushrange}
                        allowSelectionClear
                        onTimeRangeChanged={this.handleTimeRangeChange}
                    />
                    <YAxis 
                        id="axis1"
                        label="Altitude"
                        min={0}
                        max={channels.altitude.max}
                        width={70}
                        type="linear"
                        format="d"
                    />
                    <Charts>
                        <AreaChart 
                            axis="axis1"
                            style={style.areaChartStyle()}
                            columns={{ up: ["altitude"], down: [] }}
                            series={channels.altitude.series}
                        />
                    </Charts>
                </ChartRow>

            </ChartContainer>
        )
    }

    render() {

        // if (this.props.)

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
                <div className="row">
                    <Resizable>
                        {ready ? this.renderBrush() : <div />}
                    </Resizable>
                </div>
            </div>
        )
    }
}

export default Graphs