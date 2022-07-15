import React from "react"

import ol from "ol"
import Map from "ol/Map"
import View from "ol/View"
import TileLayer from "ol/layer/Tile"
import VectorLayer from "ol/layer/Vector"
import Polyline from "ol/format/Polyline"

import {fromLonLat} from "ol/proj"

import VectorSource from "ol/source/Vector"
import OSM from "ol/source/OSM"
import Projection from "ol/proj/Projection"
import { Transform } from "ol/transform"
import { TransformFunction } from "ol/proj"

import style from "../Activity.css"
import LineString from "ol/geom/LineString"
import Feature from "ol/Feature"
import Style from "ol/style/Style"
import Stroke from "ol/style/Stroke"


class MapComponent extends React.Component {

    constructor() {
        super();
        this.state = {
            loaded: false
        }

        this.mapElement = React.createRef();

    }

    componentDidMount() {

        let viewPoint = [
            this.props.data[0].position_longitude,
            this.props.data[0].position_latitude,
        ]
        // console.log(viewPoint)

        viewPoint = fromLonLat(viewPoint)      

        var map = new Map({
            target: this.mapElement.current,
            layers: [
                new TileLayer({
                    source: new OSM()
                })
            ],
            // view: new View({
            //     center: viewPoint,
            //     zoom: 15
            // }),
        });

        // const view = new View({
        //     center: viewPoint,
        //     zoom: 15
        // });

        // map.setView(view);

        

        // const points = this.props.data.map(point => {
        //     return [point.position_longitude, point.position_latitude];
        // })

        // const polyline = new LineString(points);
        // polyline.transform('EPSG:4326', 'EPSG:3857');
        // const feature = new Feature(polyline);
        // const source = new VectorSource(feature);
        // source.addFeature(feature);
        // const vectorLayer = new VectorLayer({source: source});

        // map.addLayer(vectorLayer);

        this.setState({
            map: map,
        })

        
    }

    static getDerivedStateFromProps(nextProps, prevState) {

        const map = prevState.map;

        if (map) {

            let viewPoint = [
                nextProps.data[0].position_longitude,
                nextProps.data[0].position_latitude,
            ];

            viewPoint = fromLonLat(viewPoint)   
            const newView = new View({
                center: viewPoint,
                zoom: 15
            });
            map.setView(newView);

            if (prevState.focusLayer) {
                map.removeLayer(prevState.focusLayer)
            }

            const points = nextProps.data.map(point => {
                return [point.position_longitude, point.position_latitude];
            })
    
            const polyline = new LineString(points);
            polyline.transform('EPSG:4326', 'EPSG:3857');
            const feature = new Feature(polyline);
            const source = new VectorSource(feature);
            source.addFeature(feature);
            const vectorLayer = new VectorLayer({source: source});
    
            map.addLayer(vectorLayer);

        }   

        if (nextProps.focusTimerange) {
            const start_time = nextProps.focusTimerange.start_time;
            const end_time = nextProps.focusTimerange.end_time;

            console.log(start_time + " ... " + end_time)

            const focusPoints = nextProps.data.filter(point => {
                return point.timestamp >= start_time && point.timestamp <= end_time;
            }).map(point => {
                return [point.position_longitude, point.position_latitude];
            });
            
            // const focusPoints = nextProps.data.map(point => {
            //     return [point.position_longitude, point.position_latitude];
            // })
            // console.log(focusPoints.length)
            const polyline = new LineString(focusPoints);
            polyline.transform('EPSG:4326', 'EPSG:3857');
            const feature = new Feature(polyline);

            feature.setStyle(new Style({
                stroke: new Stroke({
                    width: 5,
                    color: "#ff0000"
                })
            }))
            // console.log(feature.style)
            const source = new VectorSource(feature);
            source.addFeature(feature);
            const vectorLayer = new VectorLayer({source: source});
            
    
            map.addLayer(vectorLayer);


            return {
                focus_start_time: start_time,
                focus_end_time: end_time,
                map: map,
                focusLayer: vectorLayer
            }

        }

        return {
            map: map
        };

    }

    render() {


        // this.setState({
        //     map: map,
        // })

        return (
            <div>
                <div>This is the map</div>
                <div id="map" className="map-container" ref={this.mapElement}></div>
            </div>
        )
    }
}

export default MapComponent
