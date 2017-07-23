package shapes

import (
	"encoding/json"
	"github.com/golang/geo/s2"
)

type S2Point struct {
	s2.Point
}

type geoJSONPoint struct {
	Type   string    `json:"type"`
	BBox   []float64 `json:"bbox,omitempty"`
	Coords []float64 `json:"coordinates"`
}

func (p *S2Point) Coverage(maxLevel, maxCells int) []s2.CellID {
	rc := &s2.RegionCoverer{MaxLevel: maxLevel, MaxCells: maxCells}
	return rc.Covering(p.Point)
}

func (p *S2Point) MarshalJSON() ([]byte, error) {
	lnglat := s2.LatLngFromPoint(p.Point)
	lo := p.Point.RectBound().Lo().Normalized()
	hi := p.Point.RectBound().Hi().Normalized()
	gPoint := &geoJSONPoint{"Point", []float64{lo.Lng.Degrees(), lo.Lat.Degrees(), hi.Lng.Degrees(), hi.Lat.Degrees()},
		[]float64{lnglat.Lng.Degrees(), lnglat.Lat.Degrees()}}
	return json.Marshal(gPoint)
}

func (p *S2Point) UnmarshalJSON(in []byte) error {
	pView := &geoJSONPoint{}
	err := json.Unmarshal(in, pView)
	if err != nil {
		return err
	}

	*p = S2Point{s2.PointFromLatLng(s2.LatLngFromDegrees(pView.Coords[1], pView.Coords[0]))}

	return err
}
