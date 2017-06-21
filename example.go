package main

import (
	"./shapes"
	"encoding/json"
	"fmt"
)

func main() {
	pointStr := []byte(`{"type":"Point","coordinates":[125.6, 10.1]}`)
	p := &shapes.S2Point{}
	err := json.Unmarshal(pointStr, p)
	if err != nil {
		panic(err)
	}
	out, err := json.Marshal(&p)
	fmt.Println(string(out), err)

	//lsString := `{"type": "LineString", "coordinates": [[-100, 40], [-105, 45], [-110, 55]]}`
	lsString := []byte(`{"type":"LineString","coordinates":[[40,-100],[45,-105],[55,-110]]}`)
	pl := &shapes.S2Polyline{}
	err = json.Unmarshal(lsString, pl)
	if err != nil {
		panic(err)
	}
	out, err = json.Marshal(&pl)
	fmt.Println(string(out), err)

	pString := []byte(`{"type":"Polygon","coordinates":[[[0,0],[10,10],[10,0],[0,0]]]}`)
	poly := &shapes.S2Polygon{}
	err = json.Unmarshal(pString, poly)
	if err != nil {
		panic(err)
	}
	out, err = json.Marshal(&poly)
	fmt.Println(string(out), err)
}
