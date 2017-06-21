package shapes

import (
	"encoding/json"
	"github.com/golang/geo/s2"
)

type S2Polygon struct {
	Polygon s2.Polygon
}

type GeoJSONPolygon struct {
	Type   string        `json:"type"`
	Coords [][][]float64 `json:"coordinates"`
}

func (p *S2Polygon) MarshalJSON() ([]byte, error) {
	plnglats := [][][]float64{}
	for _, loop := range p.Polygon.Loops() {
		lnglats := [][]float64{}
		for _, point := range loop.Vertices() {
			ll := s2.LatLngFromPoint(point)
			lnglats = append(lnglats, []float64{ll.Lng.Degrees(), ll.Lat.Degrees()})
		}
		plnglats = append(plnglats, lnglats)
	}

	gP := &GeoJSONPolygon{"Polygon", plnglats}
	return json.Marshal(gP)
}

func (p *S2Polygon) UnmarshalJSON(in []byte) error {
	pView := &GeoJSONPolygon{}
	err := json.Unmarshal(in, pView)
	if err != nil {
		return err
	}

	loops := []*s2.Loop{}
	for _, lll := range pView.Coords {
		points := []s2.Point{}
		for _, ll := range lll {
			p := s2.PointFromLatLng(s2.LatLngFromDegrees(ll[1], ll[0]))
			points = append(points, p)
		}
		loop := s2.LoopFromPoints(points)
		loops = append(loops, loop)
	}

	s2p := s2.PolygonFromLoops(loops)
	*p = S2Polygon{*s2p}

	return err
}
