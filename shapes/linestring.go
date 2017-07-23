package shapes

import (
	"encoding/json"
	"github.com/golang/geo/s2"
)

type S2Polyline struct {
	s2.Polyline
}

type geoJSONLineString struct {
	Type   string      `json:"type"`
	BBox   []float64   `json:"bbox,omitempty"`
	Coords [][]float64 `json:"coordinates"`
}

func (p *S2Polyline) Coverage(maxLevel, maxCells int) []s2.CellID {
	rc := &s2.RegionCoverer{MaxLevel: maxLevel, MaxCells: maxCells}
	return rc.CellUnion(&p.Polyline)
}

func (p *S2Polyline) MarshalJSON() ([]byte, error) {
	lnglats := [][]float64{}
	lo := p.Polyline.RectBound().Lo().Normalized()
	hi := p.Polyline.RectBound().Hi().Normalized()
	for _, point := range p.Polyline {
		ll := s2.LatLngFromPoint(point)
		lnglats = append(lnglats, []float64{ll.Lng.Degrees(), ll.Lat.Degrees()})
	}
	gLS := &geoJSONLineString{"LineString", []float64{lo.Lng.Degrees(), lo.Lat.Degrees(), hi.Lng.Degrees(), hi.Lat.Degrees()}, lnglats}
	return json.Marshal(gLS)
}

func (p *S2Polyline) UnmarshalJSON(in []byte) error {
	pView := &geoJSONLineString{}
	err := json.Unmarshal(in, pView)
	if err != nil {
		return err
	}

	latlngs := []s2.LatLng{}
	for _, ll := range pView.Coords {
		latlngs = append(latlngs, s2.LatLngFromDegrees(ll[1], ll[0]))
	}

	*p = S2Polyline{*s2.PolylineFromLatLngs(latlngs)}

	return err
}
