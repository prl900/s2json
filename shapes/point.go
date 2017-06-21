package shapes

import (
	"encoding/json"
	"github.com/golang/geo/s2"
)

type S2Point struct {
	Point s2.Point
}

type GeoJSONPoint struct {
	Type   string    `json:"type"`
	Coords []float64 `json:"coordinates"`
}

func (p *S2Point) MarshalJSON() ([]byte, error) {
	lnglat := s2.LatLngFromPoint(p.Point)
	gPoint := &GeoJSONPoint{"Point", []float64{lnglat.Lng.Degrees(), lnglat.Lat.Degrees()}}
	return json.Marshal(gPoint)
}

func (p *S2Point) UnmarshalJSON(in []byte) error {
	pView := &GeoJSONPoint{}
	err := json.Unmarshal(in, pView)
	if err != nil {
		return err
	}

	*p = S2Point{s2.PointFromLatLng(s2.LatLngFromDegrees(pView.Coords[1], pView.Coords[0]))}

	return err
}
