package elevation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emurray647/run_log/internal/model"
)

//curl -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' -d '{"locations": [{"latitude": 10, "longitude": 10}]}' https://api.open-elevation.com/api/v1/lookup

type openElevationLocations struct {
	Locations []point `json:"locations"`
}

type openElevationResults struct {
	Locations []point `json:"results"`
}

type point struct {
	Lat       float32 `json:"latitude"`
	Long      float32 `json:"longitude"`
	Elevation float32 `json:"elevation,omitempty"`
}

const (
	apiMaxBlockSize = 1800
)

func FixElevations(activity *model.Activity) error {

	recordLocations := openElevationLocations{
		Locations: make([]point, len(activity.Records)),
	}
	for i, record := range activity.Records {
		recordLocations.Locations[i].Lat = record.Position.Latitude
		recordLocations.Locations[i].Long = record.Position.Longitude
	}

	for i := 0; i < len(recordLocations.Locations); i += apiMaxBlockSize {
		startPoint := i
		endPoint := i + apiMaxBlockSize
		if endPoint > len(recordLocations.Locations) {
			endPoint = len(recordLocations.Locations)
		}
		elevations, err := call(recordLocations.Locations[startPoint:endPoint])
		if err != nil {
			return err
		}

		for j := 0; j < len(elevations); j++ {
			activity.Records[i+j].Elevation = &elevations[j]
		}

	}

	ascent := 0.0
	descent := 0.0
	lastEl := *activity.Records[0].Elevation
	for _, record := range activity.Records[1:] {
		delta := *record.Elevation - lastEl
		if delta > 0 {
			ascent += float64(delta)
		} else {
			descent += float64(delta)
		}
		lastEl = *record.Elevation
	}

	activity.Ascent = float32(ascent)
	activity.Descent = float32(descent)

	return nil
}

func call(points []point) ([]float32, error) {
	result := make([]float32, len(points))

	inputLocations := openElevationLocations{Locations: points}

	content, err := json.Marshal(inputLocations)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal locations to JSON: %w", err)
	}

	url := "https://api.open-elevation.com/api/v1/lookup"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		return nil, fmt.Errorf("failed to create open-elevation api request: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("open-elevation api request failed : %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body of response: %w", err)
	}

	outputLocs := openElevationResults{}
	err = json.Unmarshal(body, &outputLocs)
	if err != nil {
		fmt.Println("failed to unmarshal")
		return nil, fmt.Errorf("could not unmarshal open-elevation response: %w", err)
	}

	for i := range result {
		result[i] = outputLocs.Locations[i].Elevation
	}

	return result, nil
}
