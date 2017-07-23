package main

import (
	"./shapes"
	"encoding/json"
	"fmt"
	"github.com/golang/geo/s2"
)

func main() {
	pointStr := []byte(`{"type":"Point","coordinates":[125.6,10.1]}`)
	p := &shapes.S2Point{}
	err := json.Unmarshal(pointStr, p)
	if err != nil {
		panic(err)
	}

	cov := p.Coverage(5, 1)
	for _, cellID := range cov {
		cell := s2.CellFromCellID(cellID)
		for i := 0; i < 4; i++ {
			ll := s2.LatLngFromPoint(cell.Vertex(i))
			fmt.Printf("(%f, %f), ", ll.Lng.Degrees(), ll.Lat.Degrees())
		}
		fmt.Println("")
	}
	out, err := json.Marshal(&p)
	fmt.Println(string(out), err)

	//lsString := `{"type": "LineString", "coordinates": [[-100, 40], [-105, 45], [-110, 55]]}`
	lsString := []byte(`{"type":"LineString","coordinates":[[143.4375,-32.990235],[151.347656,-32.990235],[147.568359375,-26.82407078047018]]}`)
	pl := &shapes.S2Polyline{}
	err = json.Unmarshal(lsString, pl)
	if err != nil {
		panic(err)
	}
	cov = pl.Coverage(5, 1)
	for _, cellID := range cov {
		cell := s2.CellFromCellID(cellID)
		for i := 0; i < 4; i++ {
			ll := s2.LatLngFromPoint(cell.Vertex(i))
			fmt.Printf("(%f, %f), ", ll.Lng.Degrees(), ll.Lat.Degrees())
		}
		fmt.Println("")
	}
	out, err = json.Marshal(&pl)
	fmt.Println(string(out), err)

	pString := []byte(`{"type":"Polygon","coordinates":[[[143.4375,-32.990235],[151.347656,-32.990235],[147.568359375,-26.82407078047018],[143.4375,-32.990235]]]}`)
	poly := &shapes.S2Polygon{}
	err = json.Unmarshal(pString, poly)
	if err != nil {
		panic(err)
	}

	cov = poly.Coverage(5, 25)
	for _, cellID := range cov {
		cell := s2.CellFromCellID(cellID)
		for i := 0; i < 4; i++ {
			ll := s2.LatLngFromPoint(cell.Vertex(i))
			fmt.Printf("(%f, %f), ", ll.Lng.Degrees(), ll.Lat.Degrees())
		}
		fmt.Println("")
	}

	out, err = json.Marshal(&poly)
	fmt.Println(string(out), err)
}
