package shapes

import (
	"encoding/json"
	"github.com/golang/geo/s2"
)

type S2Polyline struct {
	Polyline s2.Polyline
}

type GeoJSONLineString struct {
	Type   string      `json:"type"`
	Coords [][]float64 `json:"coordinates"`
}

func (p *S2Polyline) MarshalJSON() ([]byte, error) {
	lnglats := [][]float64{}
	for _, point := range p.Polyline {
		ll := s2.LatLngFromPoint(point)
		lnglats = append(lnglats, []float64{ll.Lng.Degrees(), ll.Lat.Degrees()})
	}
	gLS := &GeoJSONLineString{"LineString", lnglats}
	return json.Marshal(gLS)
}

func (p *S2Polyline) UnmarshalJSON(in []byte) error {
	pView := &GeoJSONLineString{}
	err := json.Unmarshal(in, pView)
	if err != nil {
		return err
	}

	latlngs := []s2.LatLng{}
	for _, ll := range pView.Coords {
		latlngs = append(latlngs, s2.LatLngFromDegrees(ll[1], ll[0]))
	}

	pl := s2.PolylineFromLatLngs(latlngs)
	*p = S2Polyline{*pl}

	return err
}
