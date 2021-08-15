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
        console.log(viewPoint)

        viewPoint = fromLonLat(viewPoint)      

        var map = new Map({
            target: this.mapElement.current,
            layers: [
                new TileLayer({
                    source: new OSM()
                })
            ],
            view: new View({
                center: viewPoint,
                zoom: 15
            }),
        });

        const points = this.props.data.map(point => {
            return [point.position_longitude, point.position_latitude];
        })

        const polyline = new LineString(points);
        polyline.transform('EPSG:4326', 'EPSG:3857');
        const feature = new Feature(polyline);
        const source = new VectorSource(feature);
        source.addFeature(feature);
        const vectorLayer = new VectorLayer({source: source});

        map.addLayer(vectorLayer);

        this.setState({
            map: map,
        })
    }

    render() {


        return (
            <div>
                <div>This is the map</div>
                <div id="map" className="map-container" ref={this.mapElement}></div>
            </div>
            // <div id="map" className="map" ref={this.mapElement}> </div>
        )
    }
}

export default MapComponent
