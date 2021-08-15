package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type LatLongPoint struct {
	Lat       float32 `json:"latitude"`
	Long      float32 `json:"longitude"`
	Elevation float32 `json:"elevation,omitempty"`
}

type location struct {
	Locations []LatLongPoint `json:"locations"`
}

type output struct {
	Locations []LatLongPoint `json:"results"`
}

const (
	BLOCK_SIZE = 1800
)

func GetElevations(points []LatLongPoint) ([]float32, error) {

	elevations := make([]float32, len(points))

	for i := 0; i < len(points); i += BLOCK_SIZE {
		endpoint := i + BLOCK_SIZE
		if endpoint > len(points) {
			endpoint = len(points)
		}
		_, err := callAPI(points[i:endpoint], elevations[i:endpoint])
		if err != nil {
			return elevations, err
		}
	}

	return elevations, nil
}

func callAPI(points []LatLongPoint, dst []float32) ([]float32, error) {

	// body := &bytes.Buffer{}
	locations := location{Locations: points}

	// locations.Locations = locations.Locations[:10]

	content, err := json.Marshal(locations)
	// content, err := json.MarshalIndent(locations, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	f, err := os.Create("/tmp/fittest")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	f.Write(content)

	fmt.Printf("Size: %d\n", len(content))
	fmt.Printf("Num: %d\n", len(points))

	// fmt.Println(string(content[:]))

	url := "https://api.open-elevation.com/api/v1/lookup"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(content))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	fmt.Println(res)

	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body[:]))

	outputLocs := output{}
	json.Unmarshal(body, &outputLocs)

	// fmt.Println(outputLocs)

	result := make([]float32, len(outputLocs.Locations))
	for i := range result {
		// result[i] = outputLocs.Locations[i].Elevation
		dst[i] = outputLocs.Locations[i].Elevation
	}

	return result, nil
}
